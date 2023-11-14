package main

import (
	"fmt"
	"github.com/tumivn/goblog/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		panic(err)
	}
}
