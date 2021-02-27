package main

import (
	"net/http"

	"github.com/Sing-Kai/kana-rest-api.git/quiz"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/healthy", health)
	r.GET("/quiz", quiz.HandleKanaQuiz)

	r.Run()
}

func health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "healthy",
	})
}
