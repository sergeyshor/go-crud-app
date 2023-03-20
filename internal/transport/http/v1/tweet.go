package v1

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-crud-app/internal/dto"
	"go-crud-app/internal/entity"
	"go-crud-app/internal/transport/http/v1/middleware"
	"go-crud-app/internal/usecase"
)

type tweetHandlers struct {
	usecase.Tweet
}

func newTweetHandlers(superGroup *gin.RouterGroup, t usecase.Tweet, m middleware.Middlewares) {
	handler := &tweetHandlers{t}

	// Tweet group's routes
	tweetGroup := superGroup.Group("/tweet")
	{
		tweetGroup.GET("/feed", handler.showFeed)
		tweetGroup.GET("/:id", handler.getTweetByID)
		tweetGroup.POST("/", m.RequireAuth, handler.post)
		tweetGroup.DELETE("/:id", m.RequireAuth, handler.deleteTweetByID)
	}
}

func (h *tweetHandlers) showFeed(c *gin.Context) {
	// Get params from request
	var req dto.ListTweetsRequest
	if c.ShouldBindQuery(&req) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read uri",
		})
		return
	}

	arg := dto.ListAllTweetsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	tweets, err := h.ListAllTweets(context.Background(), arg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, tweets)
}

func (h *tweetHandlers) getTweetByID(c *gin.Context) {
	// Get params from request
	var req dto.GetTweetByIdParams
	if c.ShouldBindUri(&req) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read uri",
		})
		return
	}

	// Get tweet existence by id
	tweet, err := h.GetByID(context.Background(), req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "tweet with this id does not exist",
		})
		return
	}

	c.JSON(http.StatusOK, tweet)
}

func (h *tweetHandlers) post(c *gin.Context) {
	// Get params from body
	var params dto.CreateTweetParams
	if c.Bind(&params) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Check if user is authorized
	userKey, _ := c.Get("user")
	if _, ok := userKey.(*entity.User); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user must be authorized to post tweet",
		})
		return
	}

	arg := dto.CreateTweetDTO{
		AuthorID: userKey.(*entity.User).ID,
		Content:  params.Content,
	}

	// Create post
	tweet, err := h.Create(context.Background(), arg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, tweet)
}

func (h *tweetHandlers) deleteTweetByID(c *gin.Context) {
	// Get params from request
	var req dto.DeleteTweetByIdParams
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
			"error": "user must be authorized to post tweet",
		})
		return
	}

	// Check if tweet to be deleted is userKey's tweet
	tweet, err := h.GetAuthorID(context.Background(), req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "tweet with this id does not exist",
		})
		return
	}
	if userKey.(*entity.User).ID != tweet.AuthorID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "user must have privileges",
		})
		return
	}

	// Check if tweet with this id exists and delete it if so
	err = h.DeleteByID(context.Background(), req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "tweet with this id does not exist",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("tweet with id=%v is deleted successful", req.ID),
	})
}
