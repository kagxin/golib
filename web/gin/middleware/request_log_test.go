package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kagxin/golib/web/gin/tool"
)

func getUserID(c *gin.Context) int {
	return 1
}

func getPlatform(c *gin.Context) string {
	return "IOS"
}

func printLog(log string) {
	fmt.Println(log)
}

func Test_RequestLog(t *testing.T) {
	router := gin.Default()
	router.Use(RequestLog(getUserID, getPlatform, printLog))
	router.Any("/", func(c *gin.Context) {
		// c.String(http.StatusOK, "i am response")
		c.JSON(http.StatusOK, map[string]string{"iam": "ok"})
	})
	body := strings.NewReader("i am request body.")
	w := tool.PerformRequest(router, "POST", "/", body)
	fmt.Println(w)
}
