package main

import (
	"back/handlers"
	"back/middleware"
	"back/storage"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"os"
)

func main() {
	// Инициализация подключения к PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "1234"),
		getEnv("DB_NAME", "requests_db"),
		getEnv("DB_PORT", "5432"),
	)

	store, err := storage.NewPostgresStorage(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Инициализация Fiber
	app := fiber.New()

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	// Инициализация обработчиков
	authHandler := handlers.NewAuthHandler(store)
	requestHandler := handlers.NewRequestHandler(store)
	userHandler := handlers.NewUserHandler(store)
	chatHandler := handlers.NewChatHandler(store)

	// Маршруты аутентификации
	auth := app.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// Маршруты для заявок
	requests := app.Group("/requests", middleware.Protected())
	requests.Post("/submit", requestHandler.SubmitRequest)
	requests.Get("/", requestHandler.GetRequests)
	requests.Get("/:id", requestHandler.GetRequest)
	requests.Put("/:id/status", requestHandler.UpdateRequestStatus)

	// Маршруты для чата
	chat := app.Group("/chat", middleware.Protected())
	chat.Post("/messages", chatHandler.SendMessage)
	chat.Get("/messages/:requestId", chatHandler.GetMessages)

	// Маршруты для пользователей
	users := app.Group("/users", middleware.Protected())
	users.Get("/profile", userHandler.Profile)
	users.Put("/profile", userHandler.UpdateProfile)
	users.Get("/requests", userHandler.GetUserRequests)

	// Запуск сервера
	port := getEnv("PORT", "3000")
	log.Printf("Server is running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
