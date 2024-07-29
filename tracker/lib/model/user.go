package model

import (
	"tracker/lib/config"
	"tracker/lib/db"
)

type User struct {
	ID    int64  `json:"user_id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func GetUserByID(userID float64) (*User, error) {
	dbConn, err := db.ConnectToDB(config.LoadConfig())
	if err != nil {
		return nil, err
	}
	defer dbConn.Close()

	row := dbConn.QueryRow("SELECT user_id, email, name  FROM users WHERE user_id = $1", userID)
	var user User
	err = row.Scan(&user.ID, &user.Email, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
