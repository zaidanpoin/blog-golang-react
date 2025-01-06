package Model

import (
	"github.com/zaidanpoin/blog-go/Database"
)

type Category struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique"`
}

type Post struct {
	ID         uint     `json:"id" gorm:"primaryKey"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	UserID     string   `-`
	Thumbnail  string   `json:"thumbnail"`
	Url        string   `json:"url"`
	User       User     `json:"user" gorm:"foreignKey:UserID"`
	CategoryID uint     `json:"category_id"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryID"`
}

func (p *Post) GetData(limit int, offset int) ([]Post, error) {
	var posts []Post
	err := Database.Database.Joins("User").Joins("Category").Limit(limit).Offset(offset).Find(&posts).Error
	if err != nil {
		return posts, err
	}

	return posts, nil
}

func (p *Post) GetPostById(id string, limit int, offset int) ([]Post, error) {
	var post []Post

	err := Database.Database.Preload("User").Preload("Category").Limit(limit).Offset(offset).Where("id=?", id).Find(&post).Error

	if err != nil {
		return post, err
	}

	return post, nil

}

func (p *Post) Save() error {

	err := Database.Database.Create(&p).Error
	if err != nil {
		return err
	}

	return nil

}

func (p *Post) Delete(id string) error {

	err := Database.Database.Where("id=?", id).Delete(&p).Error
	if err != nil {
		return err
	}

	return nil

}

func (p *Post) Update(id string) error {

	err := Database.Database.Model(&Post{}).Where("id=?", id).Updates(p).Error
	if err != nil {
		return err
	}

	return nil

}
