package store

import (
	"context"
	"database/sql"
	"fmt"
	"go-monolith-template/ent"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteAdapterOptions struct {
	InMemory bool
	FileName string
}

func ConnectSQLITE3(opts SqliteAdapterOptions, migrate bool) (*ent.Client, error) {
	if opts.InMemory {
		client, err := ent.Open("sqlite3", "file:ent?mode=memory&_fk=1")
		if err != nil {
			return nil, fmt.Errorf("failed opening connection to sqlite: %v", err)
		}
		if migrate {
			err = client.Schema.Create(context.Background())
			if err != nil {
				return nil, fmt.Errorf("failed creating schema resources: %v", err)
			}
		}
		return client, nil
	}
	dbConn, err := sql.Open("sqlite3", fmt.Sprintf("%v.db?_fk=1&mode=rwc&cache=shared&_journal_mode=WAL", opts.FileName))
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to sqlite: %v", err)
	}
	drv := entsql.OpenDB("sqlite3", dbConn)
	client := ent.NewClient(ent.Driver(drv))
	if migrate {
		err = client.Schema.Create(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed creating schema resources: %v", err)
		}
	}
	return client, nil
}

type PostgresAdapterOptions struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	SSLMode  string
}

func ConnectPostgres(opts PostgresAdapterOptions, migrate bool) (*ent.Client, error) {
	client, err := ent.Open("postgres",
		fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v",
			opts.Host, opts.Port, opts.Username, opts.Database, opts.Password, opts.SSLMode))
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %v", err)
	}
	if migrate {
		// Run the auto migration tool.
		if err := client.Schema.Create(context.Background()); err != nil {
			return nil, fmt.Errorf("failed creating schema resources: %v", err)
		}
	}
	return client, nil
}
