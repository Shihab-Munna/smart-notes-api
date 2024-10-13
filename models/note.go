package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
    ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey:" json:"id"`
    Title   string    `gorm:"size:255;not null" json:"title"`
    Content string    `gorm:"size:555; not null" json:"content"`
}

// BeforeCreate is a GORM hook that runs before inserting a new note
func (note *Note) BeforeCreate(tx *gorm.DB) (err error) {
    note.ID = uuid.New()
    return
}