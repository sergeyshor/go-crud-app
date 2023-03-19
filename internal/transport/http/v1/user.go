package v1

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"go-crud-app/internal/dto"
	"go-crud-app/internal/entity"
	"go-crud-app/internal/transport/http/v1/middleware"
	"go-crud-app/internal/usecase"
)

type userHandlers struct {
	usecase.User
}

func newUserHandlers(superGroup *gin.RouterGroup, u usecase.User, m middleware.Middlewares) {
	handler := &userHandlers{u}

	// User group's routes
	userGroup := superGroup.Group("/user")
	{
		userGroup.GET("/:id", m.RequireAuth, handler.getUserByID)
		userGroup.POST("/signup", handler.signUp)
		userGroup.POST("/login", handler.login)
		userGroup.DELETE("/:id", m.RequireAuth, handler.deleteUserByID)
	}
}

func (h *userHandlers) getUserByID(c *gin.Context) {
	// Get params from request
	var req dto.GetUserByIdParams
	if c.ShouldBindUri(&req) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read uri",
		})
		return
	}

	// Check if user is authorized
	userKey, _ := c.Get("user")
	if _, ok := userKey.(*entity.User); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user must be authorized to get other user information",
		})
		return
	}

	// Check user existence by id
	user, err := h.GetByID(context.Background(), req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user with this is id does not exist",
		})
		return
	}

	response := dto.GetUserByIdDTO{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}
	c.JSON(http.StatusOK, response)
}

func (h *userHandlers) signUp(c *gin.Context) {
	// Get params from body
	var body dto.CreateUserDTO
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	body.Password = string(hash)

	// Create user
	err = h.Create(context.Background(), body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "account is successfully created",
	})
}

func (h *userHandlers) login(c *gin.Context) {
	// Get params from the body
	var body dto.LoginDTO
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Look up requested user in DB
	user, err := h.GetByEmail(context.Background(), body.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Compare sent in password with saved user password hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// Send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"msg": "welcome!",
	})
}

func (h *userHandlers) deleteUserByID(c *gin.Context) {
	// Get params from request
	var req dto.DeleteUserByIdParams
	if c.ShouldBindUri(&req) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read uri",
		})
		return
	}

	// Check if user with this id exists and delete it if so
	err := h.DeleteByID(context.Background(), req.ID)
	log.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user with this id does not exist",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("user with id=%v is deleted successful", req.ID),
	})
}
