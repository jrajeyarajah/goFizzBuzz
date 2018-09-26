package main

import (
	"strconv"
	"time"
)

func fizzbuzzPrint(number, Fizz, Buzz int) string {
	var output = ""
	if number%Fizz == 0 {
		output = "Fizz"
	}
	if number%Buzz == 0 {
		output += "Buzz"
	}
	if output == "" {
		output = strconv.Itoa(number)
	}
	return output
}

func goThrough(start int, end int) {
	for i := start; i <= end; i++ {
		println(fizzbuzzPrint(i, 5, 7))
	}
}

func main() {
	start := time.Now()
	goThrough(1, 350)
	diff := time.Since(start)
	println(diff)
}
