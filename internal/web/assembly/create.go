package assembly

import (
	"asamblea/internal/models"
	"asamblea/internal/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Create(c *gin.Context) {
	var request models.AssemblyRequest
	if err := c.ShouldBind(&request); err != nil {
		web.HandlerError(c, err)
		return
	}
	bo, err := mapRequestToBO(request)
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	if err = h.service.Create(c, bo); err != nil {
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
