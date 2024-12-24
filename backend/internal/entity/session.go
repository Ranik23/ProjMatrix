package entity

import "github.com/google/uuid"

func GenerateSessionID() string {
	// Генерация нового UUID
	return uuid.New().String()
}
