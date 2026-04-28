package unit

import (
	"asamblea/internal/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) EditForm(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	data, err := h.service.Get(c, id)
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	all, err := h.service.GetAll(c, false)
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	c.HTML(http.StatusOK, "unit/content", gin.H{
		"unit":    mapBOToDTO(data),
		"units":   mapBOsToDTOs(all),
		"is_edit": id != 0})
}
