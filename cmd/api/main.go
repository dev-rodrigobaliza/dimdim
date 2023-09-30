package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pborman/uuid"
)

// can be overriden with -ldflags
var Version = "dev"
var BuiltBy = "local"

var uuidWithHyphen = uuid.New()

func StartWebServer() {
	http.HandleFunc("/", handleApp)

	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = "8080"
	}

	log.Printf("Starting web server on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func handleApp(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")

	who, ok := r.URL.Query()["who"]
	if ok {
		fmt.Fprintf(w, "Hello, %s\n", who[0])
	} else {
		fmt.Fprintf(w, "Hello, World\n")
	}

	fmt.Fprintf(w, "Version: %s\n", Version)
	fmt.Fprintf(w, "BuiltBy: %s\n", BuiltBy)
}

func main() {
	fmt.Println("starting server...")
	fmt.Println(hello())

	// use 3rd party package as proof-of-concept
	fmt.Println("UUID: ", uuidWithHyphen)
	fmt.Println("Version: ", Version)
	fmt.Println("BuiltBy: ", BuiltBy)

	StartWebServer()
}

func hello() string {
	return "Hello Golang"
}