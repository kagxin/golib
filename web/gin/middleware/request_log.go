package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// request Body ref:https://github.com/gin-gonic/gin/issues/961
// reponse Body ref:https://github.com/gin-gonic/gin/issues/1363

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// RequestLog 打印日志中间件
func RequestLog(getUserID func(*gin.Context) int, getPlatform func(*gin.Context) string, printLog func(string)) gin.HandlerFunc {

	return func(c *gin.Context) {
		var (
			blw              *bodyLogWriter
			requestLogString string
		)
		func() {
			defer func() {
				if err := recover(); err != nil {
					printLog(fmt.Sprintf("RequestLog1 ERROR %#v", err))
				}
			}()
			body, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				printLog("requestLog ioutil.ReadAll(c.Request.Body) error:%#v")
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			requestLogString = fmt.Sprintf("[Request]UserID:%d,Method:%s,RequestURI:%s,Body:%s;",
				getUserID(c), c.Request.Method, c.Request.RequestURI, string(body))

			blw = &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
			c.Writer = blw
		}()

		c.Next()

		func() {
			defer func() {
				if err := recover(); err != nil {
					printLog(fmt.Sprintf("RequestLog2 ERROR %#v", err))
				}
			}()
			statusCode := c.Writer.Status()
			bodyString := blw.body.String()
			responseLogString := fmt.Sprintf("[Response],Status:%d,Body:%s\n", statusCode, bodyString)
			printLog(fmt.Sprintf("[DEBUGLOGUSERID_%d]_%s", getUserID(c), requestLogString+responseLogString))
		}()
	}
}
