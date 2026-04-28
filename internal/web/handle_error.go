package web

import (
	"asamblea/internal/domain"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerError(c *gin.Context, err error) {
	var e domain.DError
	status := http.StatusInternalServerError
	t := "danger"

	ok := errors.As(err, &e)
	if ok {
		status = e.Status
		t = e.Type

	}
	data := gin.H{"err": err.Error(),
		"type": t}
	c.HTML(status, "error", data)

}
