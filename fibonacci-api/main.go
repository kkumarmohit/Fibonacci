// Copyright: 2024 Arm Ltd. All Rights Reserved.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// FibonacciResponse represents the response to a request for a specific order of the
// Fibonacci sequence. For example a request for the 4th order of the sequence would
// have Order:4, Value: 3.
type FibonacciResponse struct {
	Order int `json:"order"`
	Value int `json:"value"`
}

// Function type for Fibonacci calculation methods
type FibonacciFunc func(int) int

func slowFibonacciHandler(method FibonacciFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Get the requested order from the URL query string (i.e. http://localhost:8080/fibonacci?order=4)
		queryParams := req.URL.Query()
		orderStr := queryParams.Get("order")
		// Error checking - has it been provided? Is it a positive integer?
		if orderStr == "" {
			http.Error(w, "You must provide a series order with your request via the 'order' query parameter", 400)
			return
		}
		orderInt, err := strconv.Atoi(orderStr)
		if err != nil || orderInt < 1 {
			http.Error(w, fmt.Sprintf("Provided order '%s' not a positive integer", orderStr), 400)
			return
		}
		// Prepare response
		value := method(orderInt)
		fibResp := FibonacciResponse{
			Order: orderInt,
			Value: value,
		}
		jsonBytes, _ := json.Marshal(&fibResp)
		fmt.Fprint(w, string(jsonBytes))
	}
}

// Return the value of the Fibonacci sequence for the given order, with 1 being the first order.
// For example, the 1st and second orders of the sequence are '1' and the 4th order is '3'.
func fibonacci(order int) int {
	// Print for logs
	fmt.Printf("Calculating Fibonacci iteratively for order %d\n", order)
	if order == 1 {
		return 1
	}
	previous := 0
	current := 1
	for i := 1; i < order; i++ {
		new := previous + current
		previous = current
		current = new
	}
	return current
}

func fibonacciRecursive(order int) int {
	fmt.Printf("Calculating Fibonacci recursively for order %d\n", order)
	memo := make(map[int]int)
	var fib func(int) int
	fib = func(n int) int {
		if n <= 2 {
			return 1
		}
		if val, ok := memo[n]; ok {
			return val
		}
		memo[n] = fib(n-1) + fib(n-2)
		return memo[n]
	}
	return fib(order)
}

func main() {
	http.HandleFunc("/fibonacci", slowFibonacciHandler(fibonacci))
	http.HandleFunc("/recursive-fibonacci", slowFibonacciHandler(fibonacciRecursive))
	http.ListenAndServe(":8080", nil)
}
