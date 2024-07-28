package model

import (
	"auth/lib/config"
	"auth/lib/db"
)

type User struct {
	ID        int64  `json:"user_id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Mobile    string `json:"mobile"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

func CreateUser(user *User) error {
	dbConn, err := db.ConnectToDB(config.LoadConfig())
	if err != nil {
		return err
	}
	defer dbConn.Close()

	_, err = dbConn.Exec("INSERT INTO users (email, name, mobile, password, created_at) VALUES ($1, $2, $3, $4, NOW())",
		user.Email, user.Name, user.Mobile, user.Password)

	return err
}

func GetUserByEmail(email string) (*User, error) {
	dbConn, err := db.ConnectToDB(config.LoadConfig())
	if err != nil {
		return nil, err
	}
	defer dbConn.Close()

	row := dbConn.QueryRow("SELECT user_id, email, name, mobile, password, created_at FROM users WHERE email = $1", email)
	var user User
	err = row.Scan(&user.ID, &user.Email, &user.Name, &user.Mobile, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdatePassword(userID int64, newPassword string) error {
	dbConn, err := db.ConnectToDB(config.LoadConfig())
	if err != nil {
		return err
	}
	defer dbConn.Close()

	_, err = dbConn.Exec("UPDATE users SET password = $1 WHERE user_id = $2", newPassword, userID)
	return err
}

func GetUserByID(userID float64) (*User, error) {
	dbConn, err := db.ConnectToDB(config.LoadConfig())
	if err != nil {
		return nil, err
	}
	defer dbConn.Close()

	row := dbConn.QueryRow("SELECT user_id, email, name, mobile, password, created_at FROM users WHERE user_id = $1", userID)
	var user User
	err = row.Scan(&user.ID, &user.Email, &user.Name, &user.Mobile, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
