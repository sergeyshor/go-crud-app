package v1

import (
	"github.com/gin-gonic/gin"

	"go-crud-app/internal/transport/http/v1/middleware"
	"go-crud-app/internal/usecase"
)

type Handlers struct {
	userHandlers
	tweetHandlers
	middleware.Middlewares
}

func NewHandlers(u usecase.User, t usecase.Tweet) *Handlers {
	return &Handlers{
		userHandlers{u},
		tweetHandlers{t},
		*middleware.New(u, t),
	}
}

func (h *Handlers) NewRouter(r *gin.Engine) {
	// All entities' routes
	superGroup := r.Group("/api")
	{
		newUserHandlers(superGroup, h.userHandlers, h.Middlewares)
		newTweetHandlers(superGroup, h.tweetHandlers, h.Middlewares)
	}
}
