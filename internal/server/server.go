package server

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetSubstring(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got request")
	if r.URL.Query().Has("input") {
		input := r.URL.Query().Get("input")

		output := longestSubstring(input)

		_, err := io.WriteString(w, output)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	http.Error(w, "must have 'input' query parameter", http.StatusBadRequest)
}

func longestSubstring(input string) string {
	output, res := "", ""
	start, maxLen := 0, 0

	for _, char := range input {
		if !strings.Contains(output, string(char)) {
			output += string(char)

			if len(output) > maxLen {
				maxLen = len(output)
				res = output
			}
		} else {
			start = strings.Index(output, string(char)) + 1
			output = output[start:]
			output += string(char)

			if len(output) > maxLen {
				maxLen = len(output)
				res = output
			}
		}
	}

	return res
}
