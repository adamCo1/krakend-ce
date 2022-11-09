package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type HelloWorldMiddleware struct{}

func InitMiddleware() *HelloWorldMiddleware {
	return &HelloWorldMiddleware{}
}

func (pHelloMid *HelloWorldMiddleware) Apply(c *gin.Context) {
	fmt.Sprint("Hello world from Hello world middleware!")
	c.Next()
}
