package main

import (
	"fmt"
	"os"

	"github.com/3hajk/grpc-http-rest-microservice/app"
)

func main() {
	if err := app.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
