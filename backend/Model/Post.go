package Model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID         string   `gorm:"type:uuid;primaryKey"` // Ubah default untuk membiarkan input manual
	Title      string   `form:"title" json:"title" binding:"required"`
	Content    string   `form:"content" json:"content" binding:"required"`
	UserID     User     `gorm:"foreignKey:UserID" json:"user_id"`
	Thumbnail  string   `form:"thumbnail" json:"thumbnail" binding:"required"`
	CategoryID string   `form:"category_id" json:"category_id" binding:"required"`
	Category   Category `gorm:"foreignKey:CategoryID" json:"category_id"`
}

type Category struct {
	gorm.Model
	Name string `form:"name" json:"name" binding:"required"`
}
