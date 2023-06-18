package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	var input, url string

	flag.StringVar(&input, "input", "abcabcbb", "your input string")
	flag.StringVar(&url, "url", "http://localhost:8080/api/substring", "url to send your string to")

	flag.Parse()

	fmt.Println(input, url)

	out, err := sendToServer(input, url)
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
}

func sendToServer(input, url string) (string, error) {
	var sendURL string

	if url != "" {
		sendURL = fmt.Sprintf("%s?input=%s", url, input)
		log.Println(sendURL)
		res, err := http.Get(sendURL)
		if err != nil {
			return "", err
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return "", err
		}
		res.Body.Close()

		log.Println(string(body))

		return string(body), nil
	}

	return "", fmt.Errorf("url cannot be empty")
}
