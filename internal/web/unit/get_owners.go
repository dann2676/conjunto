package unit

import (
	"asamblea/internal/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetOwner(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	owner, err := h.owners.GetActiveByUnit(c, id)
	if err != nil {
		c.HTML(http.StatusOK, "unit/owner_detail", gin.H{"owner": nil})
		return
	}
	c.HTML(http.StatusOK, "unit/owner_detail", gin.H{"owner": mapOwnerBOToDTO(owner)})
}
