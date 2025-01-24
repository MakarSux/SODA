package storage

import (
    "database/sql"
    "fmt"
    "back/models"
    _ "github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
    db *sql.DB
}

// InitSQLite инициализирует SQLite и создает таблицу, если она не существует
func InitSQLite(dbPath string) (*SQLiteStorage, error) {
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        return nil, fmt.Errorf("failed to open database: %w", err)
    }

    // Создаем таблицу пользователей
    query := `
    CREATE TABLE IF NOT EXISTS users (
        email TEXT PRIMARY KEY,
        name TEXT NOT NULL,
        password TEXT NOT NULL
    );`
    _, err = db.Exec(query)
    if err != nil {
        return nil, fmt.Errorf("failed to create table: %w", err)
    }

    return &SQLiteStorage{db: db}, nil
}

// SaveUser сохраняет пользователя в базе данных
func (s *SQLiteStorage) SaveUser(user models.User) error {
    query := `INSERT INTO users (email, name, password) VALUES (?, ?, ?)`
    _, err := s.db.Exec(query, user.Email, user.Name, user.Password)
    return err
}

// GetUserByEmail возвращает пользователя по email
func (s *SQLiteStorage) GetUserByEmail(email string) (*models.User, error) {
    query := `SELECT email, name, password FROM users WHERE email = ?`
    row := s.db.QueryRow(query, email)

    var user models.User
    err := row.Scan(&user.Email, &user.Name, &user.Password)
    if err == sql.ErrNoRows {
        return nil, nil // Пользователь не найден
    }
    if err != nil {
        return nil, err
    }

    return &user, nil
}