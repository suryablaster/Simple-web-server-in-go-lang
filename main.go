package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/form.html")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func main() {

	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("server is started at 8000 port \n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}
