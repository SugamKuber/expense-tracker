package db

import (
    "fmt"
    "database/sql"

    "auth/lib/config"
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

func CreateTableIfNotExists(db *sql.DB) error {
    query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        email VARCHAR(255) UNIQUE NOT NULL,
        name VARCHAR(255) NOT NULL,
        mobile VARCHAR(20),
        password VARCHAR(255) NOT NULL,
        created_at TIMESTAMPTZ DEFAULT NOW()
    );
    `
    _, err := db.Exec(query)
    if err != nil {
        return fmt.Errorf("error creating table: %w", err)
    }
    return nil
}
