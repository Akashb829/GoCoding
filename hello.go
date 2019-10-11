package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", hfunc)
	http.HandleFunc("/add", add)
	http.HandleFunc("/subtract", sub)
	http.HandleFunc("/divide", div)
	http.HandleFunc("/multiply", mult)
	http.ListenAndServe(":8080", nil)
}
func hfunc(w http.ResponseWriter, r *http.Request) {
}

func add(w http.ResponseWriter, r *http.Request) {

	//Getting request
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	//Getting the URL parameters
	x := r.URL.Query()

	//Getting values corresponding to keys
	a := x.Get("a")
	b := x.Get("b")

	//Converting string values to int
	i1, er1 := strconv.Atoi(a)
	if er1 != nil {
		fmt.Println(er1)
	}
	i2, er2 := strconv.Atoi(b)
	if er2 != nil {
		fmt.Println(er1)
	}

	//Adding
	fmt.Println(i1, i2)
	sum := i1 + i2
	fmt.Fprintf(w, "%v", sum)
}

func sub(w http.ResponseWriter, r *http.Request) {

	//Getting request
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	//Getting the URL parameters
	x := r.URL.Query()

	//Getting values corresponding to keys
	a := x.Get("a")
	b := x.Get("b")

	//Converting string values to int
	i1, er1 := strconv.Atoi(a)
	if er1 != nil {
		fmt.Println(er1)
	}
	i2, er2 := strconv.Atoi(b)
	if er2 != nil {
		fmt.Println(er1)
	}

	//Subtracting
	fmt.Println(i1, i2)
	diff := i1 - i2
	fmt.Fprintf(w, "%v", diff)
}
func div(w http.ResponseWriter, r *http.Request) {

	//Getting request
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	//Getting the URL parameters
	x := r.URL.Query()

	//Getting values corresponding to keys
	a := x.Get("a")
	b := x.Get("b")

	//Converting string values to int
	i1, er1 := strconv.Atoi(a)
	if er1 != nil {
		fmt.Println(er1)
	}
	i2, er2 := strconv.Atoi(b)
	if er2 != nil {
		fmt.Println(er1)
	}

	//Dividing
	fmt.Println(i1, i2)
	divi := i1 / i2
	fmt.Fprintf(w, "%v", divi)
}
func mult(w http.ResponseWriter, r *http.Request) {

	//Getting request
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	//Getting the URL parameters
	x := r.URL.Query()

	//Getting values corresponding to keys
	a := x.Get("a")
	b := x.Get("b")

	//Converting string values to int
	i1, er1 := strconv.Atoi(a)
	if er1 != nil {
		fmt.Println(er1)
	}
	i2, er2 := strconv.Atoi(b)
	if er2 != nil {
		fmt.Println(er1)
	}

	//Multiplying
	fmt.Println(i1, i2)
	prod := i1 * i2
	fmt.Fprintf(w, "%v", prod)
}
