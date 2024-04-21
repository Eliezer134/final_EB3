package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		verbo := ctx.Request.Method
		time := time.Now()
		url := ctx.Request.URL
		var size int

		ctx.Next()

		if ctx.Writer != nil {
			size = ctx.Writer.Size()

		}

		fmt.Printf("\nPath:%s\nVerbo:%s\nTiempo:%v\nTamanio consulta:%d\n", url, verbo, time, size)

	}
}
