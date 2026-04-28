package apartment

import (
	"asamblea/internal/models"
	"asamblea/internal/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Create(c *gin.Context) {
	var request models.ApartmentRequest
	err := c.ShouldBind(&request)
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	err = h.service.Create(c, mapRequestToBO(request))

	if err != nil {
		web.HandlerError(c, err)
		return
	}
	c.Redirect(http.StatusSeeOther, "/apartments/form/0")
}
