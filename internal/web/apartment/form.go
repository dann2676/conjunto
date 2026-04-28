package apartment

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

	all, err := h.service.GetAll(c)
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	c.HTML(http.StatusOK, "apartment/content", gin.H{
		"apartment":  mapBOToDTO(data),
		"apartments": mapBOsToDTOs(all),
		"is_edit":    id != 0})
}
