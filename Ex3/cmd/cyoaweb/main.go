package main

import (
	"flag"
	"fmt"
	cyoa "gophercises/Ex3"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "the port on which to start the CYOA web application.")
	filename := flag.String("file", "gopher.json", "JSON file in CYOA format")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JSONStory(f)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server at: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
