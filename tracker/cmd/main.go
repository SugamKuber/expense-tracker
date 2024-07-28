package main

import (
    "tracker/internal/handler"
    "tracker/internal/middleware"
    "tracker/lib/config"
    "tracker/lib/db"
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

    app.Use(middleware.AuthMiddleware)
    app.Post("/add", handler.AddExpense)
    app.Get("/track/me", handler.TrackMe)
    app.Get("/track/all",handler.TrackAll)
    app.Get("/track/all/admin", handler.TrackAllAdmin)    

    app.Listen(":3001")
}
