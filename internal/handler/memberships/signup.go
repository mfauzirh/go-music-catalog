package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mfauzirh/go-music-catalog/internal/models/memberships"
	"github.com/rs/zerolog/log"
)

func (h *Handler) SignUp(c *gin.Context) {
	var req memberships.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed bind json request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.SignUp(req); err != nil {
		log.Error().Err(err).Msg("failed to sign up user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.Status(http.StatusCreated)
}
