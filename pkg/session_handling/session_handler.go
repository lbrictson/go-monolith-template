package session_handling

import (
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"go-monolith-template/pkg/models"
	"go-monolith-template/pkg/store"
	"net/http"
	"time"
)

type SessionManager struct {
	dbConn        *store.Storage
	sessionCache  *cache.Cache
	sessionStore  sessions.Store
	notifications map[string][]Notification
	maxAgeSeconds int
	cookieName    string
}

type Notification struct {
	Header  string
	Message string
	IsError bool
}

type NewSessionManagerOptions struct {
	StorageLayer           *store.Storage
	DefaultCacheExpiration time.Duration
	CookieSecret           string
	SameSite               http.SameSite
	Secure                 bool
	HTTPOnly               bool
	CookieName             string
	MaxSessionAgeSeconds   int
}

func NewSessionManager(opts NewSessionManagerOptions) *SessionManager {
	s := &SessionManager{
		dbConn:        opts.StorageLayer,
		maxAgeSeconds: opts.MaxSessionAgeSeconds,
		cookieName:    opts.CookieName,
		notifications: make(map[string][]Notification),
	}
	s.sessionCache = cache.New(opts.DefaultCacheExpiration, opts.DefaultCacheExpiration)
	cookieStore := sessions.NewCookieStore([]byte(opts.CookieSecret))
	cookieStore.Options = &sessions.Options{
		Path:     "/",
		Secure:   opts.Secure,
		HttpOnly: opts.HTTPOnly,
		MaxAge:   opts.MaxSessionAgeSeconds,
		SameSite: opts.SameSite,
	}
	s.sessionStore = cookieStore
	return s
}

func (s *SessionManager) HasSession(c echo.Context) bool {
	sess, _ := s.sessionStore.Get(c.Request(), s.cookieName)
	_, ok := sess.Values["id"].(string)
	if !ok {
		return false
	}
	return true
}

func (s *SessionManager) Save(c echo.Context, user models.User) error {
	sessionDatabaseData, err := s.dbConn.SessionCreate(c.Request().Context(), store.CreateSessionOptions{
		UserID:       user.ID,
		MFACompleted: false,
	})
	if err != nil {
		return err
	}
	sess, _ := s.sessionStore.Get(c.Request(), s.cookieName)
	sess.Values["id"] = sessionDatabaseData.ID.String()
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return err
	}
	// save the session id in the cache, the TTL of the cache will be the same as the session duration
	s.sessionCache.Set(sessionDatabaseData.ID.String(), *sessionDatabaseData, time.Duration(s.maxAgeSeconds)*time.Second)
	return nil
}

func (s *SessionManager) Destroy(c echo.Context) error {
	sess, _ := s.sessionStore.Get(c.Request(), s.cookieName)
	id, ok := sess.Values["id"].(string)
	if !ok {
		return errors.New("session not found")
	}
	s.sessionCache.Delete(id)
	sess.Options.MaxAge = -1
	err := sess.Save(c.Request(), c.Response())
	// Remove in the database
	s.dbConn.SessionDelete(c.Request().Context(), uuid.MustParse(id))
	return err
}

func (s *SessionManager) GetSessionData(c echo.Context) (models.SessionData, error) {
	sess, _ := s.sessionStore.Get(c.Request(), s.cookieName)
	id, ok := sess.Values["id"].(string)
	if !ok {
		return models.SessionData{}, errors.New("session not found")
	}
	// Try getting from the cache first, odds are good it is there unless the server was restarted
	sessionData, found := s.sessionCache.Get(id)
	if found {
		// Try casting
		model, ok := sessionData.(models.SessionData)
		if ok {
			return model, nil
		}
	}
	// If not found in the cache, get from the database and put it in the cache
	dbSessionData, err := s.dbConn.SessionGetByID(c.Request().Context(), uuid.MustParse(id))
	if err != nil {
		return models.SessionData{}, err
	}
	s.sessionCache.Set(id, &sessionData, time.Duration(s.maxAgeSeconds)*time.Second)
	return *dbSessionData, nil
}

// Update will update the existing session in the database and the cache, returning an error if either fails or the
// session is not found.
func (s *SessionManager) Update(c echo.Context, opts store.UpdateSessionOptions) error {
	sess, _ := s.sessionStore.Get(c.Request(), s.cookieName)
	id, ok := sess.Values["id"].(string)
	if !ok {
		return errors.New("session not found")
	}
	// Update the database
	updated, err := s.dbConn.SessionUpdate(c.Request().Context(), uuid.MustParse(id), opts)
	if err != nil {
		return err
	}
	s.sessionCache.Set(id, &updated, time.Duration(s.maxAgeSeconds)*time.Second)
	return nil
}

// AddNotification will add a new notification to the session
func (s *SessionManager) AddNotification(id uuid.UUID, header string, message string, isError bool) {
	// Check if the session exists in the notification map
	_, ok := s.notifications[id.String()]
	if !ok {
		s.notifications[id.String()] = []Notification{
			{
				Header:  header,
				Message: message,
				IsError: isError,
			},
		}
		return
	}
	// Append the notification to the existing slice
	s.notifications[id.String()] = append(s.notifications[id.String()], Notification{
		Header:  header,
		Message: message,
		IsError: isError,
	})
	return
}

func (s *SessionManager) AddNotificationViaContext(c echo.Context, header string, message string, isError bool) {
	id, ok := c.Get("session_id").(uuid.UUID)
	if !ok {
		return
	}
	s.AddNotification(id, header, message, isError)
}

// GetNotifications will return the notifications for the session and clear them
func (s *SessionManager) GetNotifications(id uuid.UUID) []Notification {
	notifications, ok := s.notifications[id.String()]
	if !ok {
		return []Notification{}
	}
	// Clear the notifications
	delete(s.notifications, id.String())
	return notifications
}
