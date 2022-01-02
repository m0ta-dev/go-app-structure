package bun

import (
	//"context"
	//"log"
	//"time"
	"database/sql"

	//"github.com/go-pg/pg/extra/pgotel"
	//"github.com/go-pg/pg/v10"
	//"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/dialect/pgdialect"
	//import "github.com/uptrace/bun/driver/pgdriver"

	"github.com/m0taru/go-app-structure/app/config"
	"github.com/m0taru/go-app-structure/app/utils"
)

// Timeout is a Postgres timeout
const Timeout = 5

// DB is a shortcut structure to a Postgres DB
type DB struct {
	*bun.DB
}

// Dial creates new database connection to postgres
func Dial() (*DB, error) {
	cfg := config.Get()
	if cfg.PgURL == "" {
		return nil, utils.ErrorNew("No URL to connect Postgre (bun/Dial)")
	}

	/*dsn, err := bun.ParseURL(cfg.PgURL)
	if err != nil {
		return nil, err
	}*/
	
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.PgURL)))
	db := bun.NewDB(sqldb, pgdialect.New())

	_, err := db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}