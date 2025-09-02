package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(URl string) error {
	db, _ = sql.Open("postgres", URl)
	return db.Ping()
}

func CreateTable() error {
	q := `CREATE TABLE IF NOT EXISTS history(
    id SERIAL PRIMARY KEY,
    user_ID BIGINT NOT NULL,
    date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    source_currency VARCHAR(3) NOT NULL,
    target_currency  VARCHAR(3) NOT NULL,
    amount NUMERIC(15, 2) NOT NULL,
    result NUMERIC(15, 2) NOT NULL
	);`
	_, err := db.Exec(q)
	if err != nil {
		return err
	}
	return nil
}
