package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primarykey"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

type Request struct {
	gorm.Model
	ID          uint          `json:"id" gorm:"primarykey"`
	UserID      uint          `json:"userId"`
	User        User          `json:"user" gorm:"foreignKey:UserID"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Status      string        `json:"status"`
	Attachments []Attachment  `json:"attachments,omitempty"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
}

type Attachment struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primarykey"`
	RequestID uint   `json:"requestId"`
	Filename  string `json:"filename"`
	Type      string `json:"type"`
	Size      int64  `json:"size"`
}

type ChatMessage struct {
	gorm.Model
	ID         uint      `json:"id" gorm:"primarykey"`
	RequestID  uint      `json:"requestId"`
	Request    Request   `json:"request" gorm:"foreignKey:RequestID"`
	UserID     uint      `json:"userId"`
	User       User      `json:"user" gorm:"foreignKey:UserID"`
	Message    string    `json:"message"`
	CreatedAt  time.Time `json:"createdAt"`
}
