package assembly

import (
	"asamblea/internal/domain"
	"asamblea/internal/web"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) DownloadReport(c *gin.Context) {
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

	if assembly.Status != "closed" {
		web.HandlerError(c, domain.AsociationErr("la asamblea debe estar cerrada para generar el reporte"))
		return
	}

	pdf, err := h.service.GenerateReport(c, id)
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	filename := fmt.Sprintf("acta-%s-%s.pdf",
		assembly.Slug,
		assembly.Date.Format("2006-01-02"),
	)

	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/pdf")
	c.Data(http.StatusOK, "application/pdf", pdf)
}
