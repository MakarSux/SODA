package handlers

import (
	"back/models"
	"back/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type ChatHandler struct {
	store storage.Storage
}

func NewChatHandler(store storage.Storage) *ChatHandler {
	return &ChatHandler{store: store}
}

// SendMessage обработчик для отправки сообщения в чат
func (h *ChatHandler) SendMessage(c *fiber.Ctx) error {
	var message models.ChatMessage
	if err := c.BodyParser(&message); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный формат данных",
		})
	}

	// Получаем ID пользователя из JWT токена
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))
	message.UserID = userID
	message.CreatedAt = time.Now()

	// Проверяем существование заявки
	request, err := h.store.GetRequestByID(message.RequestID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Заявка не найдена",
		})
	}

	userInfo, err := h.store.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ошибка при получении данных пользователя",
		})
	}

	// Проверяем права доступа (только админ или создатель заявки могут отправлять сообщения)
	if userInfo.Role != "admin" && request.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Нет прав для отправки сообщений в этот чат",
		})
	}

	// Сохраняем сообщение
	if err := h.store.CreateChatMessage(&message); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ошибка при сохранении сообщения",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(message)
}

// GetMessages получение всех сообщений для конкретной заявки
func (h *ChatHandler) GetMessages(c *fiber.Ctx) error {
	requestID, err := c.ParamsInt("requestId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный ID заявки",
		})
	}

	// Получаем ID пользователя из JWT токена
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	// Проверяем права доступа
	request, err := h.store.GetRequestByID(uint(requestID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Заявка не найдена",
		})
	}

	userInfo, err := h.store.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ошибка при получении данных пользователя",
		})
	}

	// Проверяем, является ли пользователь админом или создателем заявки
	if userInfo.Role != "admin" && request.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Нет прав для просмотра сообщений этого чата",
		})
	}

	// Получаем сообщения
	messages, err := h.store.GetChatMessagesByRequestID(uint(requestID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ошибка при получении сообщений",
		})
	}

	return c.JSON(messages)
}
