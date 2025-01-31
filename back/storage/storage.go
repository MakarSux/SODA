package storage

import (
	"back/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage interface {
	// User methods
	GetUserByID(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error

	// Request methods
	CreateRequest(request *models.Request) error
	GetRequestByID(id uint) (*models.Request, error)
	GetAllRequests() ([]models.Request, error)
	GetRequestsByUserID(userID uint) ([]models.Request, error)
	UpdateRequest(request *models.Request) error

	// Attachment methods
	CreateAttachment(attachment *models.Attachment) error
	GetAttachmentsByRequestID(requestID uint) ([]models.Attachment, error)

	// Chat methods
	CreateChatMessage(message *models.ChatMessage) error
	GetChatMessagesByRequestID(requestID uint) ([]models.ChatMessage, error)
}

type PostgresStorage struct {
	db *gorm.DB
}

func NewPostgresStorage(dsn string) (*PostgresStorage, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Автоматическая миграция схемы
	err = db.AutoMigrate(&models.User{}, &models.Request{}, &models.Attachment{}, &models.ChatMessage{})
	if err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}

// User methods implementation
func (s *PostgresStorage) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := s.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (s *PostgresStorage) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := s.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (s *PostgresStorage) CreateUser(user *models.User) error {
	return s.db.Create(user).Error
}

func (s *PostgresStorage) UpdateUser(user *models.User) error {
	return s.db.Save(user).Error
}

// Request methods implementation
func (s *PostgresStorage) CreateRequest(request *models.Request) error {
	return s.db.Create(request).Error
}

func (s *PostgresStorage) GetRequestByID(id uint) (*models.Request, error) {
	var request models.Request
	result := s.db.Preload("Attachments").First(&request, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &request, nil
}

func (s *PostgresStorage) GetAllRequests() ([]models.Request, error) {
	var requests []models.Request
	result := s.db.Preload("Attachments").Find(&requests)
	if result.Error != nil {
		return nil, result.Error
	}
	return requests, nil
}

func (s *PostgresStorage) GetRequestsByUserID(userID uint) ([]models.Request, error) {
	var requests []models.Request
	result := s.db.Preload("Attachments").Where("user_id = ?", userID).Find(&requests)
	if result.Error != nil {
		return nil, result.Error
	}
	return requests, nil
}

func (s *PostgresStorage) UpdateRequest(request *models.Request) error {
	return s.db.Save(request).Error
}

// Attachment methods implementation
func (s *PostgresStorage) CreateAttachment(attachment *models.Attachment) error {
	return s.db.Create(attachment).Error
}

func (s *PostgresStorage) GetAttachmentsByRequestID(requestID uint) ([]models.Attachment, error) {
	var attachments []models.Attachment
	result := s.db.Where("request_id = ?", requestID).Find(&attachments)
	if result.Error != nil {
		return nil, result.Error
	}
	return attachments, nil
}

// Chat methods implementation
func (s *PostgresStorage) CreateChatMessage(message *models.ChatMessage) error {
	return s.db.Create(message).Error
}

func (s *PostgresStorage) GetChatMessagesByRequestID(requestID uint) ([]models.ChatMessage, error) {
	var messages []models.ChatMessage
	result := s.db.Where("request_id = ?", requestID).
		Preload("User").
		Order("created_at asc").
		Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}
	return messages, nil
}
