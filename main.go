package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const filename = "problems.csv"

func buildUpQuiz(quizLine [][]string) map[string]string {
	questionAndAnswers := make(map[string]string)
	for _, question := range quizLine {
		questionAndAnswers[question[0]] = question[1]
	}
	return questionAndAnswers
}

func setUpAndAskQuestion(question string, number *int) {
	fmt.Printf("Вопрос №%d: %s? ", *number, question)
	*number++
}

func getAndCheckAnswer(rightAnswer string) bool {
	rightAnswer = strings.ToLower(rightAnswer)
	rightAnswer = strings.TrimSpace(rightAnswer)

	var userAnswer string
	fmt.Scan(&userAnswer)
	userAnswer = strings.ToLower(userAnswer)

	return rightAnswer == userAnswer

}

func calcScore(correctAnswers, incorrectAnswers *[]string, correctness bool, answer string) {
	// Save answer and correctness for future improving of results printing
	if correctness {
		*correctAnswers = append(*correctAnswers, answer)
	} else {
		*incorrectAnswers = append(*incorrectAnswers, answer)
	}
}

func main() {
	fileName := flag.String("f", filename, "Insert filename")
	flag.Parse()

	var questionNumber = 1
	var correctAnswers, incorrectAnswers []string

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	questionAndAnswer, _ := reader.ReadAll()

	for question, answer := range buildUpQuiz(questionAndAnswer) {
		setUpAndAskQuestion(question, &questionNumber)
		answerAccuracy := getAndCheckAnswer(answer)
		calcScore(&correctAnswers, &incorrectAnswers, answerAccuracy, answer)
	}

	fmt.Printf("Количество правильных ответов: %d\n", len(correctAnswers))
	fmt.Printf("Количество неправильных ответов: %d\n", len(incorrectAnswers))

}
