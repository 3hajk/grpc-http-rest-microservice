package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var address = flag.String("server", "http://localhost:8080", "HTTP gateway url, e.g. http://localhost:8080")

func main() {
	// get configuration
	flag.Parse()

	var body string

	// Call Create
	resp, err := http.Post(*address+"/v1/info", "application/json", strings.NewReader(`{"api":"v1"}`))
	if err != nil {
		log.Fatalf("failed to call Create method: %v", err)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Create response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Info response: Code=%d, Body=%s\n", resp.StatusCode, body)
}
