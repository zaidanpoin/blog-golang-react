package Controller

import (
	"net/http"

	helper "github.com/zaidanpoin/blog-go/Helper"
	"github.com/zaidanpoin/blog-go/Model"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(context *gin.Context) {
	var input Model.AuthenticationInput

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := Model.User{
		ID:       helper.GenerateUUID(),
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
		Name:     input.Name,
		Role:     input.Role,
	}

	isi, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User created successfully!",
		"data":    isi,
	})
}

func Login(context *gin.Context) {
	var input LoginInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := Model.FindUserByUsername(input.Username)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.Writer.Header().Set("Authorization", "Bearer "+jwt)
	context.JSON(http.StatusOK, gin.H{"message": "Login success", "token": jwt})
}
