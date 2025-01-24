package handlers

import (
    "errors"
    "fmt"
    "github.com/gofiber/fiber/v2"
    "back/models"
    "back/storage"
    "back/utils"
)

type AuthHandler struct {
    storage *storage.SQLiteStorage
}

func NewAuthHandler(storage *storage.SQLiteStorage) *AuthHandler {
    return &AuthHandler{storage: storage}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
    var regReq models.RegisterRequest
    if err := c.BodyParser(&regReq); err != nil {
        return fmt.Errorf("body parser: %w", err)
    }

    // Проверяем, существует ли пользователь
    existingUser, err := h.storage.GetUserByEmail(regReq.Email)
    if err != nil {
        return fmt.Errorf("failed to check user existence: %w", err)
    }
    if existingUser != nil {
        return errors.New("the user already exists")
    }

    // Сохраняем пользователя
    user := models.User{
        Email:    regReq.Email,
        Name:     regReq.Name,
        Password: regReq.Password, // В реальном проекте пароль должен быть хэширован
    }
    if err := h.storage.SaveUser(user); err != nil {
        return fmt.Errorf("failed to save user: %w", err)
    }

    return c.SendStatus(fiber.StatusCreated)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
    var loginReq models.LoginRequest
    if err := c.BodyParser(&loginReq); err != nil {
        return fmt.Errorf("body parser: %w", err)
    }

    user, err := h.storage.GetUserByEmail(loginReq.Email)
    if err != nil {
        return fmt.Errorf("failed to get user: %w", err)
    }
    if user == nil || user.Password != loginReq.Password {
        return errors.New("email or password is incorrect")
    }

    token, err := utils.GenerateJWT(user.Email)
    if err != nil {
        return c.SendStatus(fiber.StatusInternalServerError)
    }

    return c.JSON(models.LoginResponse{AccessToken: token})
}