package db

import (
    "database/sql"

    "file-manager/lib/config"
    _ "github.com/lib/pq"
)

func ConnectToDB(cfg *config.Config) (*sql.DB, error) {
    db, err := sql.Open("postgres", cfg.DB_URI)
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}
