package main

import "fmt"

func main() {
	app := NewApp()

	err := app.Router.Run(":8080")
	if err != nil {
		fmt.Printf("Error running the server: %v\n", err)
	}
}
