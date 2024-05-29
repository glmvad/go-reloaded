package main

import (
	"fmt"
	"go-reloaded/utils"
	"os"
	"regexp"
	"strings"
)

func main() {

	// Check if the correct number of arguments are provided.
	if len(os.Args) != 3 {
		fmt.Println("Error(missing text files)")
		os.Exit(1)
	}

	fileSample := os.Args[1]
	fileResult := os.Args[2]
	// Read the content of the input file
	input, err := os.ReadFile(fileSample)
	if err != nil {
		fmt.Println("Error, reading file", err)
		os.Exit(1)
	}
	// Conver the content to a string
	text := string(input)
	// Define a regular expression to get words from the text.
	regex := regexp.MustCompile(`\b\w*-\w*\w*\b|\b\w*'\w*\w*\b|\(\w+\)|\(\w+,\s*\w+\)|\b\w+\b|[^\w\s]|\n`)
	words := regex.FindAllString(text, -1)

	// Loop through the words and replace single quotes with '#' if they are part of a word.
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], "'") && len(words[i]) > 1 {
			words[i] = strings.ReplaceAll(words[i], "'", "#")
		}
	}
	// Process the words using a custom function
	processWord := utils.ProcessFunction(words)
	// Join the processed words into a string.
	resultStr := joinWords(processWord)
	// Create the output file and write the result string to it.
	file, err := os.Create(fileResult)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.WriteString(resultStr)
}

// joinWords concatenates processed words into a string.
func joinWords(words []string) string {

	strRes := ""
	count := 0

	for _, word := range words {
		// Handle punctuation marks and single quotes.
		if word == "!" || word == "?" || word == "." || word == "," || word == ":" || word == ";" {
			strRes = strRes[:len(strRes)-1] + word + " "
		} else if word == "'" {
			count++
			if count == 1 {
				strRes += word
			} else if count == 2 {
				strRes = strRes[:len(strRes)-1] + word + " "
				count = 0
			}
		} else if word == "(" {
			strRes += word
		} else if word == ")" {
			strRes = strRes[:len(strRes)-1] + word + " "
		} else {
			strRes += word + " "
		}
	}
	// Replace '#' with a single quote in the result string.
	if strings.Contains(strRes, "#") {
		strRes = strings.ReplaceAll(strRes, "#", "'")
	}
	strRes = strings.Trim(strRes, " ")
	return strRes
}
