package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "invalid route", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "invalid method", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "HEllo")

}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() %v", err)
		return
	}
	fmt.Fprintf(w, "request posted succensfully")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "name:%v", name)
	fmt.Fprintf(w, "address:%v", address)

}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("listening to server 12000\n")
	if err := http.ListenAndServe(":12000", nil); err != nil {
		log.Fatal(err)
	}
}
