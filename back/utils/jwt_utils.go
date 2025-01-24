package utils

import (
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v5"
    "time"
)

var jwtSecretKey = []byte("very-secret-key") // Замените на получение из переменных окружения

func GenerateJWT(email string) (string, error) {
    payload := jwt.MapClaims{
        "sub": email,
        "exp": time.Now().Add(time.Hour * 72).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
    return token.SignedString(jwtSecretKey)
}

func GetEmailFromJWT(c *fiber.Ctx) (string, error) {
    user := c.Locals("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    return claims["sub"].(string), nil
}