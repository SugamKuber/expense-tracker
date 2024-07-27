package main

import (
    "auth/internal/handler"
    "auth/internal/middleware"
    "auth/lib/config"
    "auth/lib/db"
    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
    "log"
    "fmt"

)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
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
        panic(fmt.Sprintf("Failed to create table: %s", err))
    }
    app.Post("/signup", handler.SignupHandler)
    app.Post("/login", handler.LoginHandler)

    app.Use(middleware.AuthMiddleware)
    app.Get("/me", handler.CheckMeHandler)
    app.Post("/change-password", handler.ChangePasswordHandler)
    app.Post("/logout", handler.LogoutHandler)

    app.Listen(":3000")

}
