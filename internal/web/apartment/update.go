package apartment

import (
	"asamblea/internal/models"
	"asamblea/internal/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) Update(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	var request models.ApartmentRequest
	err = c.ShouldBind(&request)
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	data := mapRequestToBO(request)
	data.ID = id

	err = h.service.Update(c, data)

	if err != nil {
		web.HandlerError(c, err)
		return
	}
	c.Redirect(http.StatusSeeOther, "/apartments/form/0")
}
