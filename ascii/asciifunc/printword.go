package asciifunc

import (
	"strings"
)

func PrintWord(input string) string {
	flag := true
	for i := 0; i < len(input); i++ {
		flag = true
		if input[i] != '\n' {
			flag = false
			i++
			break
		}
	}
	// Splites string to slice of string
	wordSlice := strings.Split(input, "\n")
	str := "" // final string
	for _, v := range wordSlice {
		str += returnAsciiStr(v)
		str += "\n"
	}
	if flag {
		str = strings.TrimSuffix(str, "\n")
		str = strings.TrimSuffix(str, "\n")
		// fmt.Println(str)
		return ""
	}
	str = strings.TrimSuffix(str, "\n")
	return str
}

func returnAsciiStr(input string) string {
	// Creates doulbe slice, where each embedded slice is letter([]strings)
	asciiWord := [][]string{}
	for _, l := range input {
		asciiWord = append(asciiWord, strings.Split(Dict[byte(l)], "\n"))
	}

	// Checks if width of ascii representation could be placed in terminal
	strFirstRow := ""
	for _, l := range asciiWord {
		strFirstRow += l[1]
	}

	str := "" // string represents input input in ascii template
	for i := 0; i < 8; i++ {
		for _, l := range asciiWord {
			str += l[i]
		}
		str += "\n"
	}
	str = strings.Trim(str, "\n")
	return str
}
