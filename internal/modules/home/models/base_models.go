package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	Uuid      uuid.UUID      `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time      // Popolato automaticamente da GORM alla creazione
	UpdatedAt time.Time      // Aggiornato automaticamente da GORM all'aggiornamento
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (baseModel *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if baseModel.Uuid == uuid.Nil {
		baseModel.Uuid = uuid.New()
	}
	return nil
}
