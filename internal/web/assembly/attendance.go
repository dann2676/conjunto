package assembly

import (
	"asamblea/internal/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetAllList(c *gin.Context) {
	includeInactive := c.Query("todos") == "true"

	r, err := h.service.GetAll(c, includeInactive)
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	c.HTML(http.StatusOK, "owner/list", gin.H{
		"owners":        mapBOsToDTOs(r),
		"show_inactive": includeInactive,
	})
}
