package assembly

import (
	"asamblea/internal/web"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func (h *handler) GetCodes(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	assembly, err := h.service.Get(c, id)
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	codes, err := h.service.GetCodes(c, id)
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	c.HTML(http.StatusOK, "assembly/codes", gin.H{
		"assembly": mapBOToDTO(assembly),
		"codes":    codes,
		"host":     c.Request.Host,
	})
}

func (h *handler) DownloadQR(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	assembly, err := h.service.Get(c, id)
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	code := c.Param("code")
	url := fmt.Sprintf("http://%s/assemblies/%s/attendance?code=%s",
		c.Request.Host, assembly.Slug, code)

	var png []byte
	png, err = qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=qr-unidad-%s.png", code))
	c.Data(http.StatusOK, "image/png", png)
}

func (h *handler) QRImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		web.HandlerError(c, err)
		return
	}
	assembly, err := h.service.Get(c, id)
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	code := c.Param("code")
	url := fmt.Sprintf("http://%s/assemblies/%s/attendance?code=%s",
		c.Request.Host, assembly.Slug, code)

	png, err := qrcode.Encode(url, qrcode.Medium, 200)
	if err != nil {
		web.HandlerError(c, err)
		return
	}

	c.Data(http.StatusOK, "image/png", png)
}
