package helpers

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadInput(input string) string {
	// Open the file
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new scanner
	scanner := bufio.NewScanner(file)

	// Create a string builder to store the result
	var result strings.Builder

	// Read the file line by line
	for scanner.Scan() {
		result.WriteString(scanner.Text())
		result.WriteString("\n")
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result.String()
}
