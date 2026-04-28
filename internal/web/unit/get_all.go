package unit

import (
	"asamblea/internal/models"
	"asamblea/internal/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetAll(c *gin.Context) {
	r, err := h.service.GetAll(c, false)
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	c.HTML(http.StatusOK, "unit/base", gin.H{"units": mapBOsToDTOs(r),
		"unit":    models.UnitDTO{},
		"is_edit": false})
}
