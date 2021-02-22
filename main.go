package main

import (
	"net/http"

	"github.com/Sing-Kai/kana-rest-api.git/hirigana"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/healthy", handleTest)
	r.GET("/hirigana", hirigana.HandleHirigana)
	r.Run()
}

func handleTest(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}
