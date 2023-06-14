package model

import (
	"time"

	"gorm.io/gorm"
)

type DefaultModel struct {
	Id        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
