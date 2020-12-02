package middleware

import (
	mylog "Go-000/Week02/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"runtime/debug"
	"strings"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				stack := string(debug.Stack())
				mylog.Logger.Error("panic_error:", zap.String("stack_info:", fmt.Sprintf("error=%v, stack=%s", err, strings.ReplaceAll(stack, "\n", " <--- "))))
			}
		}()
		c.Next()
	}
}
