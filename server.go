package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func addHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err:%v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful \n")
	numOne, errOne := strconv.Atoi(r.FormValue("numOne"))
	if errOne != nil {
		//handle error
	}
	numTwo, errTwo := strconv.Atoi(r.FormValue("numTwo"))
	if errTwo != nil {
		//handle error
	}

	fmt.Fprintf(w, "Result of Addition = %d\n", numOne+numTwo)

}

func subtractHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err:%v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful \n")
	subtrahend, errOne := strconv.Atoi(r.FormValue("subtrahend"))
	if errOne != nil {
		//handle error
	}
	minuend, errTwo := strconv.Atoi(r.FormValue("minuend"))
	if errTwo != nil {
		//handle error
	}

	fmt.Fprintf(w, "Result of Subtraction = %d\n", subtrahend-minuend)

}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/subtract", subtractHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
