package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)
// It modifies the input words based on the given commandName, count, and current index.
func ProcessFunctionWord(words []string, comandName string, count int, currInd int) []string {
	splitInd := currInd
	currInd--
	punc := ".,?!':;()"
	// Switch statement to handle different command types.
	switch comandName {
	case "cap":
		for currInd >= 0 && count > 0 {
			// Check if the current word is not a space or a punctuation mark.
			if !unicode.IsSpace(rune(words[currInd][0])) && !strings.Contains(punc, string(words[currInd][0])) { 
				words[currInd] = Cap(words[currInd])
				count--
			}
			currInd--
		}
		// Remove the original command word from the words slice.
		return append(words[:splitInd], words[splitInd+1:]...)
	case "up":
		for currInd >= 0 && count > 0 {
			if !unicode.IsSpace(rune(words[currInd][0])) && !strings.Contains(punc, string(words[currInd][0])) { 
				words[currInd] = Up(words[currInd])
				count--

			}
			currInd--
		}
		return append(words[:splitInd], words[splitInd+1:]...)
	case "low":
		for currInd >= 0 && count > 0 {

			if !unicode.IsSpace(rune(words[currInd][0])) && !strings.Contains(punc, string(words[currInd][0])) { 
				words[currInd] = Low(words[currInd])
				count--

			}
			currInd--
		}
		return append(words[:splitInd], words[splitInd+1:]...)
	case "hex":
		for currInd >= 0 && count > 0 {
			if unicode.IsSpace(rune(words[currInd][0])) || strings.Contains(punc, string(words[currInd][0])) { 
				currInd--
				continue
			}
			// Convert the word to hexadecimal and update the words slice.
			words[currInd] = strconv.FormatInt(Hex(words[currInd]), 10)
			currInd--
			count--
		}
		return append(words[:splitInd], words[splitInd+1:]...)
	case "bin":
		for currInd >= 0 && count > 0 {
			if unicode.IsSpace(rune(words[currInd][0])) || strings.Contains(punc, string(words[currInd][0])) { 
				currInd--
				continue
			}
			// Convert the word to binary and update the words slice.
			words[currInd] = strconv.FormatInt(Bin(words[currInd]), 10)
			currInd--
			count--
		}
		return append(words[:splitInd], words[splitInd+1:]...)
	}
	return words
}
// Cap capitalizes the first letter of a word.
func Cap(word string) string {
	res := strings.ToLower(word)
	res = string(unicode.ToUpper(rune(word[0]))) + res[1:]
	return res
}
// Up converts a word to uppercase.
func Up(word string) string {
	res := strings.ToUpper(word)
	return res
}
// Low converts a word to lowercase.
func Low(word string) string {
	res := strings.ToLower(word)
	return res
}
// Hex converts a hexadecimal string to an int64.
func Hex(word string) int64 {
	res, err := strconv.ParseInt(word, 16, 64)
	if err != nil {
		fmt.Println("error")
		os.Exit(1)
	}
	return res
}
// Bin converts a binary string to an int64.
func Bin(word string) int64 {
	res, err := strconv.ParseInt(word, 2, 64)
	if err != nil {
		fmt.Println("error")
		os.Exit(1)
	}
	return res
}
