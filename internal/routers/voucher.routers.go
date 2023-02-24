package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VoucherRouter(rg *gin.RouterGroup) {
	voucher := rg.Group("/vouchers")

	voucher.GET("/get", func(c *gin.Context) {
		full_link := c.FullPath()
		c.JSON(http.StatusOK, gin.H{"Data": full_link})
	})
}
