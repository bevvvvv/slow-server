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
	if !ok {
		http.Error(w, "error parsing integer parameter n", http.StatusBadRequest)
		return
	}
	n, err := strconv.Atoi(values[0])
	if err != nil {
		http.Error(w, "error parsing integer parameter n", http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, fmt.Sprintf("Fibonacci number %d is %d", n, fib(n)))
}

func main() {
	http.HandleFunc("/", ServeFib)

	http.ListenAndServe(":3000", nil)
}
