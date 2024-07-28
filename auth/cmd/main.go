package main

import (
	"auth/internal/handler"
	"auth/internal/middleware"
	"auth/lib/config"
	"auth/lib/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	app := fiber.New()
	cfg := config.LoadConfig()

	dbConn, err := db.ConnectToDB(cfg)
	if err != nil {
		panic("Failed to connect to database")
	}
	defer dbConn.Close()

	err = db.CreateTableIfNotExists(dbConn)
	if err != nil {
		panic("Failed to init database")
	}

	app.Get("/h", func(c *fiber.Ctx) error { return c.SendString("running") })

	app.Post("/signup", handler.SignupHandler)
	app.Post("/login", handler.LoginHandler)

	app.Use(middleware.AuthMiddleware)
	app.Get("/me", handler.CheckMeHandler)
	app.Post("/change-password", handler.ChangePasswordHandler)
	app.Post("/logout", handler.LogoutHandler)

	app.Listen(":3000")
}
