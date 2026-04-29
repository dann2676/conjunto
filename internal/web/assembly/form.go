package assembly

import (
	"asamblea/internal/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) EditForm(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
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
	c.HTML(http.StatusOK, "assembly/content", gin.H{
		"assembly":   mapBOToDTO(data),
		"assemblies": mapBOsToDTOs(all),
		"is_edit":    id != 0,
	})
}
