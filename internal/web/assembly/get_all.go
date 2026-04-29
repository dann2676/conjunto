package assembly

import (
	"asamblea/internal/models"
	"asamblea/internal/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetAll(c *gin.Context) {
	assemblies, err := h.service.GetAll(c, false)
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	c.HTML(http.StatusOK, "assembly/base", gin.H{
		"assemblies": mapBOsToDTOs(assemblies),
		"assembly":   models.AssemblyDTO{},
		"is_edit":    false,
	})
}
