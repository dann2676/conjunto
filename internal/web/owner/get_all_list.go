package owner

import (
	"asamblea/internal/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetAllList(c *gin.Context) {
	r, err := h.service.GetAll(c)

	if err != nil {
		web.HandlerError(c, err)
		return
	}

	c.HTML(http.StatusOK, "owner/list", gin.H{"owners": mapBOsToDTOs(r)})
}
