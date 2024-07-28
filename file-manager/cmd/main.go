package main

import (
    "file-manager/internal/handler"
    "file-manager/internal/middleware"
    "file-manager/lib/config"
    "file-manager/lib/db"
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
    
    app.Use(middleware.AuthMiddleware)
    app.Get("/download/me", handler.TrackMe)
    app.Get("/download/all",handler.TrackAll)

    app.Listen(":3002")
}
