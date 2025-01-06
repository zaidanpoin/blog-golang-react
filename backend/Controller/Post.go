package Controller

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zaidanpoin/blog-go/Model"
)

type PostResponse struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	UserID   string `-`
	UserName string `-`
	Username string `json:"username"`

	CategoryName string `json:"category_id"`
}

func GetPosts(c *gin.Context) {

	var post Model.Post
	posts, err := post.GetData(10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response []PostResponse
	for i := range posts {
		response = append(response, PostResponse{
			ID:       posts[i].ID,
			Title:    posts[i].Title,
			Content:  posts[i].Content,
			UserID:   posts[i].UserID,
			UserName: posts[i].User.Name,

			CategoryName: posts[i].Category.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"status":  http.StatusOK,
		"data":    response,
	})

}

func GetPostByID(c *gin.Context) {

	id := c.Param("id")

	var post Model.Post
	posts, err := post.GetPostById(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})

}

func CreatePost(c *gin.Context) {

	var input Model.PostInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "data": input})
		return
	}

	file, err := c.FormFile("thumbnail")
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Failed to retrieve file",
			"error":   err.Error(),
		})
		return
	}

	// Check file extension
	allowedExtensions := map[string]bool{
		".webp": true,
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	ext := filepath.Ext(file.Filename)
	if !allowedExtensions[ext] {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid file type. Only .webp, .jpg, and .png are allowed",
		})
		return
	}

	md5Hash := md5.New()
	md5Hash.Write([]byte(file.Filename + fmt.Sprintf("%d", time.Now().UnixNano())))
	hashedFilename := fmt.Sprintf("%x", md5Hash.Sum(nil))

	file.Filename = hashedFilename + ext

	// Simpan file ke lokasi tertentu (opsional)
	uploadPath := "./uploads/" + file.Filename
	url := "http://localhost:8080/uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Failed to save file",
			"error":   err.Error(),
		})
		return
	}

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "data": input})
		return
	}

	post := Model.Post{
		Title:      input.Title,
		Content:    input.Content,
		Thumbnail:  file.Filename,
		Url:        url,
		CategoryID: input.CategoryID,

		UserID: c.GetString("user_id"),
	}

	err1 := post.Save()

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error(), "data": post})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Post created successfully!",
	})

}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	var GetPost Model.Post

	GetPosts, err1 := GetPost.GetPostById(id, 1, 0)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data not found"})
		return
	}

	os.Remove("./uploads/" + GetPosts[0].Thumbnail)

	var post Model.Post
	err := post.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Post deleted successfully!",
	})

}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("user_id")

	post := Model.Post{}
	input := Model.PostInput{}

	GetData, err1 := post.GetPostById(id, 1, 0)
	GetUser, err2 := Model.FindUserById(userID)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data not found"})
		return
	}

	if GetData[0].UserID != userID && GetUser.Role != "Admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You are not authorized to update this post"})
		return
	}

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post = Model.Post{Title: input.Title, Content: input.Content, CategoryID: input.CategoryID}

	file, err := c.FormFile("thumbnail")
	if err == nil {

		err := os.Remove("./uploads/" + GetData[0].Thumbnail)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		allowedExtensions := map[string]bool{
			".webp": true,
			".jpg":  true,
			".jpeg": true,
			".png":  true,
		}

		ext := filepath.Ext(file.Filename)
		if !allowedExtensions[ext] {
			c.JSON(400, gin.H{
				"status":  "error",
				"message": "Invalid file type. Only .webp, .jpg, and .png are allowed",
			})
			return
		}

		md5Hash := md5.New()
		md5Hash.Write([]byte(file.Filename + fmt.Sprintf("%d", time.Now().UnixNano())))

		hashedFilename := fmt.Sprintf("%x", md5Hash.Sum(nil))

		file.Filename = hashedFilename + ext

		uploadPath := "./uploads/" + file.Filename

		url := "http://localhost:8080/uploads/" + file.Filename

		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			c.JSON(500, gin.H{
				"status":  "error",
				"message": "Failed to save file",
				"error":   err.Error(),
			})

			return
		}

		post.Thumbnail = file.Filename
		post.Url = url

	}

	if err := post.Update(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Post updated successfully!"})
}
