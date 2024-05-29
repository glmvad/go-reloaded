package utils

import (
	"strconv"
	"strings"
)

// ProcessFunction applies various text processing operations to the input words.
func ProcessFunction(words []string) []string {
	
	for i := 0; i < len(words); i++ {
		commandName := ""
		if strings.HasPrefix(words[i], "(") {
			if strings.HasSuffix(words[i], ")") {
				oper := strings.Trim(words[i], "()")
				if strings.Contains(oper, ", ") {
					comm := strings.Split(oper, ", ")
					commandName = comm[0]
					count, _ := strconv.Atoi(comm[1])
					// Call ProcessFunctionWord to handle the specific command.
					words = ProcessFunctionWord(words, commandName, count, i)
					// Adjust the loop index to account for changes in the words slice.
					i--
				} else if oper == "cap" || oper == "up" || oper == "low" || oper == "hex" || oper == "bin" {
					commandName = oper
					words = ProcessFunctionWord(words, commandName, 1, i)
					i--
				}
			}
		}
	}
	// Loop through words again to perform additional processing for articles.
	for i := 0; i < len(words); i++ {
		if (words[i] == "a" || words[i] == "A") && i+1 < len(words) {
			if strings.Contains("aeiouhAEIOUH", string(words[i+1][0])) && len(words[i+1]) > 1 {
				words[i] += "n"
			}
		}
	}
	return words
}
