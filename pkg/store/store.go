package store

import (
	"database/sql"
	"go-monolith-template/ent"
)

type Storage struct {
	conn             *ent.Client
	extraConnections []*sql.DB
}

func NewStorage(conn *ent.Client, underlyingConnections ...*sql.DB) *Storage {
	return &Storage{conn: conn, extraConnections: underlyingConnections}
}

func (s *Storage) Close() error {
	return s.conn.Close()
}
