package owner

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
	apartments, err := h.apartemts.GetAll(c)

	if err != nil {
		c.HTML(http.StatusBadRequest, "owner/get", gin.H{"err": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "owner/content", gin.H{
		"owner":      mapBOToDTO(data),
		"owners":     mapBOsToDTOs(all),
		"apartments": apartments,
		"is_edit":    id != 0})
}
