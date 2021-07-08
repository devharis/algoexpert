package easy

import (
	"bytes"
	"fmt"
	"strings"
)

func Test() {
	strings := []string{"Hello, World!", "Welcome to CoderPad.", "This pad is running Java "}

	for _, value := range strings {
		fmt.Println(value)
		result := remove(value, 1)
		fmt.Println(result)
	}
}

func remove(value string, limit int) string {

	if limit == 0 {
		return ""
	}

	chars := strings.Split(value, "")
	n := len(chars)
	var buffer bytes.Buffer

	current := chars[0]
	previous := ""
	matchedChars := make([]byte, limit)
	matches := 0

	for i := 0; i < n-1; i++ {
		if previous == chars[i] { // TODO: this doesn't work!
			current = chars[i]
			previous = chars[i]
			continue
		}

		if current == chars[i] {
			matches++
			matchedChars = append(matchedChars, chars[i][0])

			if limit == matches {
				buffer.WriteString(string(matchedChars))
				current = chars[i+1]
				previous = chars[i]

				matches = 0
				matchedChars = nil
			}
		} else {
			i -= matches
			current = chars[i]

			matches = 0
			matchedChars = nil
		}
	}

	return buffer.String()
}
