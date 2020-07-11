package response

import (
	"golib/web/common/eno"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSON 封装返回
func JSON(c *gin.Context, err error, data interface{}) {
	rc, ok := eno.ParseRCode(err)
	if ok != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    rc.Code,
		"message": rc.Message,
		"data":    data,
	})
}
