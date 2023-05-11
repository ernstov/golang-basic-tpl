package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// Environment variables can be read this way:
	ExampleVar1 := os.Getenv("EXAMPLE_VAR")
	ExampleVar2 := os.Getenv("EXAMPLE_VAR_2")

	// Do something with the values

	log.Println("Looking for EXAMPLE_VAR and EXAMPLE_VAR_2 environment variables:")
	log.Println("Example var 1: ", ExampleVar1)
	log.Println("Example var 2: ", ExampleVar2)
	log.Println()
	log.Println("Visit web interface to see more details.")

	// In this example we will show all environment variables
	// We can obtain all environment variables as a []string
	allVars := os.Environ()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<h2>This is an example of a GO application that displays the value of environment variables on a web server</h2>")
		for _, value := range allVars {
			fmt.Fprintf(w, "%s<br>", value)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
