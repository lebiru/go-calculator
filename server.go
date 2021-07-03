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

func multiplyHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err:%v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful \n")
	multiplier, errOne := strconv.Atoi(r.FormValue("multiplier"))
	if errOne != nil {
		//handle error
	}
	multiplicand, errTwo := strconv.Atoi(r.FormValue("multiplicand"))
	if errTwo != nil {
		//handle error
	}

	fmt.Fprintf(w, "Result of Multiplication = %d\n", multiplier*multiplicand)

}

func divisionHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err:%v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful \n")
	dividend, errOne := strconv.Atoi(r.FormValue("dividend"))
	if errOne != nil {
		//handle error
	}
	divisor, errTwo := strconv.Atoi(r.FormValue("divisor"))
	if errTwo != nil {
		//handle error
	}

	if divisor == 0 {
		fmt.Fprintf(w, "Cannot divide by zero")
		return
	}

	fmt.Fprintf(w, "Result of Division = %d\n", dividend/divisor)

}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/subtract", subtractHandler)
	http.HandleFunc("/multiply", multiplyHandler)
	http.HandleFunc("/divide", divisionHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
