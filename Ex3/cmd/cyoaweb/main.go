package main

import (
	"flag"
	"fmt"
	cyoa "gophercises/Ex3"
	"os"
)

func main() {
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

	fmt.Printf("%+v\n", story)
}
