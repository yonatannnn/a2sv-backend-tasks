package main

import (
	"fmt"
	"regexp"
	"strings"
)

func countWords(text string) map[string]int {
	count := make(map[string]int)
	words := strings.Fields(text)

	for _, word := range words {
		cleanedWord := cleanWord(word)
		if cleanedWord != "" {
			count[strings.ToLower(cleanedWord)]++
		}
	}
	return count
}

func cleanWord(word string) string {
	re := regexp.MustCompile(`[^\w\s]`)
	cleanedWord := re.ReplaceAllString(word, "")
	return cleanedWord
}


func checkPalindrome(text string) bool {
	re := regexp.MustCompile(`[^\w]`)
	cleanedText := re.ReplaceAllString(text, "")
	caseInsensitiveText := strings.ToLower(cleanedText)
	var i int = 0;
	var j int = len(caseInsensitiveText) - 1;
	for i < j {
		if caseInsensitiveText[i] != caseInsensitiveText[j] {
			return false
		}
		i++
		j--
	}
	return true
}



func main() {
	fmt.Println(countWords("Hello world! Hello"))
}
