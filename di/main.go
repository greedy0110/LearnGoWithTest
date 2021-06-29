package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, msg string) {
	fmt.Fprintf(writer, "Hello, %s", msg)
}

func MyGreeterHeandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Seungmin")
}

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHeandler)))
}
