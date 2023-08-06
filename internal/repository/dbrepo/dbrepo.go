package dbrepo

import (
	"database/sql"

	"github.com/thanhquy1105/bookings/internal/config"
	"github.com/thanhquy1105/bookings/internal/repository"
)

type postgresBDRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

type testBDRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresBDRepo{
		App: a,
		DB:  conn,
	}
}

func NewTestingRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testBDRepo{
		App: a,
	}
}
