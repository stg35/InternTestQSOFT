package api

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

type yearRequest struct {
	Year int32 `uri:"year" binding:"required"`
}

func NewServer() *Server {
	server := &Server{}
	router := gin.Default()

	router.GET("/when/:year", server.checkDate)

	server.router = router
	return server
}

func (server *Server) checkDate(ctx *gin.Context) {
	var req yearRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Fatal(err.Error())
		return
	}

	nowDate := time.Now().UTC()
	nowDate = Date(nowDate.Year(), int(nowDate.Month()), nowDate.Day())
	whenDate := Date(int(req.Year), 1, 1)

	result := nowDate.Sub(whenDate).Hours() / 24

	if result > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"Days gone": result,
		})
		fmt.Println("Days gone: ", result)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Days left": math.Abs(result),
		})
		fmt.Println("Days left: ", math.Abs(result))
	}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}