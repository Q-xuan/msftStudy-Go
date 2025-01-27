package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type customWriter struct{}

type GithubResponse []struct {
	FullName string `json:"full_name"`
}

func (w customWriter) Write(p []byte) (n int, err error) {
	var resp GithubResponse
	json.Unmarshal(p, &resp)
	for _, r := range resp {
		fmt.Println(r.FullName)
	}
	return len(p), nil
}

func main() {
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		log.Fatal(err)
	}
	writer := customWriter{}
	io.Copy(writer, resp.Body)
}
