package assembly

import (
	"asamblea/internal/models"
	"asamblea/internal/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	if err = h.service.Delete(c, id); err != nil {
		web.HandlerError(c, err)
		return
	}
	all, err := h.service.GetAll(c, false)
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	c.HTML(http.StatusOK, "assembly/content", gin.H{
		"assemblies": mapBOsToDTOs(all),
		"assembly":   models.AssemblyDTO{},
		"is_edit":    false,
	})
}
