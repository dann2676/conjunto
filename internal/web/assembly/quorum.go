package assembly

import (
	"asamblea/internal/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetQuorum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	assembly, err := h.service.Get(c, id)
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	attendance, err := h.service.GetAttendance(c, id)
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	quorum, err := h.service.GetQuorum(c, id)
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	c.HTML(http.StatusOK, "assembly/quorum_panel", gin.H{
		"assembly":   mapBOToDTO(assembly),
		"attendance": attendance,
		"quorum":     quorum,
	})
}
