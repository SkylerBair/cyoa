package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"skyler/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the Cyoa web application on")
	filename := flag.String("file", "gopher.json", "the JSON file witht the CYOA story")
	flag.Parse()
	fmt.Printf("using the story in %s. \n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
