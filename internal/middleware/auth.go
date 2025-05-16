package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bachacode/go-auth-jwt-example/internal/database"
	"github.com/bachacode/go-auth-jwt-example/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("Authorization")

		if tokenString == "" || err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		secret := []byte(os.Getenv("SECRET"))

		token, err := jwt.ParseWithClaims(tokenString, &handlers.Claims{}, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return secret, nil
		})

		if err != nil {
			log.Printf("Error parsing token: %v", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Now, check if token is nil BEFORE accessing its claims
		if token == nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(*handlers.Claims); ok {
			if time.Now().Unix() > claims.ExpiresAt.Unix() {
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			user := database.User{}
			database.DB.Where("email = ?", claims.Email).First(&user)

			if user.ID == 0 {
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			c.Set("user", user)

			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
