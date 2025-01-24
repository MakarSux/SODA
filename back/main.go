package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/sirupsen/logrus"
    "back/handlers"
    "back/storage"
    "back/middleware"
    _ "github.com/mattn/go-sqlite3" // Импортируем драйвер SQLite
)

func main() {
    db, err := InitSQLite("users.db")
    if err != nil {
        logrus.Fatal(err)
    }
    defer db.Close()

    app := fiber.New()


    authHandler := handlers.NewAuthHandler(db)
    userHandler := handlers.NewUserHandler(db)

    // Public routes
    app.Post("/register", authHandler.Register)
    app.Post("/login", authHandler.Login)

    // Protected routes
    authorizedGroup := app.Group("")
    authorizedGroup.Use(middleware.JWTMiddleware())
    authorizedGroup.Get("/profile", userHandler.Profile)

    logrus.Fatal(app.Listen(":80"))
}