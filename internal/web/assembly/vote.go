package assembly

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"asamblea/internal/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) VotePage(c *gin.Context) {
	slug := c.Param("id")
	assembly, err := h.service.GetBySlug(c, slug)
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	// Verificar que está registrado por cookie
	cookieID, err := c.Cookie("unit_id")
	if err != nil || cookieID == "" {
		c.Redirect(http.StatusSeeOther, "/assemblies/"+slug+"/attendance")
		return
	}

	// Buscar punto abierto a votación
	items, err := h.service.GetAgendaItems(c, assembly.ID)
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	var openItem *models.AgendaItemDTO
	for _, item := range items {
		if item.Status == "open" {
			dto := mapAgendaBOToDTO(item)
			openItem = &dto
			break
		}
	}

	// Buscar unidades del asistente
	attendance, err := h.service.GetAttendance(c, assembly.ID)
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	var myUnits []models.AssemblyUnitBO
	for _, a := range attendance {
		if a.AttendedByID == cookieID {
			myUnits = append(myUnits, a)
		}
	}

	c.HTML(http.StatusOK, "assembly/vote", gin.H{
		"assembly": mapBOToDTO(assembly),
		"item":     openItem,
		"myUnits":  myUnits,
	})
}

func (h *handler) RegisterVote(c *gin.Context) {
	slug := c.Param("id")
	assembly, err := h.service.GetBySlug(c, slug)
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	cookieID, err := c.Cookie("unit_id")
	if err != nil || cookieID == "" {
		web.HandlerError(c, domain.AsociationErr("debe registrar asistencia primero"))
		return
	}

	itemID, err := strconv.Atoi(c.PostForm("agenda_item_id"))
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	value := c.PostForm("value")
	unitIDs := c.PostFormArray("unit_id")

	for _, unitIDStr := range unitIDs {
		unitID, err := strconv.Atoi(unitIDStr)
		if err != nil || unitID == 0 {
			continue
		}
		bo := models.VoteBO{
			AgendaItemID: itemID,
			UnitID:       unitID,
			Value:        value,
		}
		// ignorar error de duplicado — ya votó esta unidad
		h.service.RegisterVote(c, bo)
	}

	c.HTML(http.StatusOK, "assembly/vote_confirmed", gin.H{
		"assembly": mapBOToDTO(assembly),
	})
}

func (h *handler) GetVoteResults(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	itemID, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	votes, err := h.service.GetVotes(c, itemID)
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	// calcular resultados
	results := map[string]float32{"yes": 0, "no": 0, "abstain": 0}
	counts := map[string]int{"yes": 0, "no": 0, "abstain": 0}
	for _, v := range votes {
		results[v.Value] += v.Coeficient
		counts[v.Value]++
	}

	c.HTML(http.StatusOK, "assembly/vote_results", gin.H{
		"assemblyID": id,
		"itemID":     itemID,
		"results":    results,
		"counts":     counts,
		"total":      len(votes),
	})
}
