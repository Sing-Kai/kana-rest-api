package hirigana

import (
	"net/http"

	"github.com/Sing-Kai/kana-rest-api.git/quiz"
	"github.com/gin-gonic/gin"
)

func HandleHirigana(ctx *gin.Context) {

	q := quiz.GetQuiz()
	// t, _ := json.Marshal(q)

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"quiz": t,
	// })

	ctx.JSON(http.StatusOK, q)
}
