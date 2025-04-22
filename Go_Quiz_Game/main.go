package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	// --- Opening the file
	fi, err := os.Open("problems.csv")
	if err != nil {
		panic(err)
	}

	// --- Defer func to close the file when the program ends
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
		fmt.Println("File closed.")
	}()

	// --- Getting a string slice divided by line from the CSV
	grossFile := readCSV(fi)

	// --- Getting the map Test, k is question, v is answer
	mapTest := splitKeyValue(grossFile)

	// --- Requesting user answers
	fmt.Println("You have 15 seconds, try to think fast!")
	answers := printingTest(mapTest)

	// --- Getting the result
	comparingResults(mapTest, answers)

}

// --- Creating a scanner to read the CSV line by line and writing it to a string slice

func readCSV(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var grossFile []string

	for scanner.Scan() {
		grossFile = append(grossFile, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return grossFile
}

// --- Spliting key - value from the grossFile slice. Ex Key - [5+5] Value - [10], divide by ','

func splitKeyValue(grossFile []string) map[string]string {
	formatedFileMap := make(map[string]string)

	for i := range grossFile {
		grossLine := strings.Split(grossFile[i], ",")
		formatedFileMap[grossLine[0]] = grossLine[1]
	}
	return formatedFileMap
}

// --- Printing Keys to console & receiving user answers

func printingTest(fileLines map[string]string) map[string]string {
	reader := bufio.NewReader(os.Stdin)
	userAnswers := make(map[string]string)
	i := 1
	finish := time.Now().Add(time.Second * 15)
	for k := range fileLines {
		if time.Now().After(finish) {
			fmt.Println("NO MORE TIME!")
			break
		}
		fmt.Printf("Problem #%d - %s:", i, k)
		answer, err := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		if err != nil {
			fmt.Println("Error", err)
		}

		userAnswers[k] = answer
		i++
	}
	return userAnswers
}

// --- Comparing Results

func comparingResults(fileLines map[string]string, answers map[string]string) {

	var correctAnswers int
	for k := range fileLines {
		if fileLines[k] == answers[k] {
			correctAnswers++
		}
	}
	fmt.Printf("Total correct answers: %d/%d\n", correctAnswers, len(fileLines))
}
