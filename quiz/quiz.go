package quiz

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

type Quiz struct {
	Question Kana
	Answers  []Kana
}

type Kana struct {
	Id        int    `json:"id"`
	Character string `json:"character"`
}

func GetQuiz() Quiz {

	numOfAnswers := 5
	kanaSlice := ReadFile()

	length := len(kanaSlice)
	question := kanaSlice[rand.Intn(length)]
	answerIds := getAnswerIds(numOfAnswers, question.Id, length)

	answers := make([]Kana, 0)

	for _, id := range answerIds {
		for _, k := range kanaSlice {
			if k.Id == id {
				answers = append(answers, k)
			}
		}
	}

	quiz := &Quiz{Question: question, Answers: answers}

	return *quiz

}

//gets rand list of character ids including answer
func getAnswerIds(total, id, length int) []int {

	ans := make([]int, 0)
	i := 0

	for i < total-1 {
		//ensure randNum starts with 1
		randNum := rand.Intn(length-1) + 1
		if randNum != id {
			ans = append(ans, randNum)
			i++
		}
	}

	ans = append(ans, id)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ans), func(i, j int) { ans[i], ans[j] = ans[j], ans[i] })

	return ans
}

func ReadFile() []Kana {

	file, _ := ioutil.ReadFile("./hirigana.json")
	data := make([]Kana, 0)
	_ = json.Unmarshal(file, &data)

	return data
}

func randomIndex(num int) int {
	rand := rand.Intn(num)
	return rand
}
