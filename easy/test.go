package easy

import (
	"bytes"
	"fmt"
	"strings"
)

func Test() {
	strings := []string{"aaabbbccaaaaaaa”", "abbbccc", "afbfbbcfcc"}

	for _, value := range strings {
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

	fmt.Println(chars)

	// TODO: Do a single loop, if matches fail, reverse i + 1 to where we started last
	matchCount := 0
	latestChar := ""
	var tmp bytes.Buffer
	for i := 0; i <= n-1; i++ {
		if latestChar == chars[i] {
			continue
		}
		latestChar = chars[i]

		for j := i; j <= n-1; j++ {
			if chars[i] == chars[j] {
				matchCount++
				tmp.WriteString(chars[i])
			} else {
				break
			}

			if limit == matchCount {
				buffer.WriteString(tmp.String())
				break
			}
		}
		tmp.Reset()
		matchCount = 0
	}

	return buffer.String()
}
