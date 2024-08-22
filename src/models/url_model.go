package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Url struct {
	ID          uint   `gorm:"primary_key"`
	OriginalUrl string `gorm:"column:original_url;type:varchar(255)" json:"original_url,omitempty" validate:"required,url"`
	ShortUrl    string `gorm:"column:short_url;type:varchar(255)" json:"short_url,omitempty" validate:"required,url"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}

func (u *Url) TableName() string {
	return "url"
}

func AutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(&Url{}); err != nil {
		log.Fatalf("Error migrating url table: %v", err)
	}
}