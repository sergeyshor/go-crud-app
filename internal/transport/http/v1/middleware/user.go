package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"go-crud-app/internal/usecase"
)

type userMiddlewares struct {
	usecase.User
}

func (m *userMiddlewares) RequireAuth(c *gin.Context) {
	// Get the cookie from request
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode  token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	// Check errors from decoding and validate if token is valid
	if errors.Is(err, jwt.ErrTokenMalformed) ||
		errors.Is(err, jwt.ErrTokenExpired) ||
		errors.Is(err, jwt.ErrTokenNotValidYet) {
		c.AbortWithStatus(http.StatusUnauthorized)
	} else {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check the exp
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			// Find the user with token sub
			sub := claims["sub"].(float64)
			user, err := m.GetByID(context.Background(), int64(sub))
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			// Attach to request
			c.Set("user", user)
		}
	}

	// Continue
	c.Next()
}
