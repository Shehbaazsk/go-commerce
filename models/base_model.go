package models

import "time"

type BaseModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	IsActive  bool      `json:"is_active"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
