package assembly

import (
	"asamblea/internal/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) Admin(c *gin.Context) {
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
	items, err := h.service.GetAgendaItems(c, id)
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

	c.HTML(http.StatusOK, "assembly/admin", gin.H{
		"host":       c.Request.Host,
		"assembly":   mapBOToDTO(assembly),
		"items":      mapAgendaBOsToDTOs(items),
		"attendance": attendance,
		"quorum":     quorum,
	})
}
