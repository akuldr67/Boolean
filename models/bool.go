package models

import (
	uuid "github.com/satori/go.uuid"
)

type Boolean struct {
	// ID uint   `json:"id" gorm:"primary_key"`
	// ID    uuid.UUID `json:"id" gorm:"type:char(36);primary_key;default:uuid_generate_v4()"`
	ID    uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	Key   string    `json:"key"`
	Value *bool     `json:"value" gorm:"not null" binding:"required"`
}

// func (b *Boolean) tableName() string {
// 	return "booleans"
// }
