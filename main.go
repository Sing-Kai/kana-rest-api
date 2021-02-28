package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Sing-Kai/kana-rest-api.git/quiz"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/healthy", health)
	r.GET("/quiz", quiz.HandleKanaQuiz)
	// r.GET("/assets/audio/:id", func(c *gin.Context) {
	// 	c.File("assets/shi.mp3")
	// })
	r.GET("/assets/audio/:id", handleAudioFile)
	r.Run()
}

func health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "healthy",
	})
}

func handleAudioFile(c *gin.Context) {
	id := c.Param("id")
	filePath := getAudioFilePath(id)
	c.File(filePath)
}

func getAudioFilePath(id string) string {

	i, err := strconv.Atoi(id)

	if err != nil {
		log.Fatal("error")
	}
	//add zere if single digit id
	if i/10 == 0 {
		id = "0" + id
	}

	filePath := search(id)

	return filePath
}

func search(id string) string {

	fileNames := make([]string, 0)
	root := "./assets/audio"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		fileNames = append(fileNames, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, file := range fileNames {
		if strings.Contains(file, id) {
			return file
		}
	}

	return ""
}
