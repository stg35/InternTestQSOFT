package api

import (
	"log"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type yearRequest struct {
	Year int32 `uri:"year" binding:"required"`
}

func (server *Server) checkDate(ctx *gin.Context) {
	var req yearRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err.Error())
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
		log.Println("Days gone: ", result)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Days left": math.Abs(result),
		})
		log.Println("Days left: ", math.Abs(result))
	}
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}