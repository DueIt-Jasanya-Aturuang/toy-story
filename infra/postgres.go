package infra

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func NewPostgresConnection() *sql.DB {
	fDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s",
		PgHost, PgPort, PgUser, PgPass, PgDB, PgSsl, PgSchema)

	db, err := sql.Open("postgres", fDB)
	if err != nil {
		log.Fatalf("failed connection to database")
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf("failed ping connection to database")
	}

	db.SetMaxIdleConns(setMaxIdleConnDB)
	db.SetMaxOpenConns(setMaxOpenConnDB)
	db.SetConnMaxIdleTime(SetConnMaxIdleTimeDB)
	db.SetConnMaxLifetime(setConnMaxLifetimeDB)

	log.Printf("successfully open connection to %s", PgDB)

	return db
}

const (
	setMaxIdleConnDB     = 5
	setMaxOpenConnDB     = 100
	SetConnMaxIdleTimeDB = 5 * time.Minute
	setConnMaxLifetimeDB = 60 * time.Minute
	pgPingTimeOut        = 2 * time.Second
)
