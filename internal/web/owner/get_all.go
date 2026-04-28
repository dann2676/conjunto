package owner

import (
	"asamblea/internal/models"
	"asamblea/internal/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetAll(c *gin.Context) {
	r, err := h.service.GetAll(c)
	apartments, err := h.apartemts.GetAll(c)

	if err != nil {
		web.HandlerError(c, err)
		return
	}

	c.HTML(http.StatusOK, "owner/base", gin.H{"owners": mapBOsToDTOs(r),
		"owner":      models.OwnerDTO{},
		"apartments": apartments,
		"is_edit":    false})
}
