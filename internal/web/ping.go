package web

import (
	"net/http"
  
	"github.com/gin-gonic/gin"
  )

type handler struct{

}

func New() handler{
	return handler{}
}

func (* handler)Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
	  "message": "pong",
	})
  }