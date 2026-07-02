package main

import (
	"log"

	"github.com/hynexis/heos-cli/internal/runtime"
)

func main() {
	app, err := runtime.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}