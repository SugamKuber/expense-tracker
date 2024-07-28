package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"tracker/lib/config"
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
    CREATE TABLE IF NOT EXISTS expenses (
        expense_id SERIAL PRIMARY KEY,
        expense_name VARCHAR(255) NOT NULL,
        total_amount DECIMAL(10, 2) NOT NULL,
        creator_id INTEGER NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (creator_id) REFERENCES users(user_id)
    );
    CREATE TABLE IF NOT EXISTS expense_tracker (
        tracker_id SERIAL PRIMARY KEY,
        expense_id INTEGER NOT NULL,
        user_id INTEGER NOT NULL,
        amount_owed DECIMAL(10, 2) NOT NULL,
        FOREIGN KEY (expense_id) REFERENCES expenses(expense_id),
        FOREIGN KEY (user_id) REFERENCES users(user_id)
    );
    
    `
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating tables: %w", err)
	}
	return nil

}
