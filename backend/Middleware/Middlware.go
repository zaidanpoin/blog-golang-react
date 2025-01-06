package Middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func RoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header Authorization
		tokenString := c.GetHeader("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Parse token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return privateKey, nil
		})

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Validasi klaim token
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Cek apakah peran pengguna cocok
			fmt.Println(claims["role"])
			if claims["role"] == role {
				c.Set("user_id", claims["id"]) // Set user_id ke context
				c.Next()                       // Lanjutkan ke handler berikutnya
			} else {
				c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
				c.Abort()
			}
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid token claims"})
			c.Abort()
		}
	}
}

func CheckLogin() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Ambil token dari header Authorization
		tokenString := c.GetHeader("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Parse token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return privateKey, nil
		})

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Validasi klaim token
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["id"]) // Set user_id ke context
			c.Next()                       // Lanjutkan ke handler berikutnya
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid token claims"})
			c.Abort()
		}
	}

}

func AdminMiddleware() gin.HandlerFunc {
	return RoleMiddleware("Admin")
}

func UserMiddleware() gin.HandlerFunc {
	return RoleMiddleware("User")
}
