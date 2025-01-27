package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

type Request struct {
	gorm.Model
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
	RequestID uint   `json:"requestId"`
	Filename  string `json:"filename"`
	Type      string `json:"type"`
	Size      int64  `json:"size"`
}
