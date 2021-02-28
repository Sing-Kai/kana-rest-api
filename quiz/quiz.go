package quiz

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Quiz struct {
	QuestionId int    `json:"questionid"`
	Question   Kana   `json:"question"`
	Answers    []Kana `json:"answers"`
}

type Kana struct {
	Id        int    `json:"id"`
	Hiri      string `json:"hiri"`
	Kata      string `json:"kata"`
	Syllabary string `json:"syllabary"`
}

func HandleKanaQuiz(ctx *gin.Context) {

	totalQuestions := 5
	numOfAnswers := 5
	resp := make([]Quiz, 0)
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= totalQuestions; i++ {
		q := getQuiz(numOfAnswers, i)
		resp = append(resp, q)
	}

	ctx.JSON(http.StatusOK, resp)
}

func getQuiz(ansNum int, questionId int) Quiz {

	kanaSlice := readFile()
	length := len(kanaSlice)
	question := kanaSlice[rand.Intn(length)]
	answerIds := getKanaIds(ansNum, question.Id, length)

	answers := make([]Kana, 0)

	for _, id := range answerIds {
		for _, k := range kanaSlice {
			if k.Id == id {
				answers = append(answers, k)
			}
		}
	}

	quiz := &Quiz{QuestionId: questionId, Question: question, Answers: answers}

	return *quiz

}

//gets rand list of character ids including answer
func getKanaIds(total, questionId, length int) []int {

	ans := make([]int, 0)

	uniqueIds := make(map[int]bool, 0)

	uniqueIds[questionId] = true

	j := 0

	for j < total-1 {
		//ensure randNum starts with 1
		randNum := rand.Intn(length-1) + 1

		if _, ok := uniqueIds[randNum]; !ok {
			uniqueIds[randNum] = true
			j++
		}
	}

	for k, _ := range uniqueIds {
		ans = append(ans, k)
	}

	// i := 0

	// for i < total-1 {
	// 	//ensure randNum starts with 1
	// 	randNum := rand.Intn(length-1) + 1
	// 	if randNum != questionId {
	// 		ans = append(ans, randNum)
	// 		i++
	// 	}
	// }

	// ans = append(ans, questionId)

	//shuffle answers
	rand.Shuffle(len(ans), func(i, j int) { ans[i], ans[j] = ans[j], ans[i] })

	return ans
}

func readFile() []Kana {

	file, _ := ioutil.ReadFile("./kana.json")
	data := make([]Kana, 0)
	_ = json.Unmarshal(file, &data)

	return data
}

func randomIndex(num int) int {
	rand := rand.Intn(num)
	return rand
}
