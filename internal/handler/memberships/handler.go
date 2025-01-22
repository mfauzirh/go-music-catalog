package memberships

import (
	"github.com/gin-gonic/gin"
	"github.com/mfauzirh/go-music-catalog/internal/models/memberships"
)

type service interface {
	SignUp(req memberships.SignUpRequest) error
}

type Handler struct {
	api     *gin.Engine
	service service
}

func NewHandler(api *gin.Engine, service service) *Handler {
	return &Handler{
		api,
		service,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.api.Group("/memberships")
	route.POST("/signup", h.SignUp)
}
