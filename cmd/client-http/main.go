package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	address = flag.String("server", "rest://localhost:8080", "HTTP gateway url, e.g. rest://localhost:8080")
)

func main() {
	// get configuration
	flag.Parse()

	t := time.Now().In(time.UTC)
	pfx := t.Format(time.RFC3339Nano)

	var body string

	// Call Create
	resp, err := http.Post(*address+"/v1/info", "application/json",
		strings.NewReader(fmt.Sprintf(`{"api":"v1"}`, pfx, pfx, pfx)))
	if err != nil {
		log.Fatalf("failed to call Create method: %v", err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Create response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Create response: Code=%d, Body=%s\n\n", resp.StatusCode, body)
}
