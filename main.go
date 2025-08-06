package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Sample log file path
	// Change this to user input later?
	logFile := "example.auth.log"

	file, err := os.Open(logFile)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line) //Print each line (just to test reading for now)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file %v", err)
	}
}
