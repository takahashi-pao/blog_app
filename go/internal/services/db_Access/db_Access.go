package dbAccess

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func setupDB(dbDriver string, dsn string) *sql.DB {
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatalf("GetArticleData rows.Scan error err:%v", err)
	}

	return db
}

func AccessDB() *sql.DB {
	dbDriver := "postgres"
	dsn := "host=localhost port=5432 user=test password=password dbname=blogApp sslmode=disable"
	db := setupDB(dbDriver, dsn)

	return db
}
