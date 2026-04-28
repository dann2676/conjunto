package apartment

import (
	"asamblea/internal/models"
	"asamblea/internal/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetAll(c *gin.Context) {
	r, err := h.service.GetAll(c)
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	c.HTML(http.StatusOK, "apartment/base", gin.H{"apartments": mapBOsToDTOs(r),
		"apartment": models.ApartmentDTO{},
		"is_edit":   false})
}
