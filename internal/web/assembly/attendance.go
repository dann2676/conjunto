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
	cookieID, err := c.Cookie("unit_id")
	if err == nil && cookieID != "" {
		// verificar si esa cédula ya tiene registros en esta asamblea
		attendance, _ := h.service.GetAttendance(c, assembly.ID)
		for _, a := range attendance {
			if a.AttendedByID == cookieID {
				c.HTML(http.StatusOK, "assembly/already_registered", gin.H{
					"assembly": mapBOToDTO(assembly),
				})
				return
			}
		}
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
	slug := c.Param("id")

	assembly, err := h.service.GetBySlug(c, slug)
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	if assembly.Status != "open" {
		web.HandlerError(c, domain.AsociationErr("asamblea no está abierta"))
		return
	}

	// Datos del asistente
	attendedBy := c.PostForm("attended_by")
	attendedByID := c.PostForm("attended_by_id")
	ownerIDStr := c.PostForm("owner_id")

	var ownerID *int
	if ownerIDStr != "" {
		id, err := strconv.Atoi(ownerIDStr)
		if err == nil {
			ownerID = &id
		}
	}

	// Puede registrar múltiples unidades en un solo submit
	unitIDs := c.PostFormArray("unit_id")
	isProxies := c.PostFormArray("is_proxy")
	proxyFors := c.PostFormArray("proxy_for")

	var registered int
	for i, unitIDStr := range unitIDs {
		unitID, err := strconv.Atoi(unitIDStr)
		if err != nil || unitID == 0 {
			continue
		}

		isProxy := i < len(isProxies) && isProxies[i] == "true"
		proxyFor := ""
		if i < len(proxyFors) {
			proxyFor = proxyFors[i]
		}

		bo := models.AssemblyUnitBO{
			AssemblyID:   assembly.ID,
			UnitID:       unitID,
			OwnerID:      ownerID,
			AttendedBy:   attendedBy,
			AttendedByID: attendedByID,
			IsProxy:      isProxy,
			ProxyFor:     proxyFor,
		}

		if err = h.service.RegisterAttendance(c, bo); err != nil {
			// si ya está registrada esa unidad, continuar con las demás
			continue
		}
		registered++
	}

	if registered == 0 {
		web.HandlerError(c, domain.DuplicatedErr("todas las unidades ya estaban registradas"))
		return
	}

	// Setear cookie
	c.SetCookie("unit_id", attendedByID, 60*60*8, "/", "", false, true)

	c.HTML(http.StatusOK, "assembly/attendance_confirmed_partial", gin.H{
		"assembly":   mapBOToDTO(assembly),
		"registered": registered,
	})
}

func (h *handler) LookupOwner(c *gin.Context) {
	slug := c.Param("id")
	assembly, err := h.service.GetBySlug(c, slug)
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	var req models.AttendanceLookupRequest
	if err = c.ShouldBind(&req); err != nil {
		web.HandlerError(c, err)
		return
	}

	// Buscar dueño por cédula
	owner, err := h.owners.GetByIdentification(c, req.Identification)

	units, _ := h.units.GetAll(c, false)

	if err != nil {
		// No encontrado — mostrar form manual
		c.HTML(http.StatusOK, "assembly/attendance_manual", gin.H{
			"assembly":       mapBOToDTO(assembly),
			"identification": req.Identification,
			"units":          units,
		})
		return
	}

	// Encontrado — mostrar sus unidades
	ownerUnits, _ := h.owners.GetUnitsByOwner(c, owner.ID)
	c.HTML(http.StatusOK, "assembly/attendance_owner", gin.H{
		"assembly":   mapBOToDTO(assembly),
		"owner":      owner,
		"units":      units,
		"ownerUnits": ownerUnits,
	})
}
