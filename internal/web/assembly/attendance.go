package assembly

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"asamblea/internal/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) AttendancePage(c *gin.Context) {
	slug := c.Param("id")

	assembly, err := h.service.GetBySlug(c, slug)
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	switch assembly.Status {
	case "open":
		// continúa normal
	case "closed":
		c.HTML(http.StatusOK, "assembly/closed", gin.H{
			"assembly": mapBOToDTO(assembly),
		})
		return
	default: // draft
		c.HTML(http.StatusOK, "assembly/not_open", gin.H{
			"assembly": mapBOToDTO(assembly),
		})
		return
	}

	// Verificar si ya registró asistencia por cookie
	// Verificar cookie
	cookieUnitID, err := c.Cookie("unit_id")
	if err == nil && cookieUnitID != "" {
		// además verificar en DB que realmente está registrado
		unitID, _ := strconv.Atoi(cookieUnitID)
		attendance, _ := h.service.GetAttendance(c, assembly.ID)
		for _, a := range attendance {
			if a.UnitID == unitID {
				c.HTML(http.StatusOK, "assembly/already_registered", gin.H{
					"assembly": mapBOToDTO(assembly),
				})
				return
			}
		}
		// si tiene cookie pero no está en DB, limpiar cookie y continuar
		c.SetCookie("unit_id", "", -1, "/", "", false, true)
	}

	units, err := h.units.GetAll(c, false)
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	c.HTML(http.StatusOK, "assembly/attendance", gin.H{
		"assembly": mapBOToDTO(assembly),
		"units":    units,
	})
}

func (h *handler) RegisterAttendance(c *gin.Context) {
	id := c.Param("id")

	assembly, err := h.service.GetBySlug(c, id)
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	if assembly.Status != "open" {
		web.HandlerError(c, domain.AsociationErr("asamblea no está abierta"))
		return
	}

	var request models.AttendanceRequest
	if err = c.ShouldBind(&request); err != nil {
		web.HandlerError(c, err)
		return
	}

	bo := models.AssemblyUnitBO{
		AssemblyID:    assembly.ID,
		UnitID:        request.UnitID,
		AttendedBy:    request.AttendedBy,
		RepresentedBy: request.RepresentedBy,
	}

	if err = h.service.RegisterAttendance(c, bo); err != nil {
		web.HandlerError(c, err)
		return
	}

	// Setear cookie con HttpOnly y SameSite
	c.SetCookie(
		"unit_id",
		strconv.Itoa(request.UnitID),
		60*60*8, // 8 horas
		"/",
		"",
		false,
		true, // HttpOnly
	)

	c.HTML(http.StatusOK, "assembly/attendance_confirmed_partial", gin.H{
		"assembly": mapBOToDTO(assembly),
	})
}
