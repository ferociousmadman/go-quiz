package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	text        string = "sir/madam"
	storedInput []string
)

func main() {
	shuffle := flag.Bool("shuffle", false, "shuffle the quiz order")
	timeout := flag.Duration("timeout", 30*time.Second, "set a timeout duration")
	csvFilename := flag.String("csv", "problems.csv", "A csv file in the format of questions and answers")
	flag.Parse()
	fmt.Printf("You will be taking a timed quiz, %v. \n", text)
	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("csv file %s could not be opened", *csvFilename)
		os.Exit(1)
	}
	reader := csv.NewReader(file)
	records, errs := reader.ReadAll()
	if errs != nil {
		fmt.Printf("csv file critical failure")
		os.Exit(1)
	}
	if *shuffle {
		rand.Shuffle(len(records), func(i, j int) {
			records[i], records[j] = records[j], records[i]
		})
	}
	grade := 0
	scanner := bufio.NewScanner(os.Stdin)

	timer := time.NewTimer(*timeout)
	answerCh := make(chan string)
	for i := 0; i < len(records); i++ {
		userInput := ""
		fmt.Printf("Problem #%v: %v = ", i+1, records[i][0])
		go func() {
			scanner.Scan()
			answerCh <- scanner.Text()
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nTime's up! You scored %d out of %d. \n", grade, len(records))
			return

		case userInput = <-answerCh:
			userInput = strings.TrimSpace(userInput)
			if userInput == records[i][1] {
				grade++
			}
		}
		storedInput = append(storedInput, userInput)
	}

	fmt.Printf("You scored %v out of %v\n", grade, len(records))
}
