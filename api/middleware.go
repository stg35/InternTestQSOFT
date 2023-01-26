package api

import "github.com/gin-gonic/gin"

func ManageHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.Request.Header.Get("X-PING")
		if header == "ping" {
			ctx.Writer.Header().Add("X-PONG", "pong")
		}

		ctx.Next()
	}
}