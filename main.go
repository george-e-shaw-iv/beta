package main

import (
	"fmt"

	"github.com/george-e-shaw-iv/go-webapp-boilerplates/pkg/application"
)

func main() {
	app := application.New("application/", "document_root/", 3000)

	fmt.Println("CTRL+C to stop application")

	err := app.Start()
	if err != nil {
		fmt.Errorf("Error starting application: %v\n", err)
	}
}
