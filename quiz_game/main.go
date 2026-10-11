package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Problem struct {
	Question string
	Answer   string
}

func main() {
	fileFlag := flag.String("fileName", "problems.csv", "file containing quiz data")
	timeFlag := flag.Int("timeLimit", 30, "time duration of the quiz")
	flag.Parse()

	bfReader := bufio.NewReader(os.Stdin)

	// validate and open a file
	file, err := os.Open(*fileFlag)
	if err != nil {
		log.Fatalf("failed to open the file: %s", err)
	}
	defer file.Close()

	// read a csv file all at once
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV: %s", err)
	}

	problems := parseData(data)

	timeLimit := time.Duration(*timeFlag) * time.Second
	fmt.Printf("Press Enter to Start the Quiz, you have %v sec time for the quiz\n", timeLimit)
	bfReader.ReadString('\n')

	resultChan := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), timeLimit)
	defer cancel()

	go startQuiz(ctx, resultChan, problems, bfReader)

	fmt.Printf("You scored %v out of %v\n", <-resultChan, len(data))
}

func startQuiz(ctx context.Context, resultChan chan<- int, problems []Problem, reader *bufio.Reader) {
	var correctQuestions int
	answerCh := make(chan string)

	for i, p := range problems {
		fmt.Printf("Problem No. %v, %s \n", i+1, p.Question)
		go func() {
			answer, _ := reader.ReadString('\n')
			answerCh <- answer
		}()

		select {
		case <-ctx.Done():
			fmt.Print("Time Over !!!\n")
			resultChan <- correctQuestions
			return
		case answer := <-answerCh:
			if strings.TrimSpace(answer) == strings.TrimSpace(p.Answer) {
				correctQuestions++
			}
		}
	}

	resultChan <- correctQuestions
}

func parseData(data [][]string) []Problem {
	var problems []Problem
	for _, d := range data {
		problem := Problem{
			Question: d[0],
			Answer:   d[1],
		}
		problems = append(problems, problem)
	}

	return problems
}
