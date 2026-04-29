package assembly

import (
	"asamblea/internal/models"
	"asamblea/internal/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateAgendaItem(c *gin.Context) {
	assemblyID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	var request models.AgendaItemRequest
	if err = c.ShouldBind(&request); err != nil {
		web.HandlerError(c, err)
		return
	}
	if err = h.service.CreateAgendaItem(c, mapAgendaRequestToBO(request, assemblyID)); err != nil {
		web.HandlerError(c, err)
		return
	}
	h.refreshAgenda(c, assemblyID)
}

func (h *handler) DeleteAgendaItem(c *gin.Context) {
	assemblyID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	itemID, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	if err = h.service.DeleteAgendaItem(c, itemID); err != nil {
		web.HandlerError(c, err)
		return
	}
	h.refreshAgenda(c, assemblyID)
}

func (h *handler) UpdateAgendaItemStatus(c *gin.Context) {
	assemblyID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	itemID, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	status := c.Param("status")
	if err = h.service.UpdateAgendaItemStatus(c, itemID, status); err != nil {
		web.HandlerError(c, err)
		return
	}
	h.refreshAgenda(c, assemblyID)
}

func (h *handler) refreshAgenda(c *gin.Context, assemblyID int) {
	items, err := h.service.GetAgendaItems(c, assemblyID)
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	c.HTML(http.StatusOK, "assembly/agenda_list", gin.H{
		"items":      mapAgendaBOsToDTOs(items),
		"assemblyID": assemblyID,
	})
}
