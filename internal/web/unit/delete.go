package unit

import (
	"asamblea/internal/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	err = h.service.Delete(c, id)

	if err != nil {
		web.HandlerError(c, err)
		return
	}
	c.Redirect(http.StatusSeeOther, "/units/form/0")
}
