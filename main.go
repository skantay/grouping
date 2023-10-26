package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 2 {

		if isNotValid(args[0]) {
			return
		}

		letters := getLetters(args[0])

		var word string
		result := []string{}

		for i, letter := range args[1] {

			if letter == ' ' || i == len(args[1])-1 {
				times, ok := letterInWord(letters, word)

				if ok {

					if i == len(args[1])-1 {
						word += string(letter)
					}

					for times != 0 {
						result = append(result, word)
						times--
					}
				}

				word = ""
				continue
			} else if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {
				word += string(letter)
			}

		}

		for _, v := range result {
			fmt.Println(v)
		}
	}
}

func letterInWord(letters []string, word string) (int, bool) {

	var result int
	var isTrue bool



	for _, letter := range letters {

		if len(letter) > len(word) {
			continue
		}

		for i := 0; i <= len(word)-len(letter); i++ {
			if letter == word[i:i+len(letter)] {
				result++
				isTrue = true
				break
			}
		}

	}

	if isTrue {
		return result, true
	}
	return result, false
}

func getLetters(arg string) []string {

	result := []string{}
	var word string

	for _, v := range arg {

		if v == '(' {
			continue
		} else if v == ')' || v == '|' {
			result = append(result, word)
			word = ""
			continue
		}

		word += string(v)

	}
	return result
}

func isNotValid(arg string) bool {
	words := []rune(arg)
	if words[0] == '(' && words[len(words)-1] == ')' {
		return false
	}

	return true
}
