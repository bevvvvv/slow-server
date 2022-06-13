package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Computes the nth number in the fibonacci series recursivey
func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func ServeFib(w http.ResponseWriter, r *http.Request) {
	values, ok := r.URL.Query()["n"]
	if len(values) == 0 {
		http.Error(w, "no integer parameter n present", http.StatusBadRequest)
		return
	}
	fmt.Println("Responding to Fibonacci request...")
	if !ok {
		fmt.Println("Error responding to Fibonacci request.")
		http.Error(w, "error parsing integer parameter n", http.StatusBadRequest)
		return
	}
	n, err := strconv.Atoi(values[0])
	if err != nil {
		fmt.Println("Error responding to Fibonacci request.")
		http.Error(w, "error parsing integer parameter n", http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, fmt.Sprintf("Fibonacci number %d is %d", n, fib(n)))
	fmt.Println("Responded to Fibonacci request.")
}

func main() {
	http.HandleFunc("/", ServeFib)

	fmt.Println("Starting web server on port 3000...")
	http.ListenAndServe(":3000", nil)
}
