package middleware

import (
    "github.com/gofiber/fiber/v2"
    jwtware "github.com/gofiber/contrib/jwt"
)

func JWTMiddleware() fiber.Handler {
    return jwtware.New(jwtware.Config{
        SigningKey: jwtware.SigningKey{
            Key: []byte("very-secret-key"), // Замените на получение из переменных окружения
        },
        ContextKey: "user",
    })
}