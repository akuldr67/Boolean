package models

import (
	uuid "github.com/satori/go.uuid"
)

type Boolean struct {
	ID    uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	Key   string    `json:"key"`
	Value *bool     `json:"value" gorm:"not null" binding:"required"`
}
