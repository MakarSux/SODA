package handlers

import (
	"back/models"
	"back/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type UserHandler struct {
	storage storage.Storage
}

func NewUserHandler(storage storage.Storage) *UserHandler {
	return &UserHandler{
		storage: storage,
	}
}

type UpdateProfileRequest struct {
	Name     string `json:"name"`
	Password string `json:"password,omitempty"`
}

func (h *UserHandler) Profile(c *fiber.Ctx) error {
	user := c.Locals("user")
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	userProfile, err := h.storage.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Не отправляем пароль в ответе
	userProfile.Password = ""
	return c.JSON(userProfile)
}

func (h *UserHandler) UpdateProfile(c *fiber.Ctx) error {
	user := c.Locals("user")
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	userProfile, err := h.storage.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Обновляем только разрешенные поля
	var req UpdateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Name != "" {
		userProfile.Name = req.Name
	}

	if err := h.storage.UpdateUser(userProfile); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update profile",
		})
	}

	// Не отправляем пароль в ответе
	userProfile.Password = ""
	return c.JSON(userProfile)
}

func (h *UserHandler) GetUserRequests(c *fiber.Ctx) error {
	user := c.Locals("user")
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	// Получаем информацию о пользователе для проверки роли
	userProfile, err := h.storage.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ошибка при получении данных пользователя",
		})
	}

	var requests []models.Request
	if userProfile.Role == "admin" {
		// Если пользователь админ, получаем все заявки
		requests, err = h.storage.GetAllRequests()
	} else {
		// Если обычный пользователь, получаем только его заявки
		requests, err = h.storage.GetRequestsByUserID(userID)
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ошибка при получении заявок",
		})
	}

	return c.JSON(requests)
}
