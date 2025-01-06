package Model

import (
	"mime/multipart"

	"gorm.io/gorm"
)

type PostInput struct {
	gorm.Model
	// Ubah default untuk membiarkan input manual
	Title      string                `form:"title" json:"title" binding:"required"`
	Content    string                `form:"content" json:"content" binding:"required"`
	UserID     string                `gorm:"foreignKey:UserID" json:"user_id"`
	Thumbnail  *multipart.FileHeader `form:"thumbnail" json:"thumbnail" file:"true"`
	Url        string                `form:"url" json:"url"`
	CategoryID uint                  `form:"category_id" json:"category_id"`
}

type CategoryInput struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"` // Menggunakan auto increment
	Name string `form:"name" json:"name" binding:"required"`
}
