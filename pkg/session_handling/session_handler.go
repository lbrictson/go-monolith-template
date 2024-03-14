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
	maxAgeSeconds int
	cookieName    string
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
