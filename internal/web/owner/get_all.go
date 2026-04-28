package owner

import (
	"asamblea/internal/models"
	"asamblea/internal/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetAll(c *gin.Context) {
	includeInactive := c.Query("todos") == "true"

	r, err := h.service.GetAll(c, includeInactive)
	apartments, err2 := h.apartments.GetAll(c)
	if err != nil || err2 != nil {
		web.HandlerError(c, err)
		return
	}

	c.HTML(http.StatusOK, "owner/base", gin.H{
		"owners":        mapBOsToDTOs(r),
		"owner":         models.OwnerDTO{},
		"apartments":    apartments,
		"is_edit":       false,
		"show_inactive": includeInactive,
	})
}
