package handlers

import (
	"back/models"
	"back/storage"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"path/filepath"
)

type RequestHandler struct {
	storage storage.Storage
}

func NewRequestHandler(storage storage.Storage) *RequestHandler {
	return &RequestHandler{
		storage: storage,
	}
}

type SubmitRequestRequest struct {
	Title       string `form:"title"`
	Description string `form:"description"`
}

type UpdateRequestStatusRequest struct {
	Status string `json:"status"`
}

func (h *RequestHandler) SubmitRequest(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	// Парсим форму
	var req SubmitRequestRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Создаем заявку
	request := &models.Request{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		Status:      "pending",
	}

	if err := h.storage.CreateRequest(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create request",
		})
	}

	// Обрабатываем файл, если он есть
	file, err := c.FormFile("attachment")
	if err == nil && file != nil {
		// Создаем директорию для файлов, если её нет
		uploadDir := "./uploads"
		if err := c.SaveFile(file, fmt.Sprintf("%s/%d_%s", uploadDir, request.ID, file.Filename)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save file",
			})
		}

		// Сохраняем информацию о файле
		attachment := &models.Attachment{
			RequestID: request.ID,
			Filename:  file.Filename,
			Type:     filepath.Ext(file.Filename),
			Size:     int64(file.Size),
		}

		if err := h.storage.CreateAttachment(attachment); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save attachment info",
			})
		}
	}

	return c.JSON(request)
}

func (h *RequestHandler) GetRequests(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	// Только админ может видеть все заявки
	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access denied",
		})
	}

	requests, err := h.storage.GetAllRequests()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get requests",
		})
	}

	return c.JSON(requests)
}

func (h *RequestHandler) GetRequest(c *fiber.Ctx) error {
	requestID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request ID",
		})
	}

	request, err := h.storage.GetRequestByID(uint(requestID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Request not found",
		})
	}

	// Проверяем права доступа
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))
	role := claims["role"].(string)

	// Только админ или владелец заявки может её просматривать
	if role != "admin" && request.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access denied",
		})
	}

	return c.JSON(request)
}

func (h *RequestHandler) UpdateRequestStatus(c *fiber.Ctx) error {
	// Проверяем права доступа
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	// Только админ может обновлять статус
	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access denied",
		})
	}

	requestID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request ID",
		})
	}

	var req UpdateRequestStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Проверяем существование заявки
	request, err := h.storage.GetRequestByID(uint(requestID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Request not found",
		})
	}

	// Обновляем статус
	request.Status = req.Status
	if err := h.storage.UpdateRequest(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update request status",
		})
	}

	return c.JSON(request)
}
