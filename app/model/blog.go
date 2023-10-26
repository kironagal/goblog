package model

import "gorm.io/gorm"

type Blog struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Blog{})
}
