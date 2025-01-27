package middleware

import (
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v4"
    "net/http"
)

const SecretKey = "your-secret-key" // В продакшене использовать значение из конфига

// JWTMiddleware проверяет токен
func JWTMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        tokenString := c.Get("Authorization")
        if tokenString == "" {
            return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
                "error": "Unauthorized",
            })
        }

        // Удаляем префикс "Bearer " если он есть
        if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
            tokenString = tokenString[7:]
        }

        // Проверяем токен
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, http.ErrNotSupported
            }
            return []byte(SecretKey), nil
        })

        if err != nil || !token.Valid {
            return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
                "error": "Unauthorized",
            })
        }

        // Добавляем информацию о пользователе в контекст
        tokenClaims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
                "error": "Invalid token claims",
            })
        }
        userId := uint(tokenClaims["id"].(float64))
        userEmail := tokenClaims["email"].(string)
        userRole := tokenClaims["role"].(string)

        // Сохраняем данные в контексте
        c.Locals("user", token)  // Добавляем токен в контекст
        c.Locals("userId", userId)
        c.Locals("userEmail", userEmail)
        c.Locals("userRole", userRole)

        return c.Next()
    }
}

// AdminOnly middleware для проверки роли администратора
func AdminOnly() fiber.Handler {
    return func(c *fiber.Ctx) error {
        role := c.Locals("userRole")
        if role != "admin" {
            return c.Status(http.StatusForbidden).JSON(fiber.Map{
                "error": "Access denied",
            })
        }
        return c.Next()
    }
}

// Protected возвращает JWT middleware для защиты маршрутов
func Protected() fiber.Handler {
    return JWTMiddleware()
}