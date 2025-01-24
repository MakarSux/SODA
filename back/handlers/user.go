package handlers

import (
    "errors"
    "github.com/gofiber/fiber/v2"
    "back/models"
    "back/storage"
    "back/utils"

    "fmt"
)

type UserHandler struct {
    storage *storage.SQLiteStorage
}

func NewUserHandler(storage *storage.SQLiteStorage) *UserHandler {
    return &UserHandler{storage: storage}
}

func (h *UserHandler) Profile(c *fiber.Ctx) error {
    email, err := utils.GetEmailFromJWT(c)
    if err != nil {
        return c.SendStatus(fiber.StatusUnauthorized)
    }

    user, err := h.storage.GetUserByEmail(email)
    if err != nil {
        return fmt.Errorf("failed to get user: %w", err)
    }
    if user == nil {
        return errors.New("user not found")
    }

    return c.JSON(models.ProfileResponse{
        Email: user.Email,
        Name:  user.Name,
    })
}