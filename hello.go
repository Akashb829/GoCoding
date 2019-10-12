//GO CALCULATOR
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

	var sum int
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	for key, value := range r.Form {
		fmt.Printf("Key = %v | Value = %v", key, value[0])
		i1, er1 := strconv.Atoi(value[0])
		if er1 != nil {
			fmt.Println(er1)
		}
		sum += i1
	}
	fmt.Fprintf(w, "%v", sum)
}

func sub(w http.ResponseWriter, r *http.Request) {

	var diff int
	var vals [2]string
	var val [2]int
	i := 0
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	for key, value := range r.Form {
		fmt.Printf("Key = %v | Value = %v", key, value[0])
		vals[i] = value[0]
		i1, er1 := strconv.Atoi(vals[i])
		if er1 != nil {
			fmt.Println(er1)
		}
		val[i] = i1
		i++
	}
	diff = val[0] - val[1]
	fmt.Fprintf(w, "%v", diff)
}
func div(w http.ResponseWriter, r *http.Request) {

	var divi int
	var vals [2]string
	var val [2]int
	i := 0
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	for key, value := range r.Form {
		fmt.Printf("Key = %v | Value = %v", key, value[0])
		vals[i] = value[0]
		i1, er1 := strconv.Atoi(vals[i])
		if er1 != nil {
			fmt.Println(er1)
		}
		val[i] = i1
		i++
	}
	divi = val[0] / val[1]
	fmt.Fprintf(w, "%v", divi)
}
func mult(w http.ResponseWriter, r *http.Request) {

	var prod int
	var vals [2]string
	var val [2]int
	i := 0
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	for key, value := range r.Form {
		fmt.Printf("Key = %v | Value = %v", key, value[0])
		vals[i] = value[0]
		i1, er1 := strconv.Atoi(vals[i])
		if er1 != nil {
			fmt.Println(er1)
		}
		val[i] = i1
		i++
	}
	prod = val[1] * val[0]
	fmt.Fprintf(w, "%v", prod)
}
