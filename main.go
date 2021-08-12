package main

import (
	"fmt"
	"time"

	"github.com/kikiklang/go-playground/sum_of_integers_in_a_string"
)

func main() {
	start := time.Now()
	fmt.Println("**********************************")
	fmt.Println(sum_of_integers_in_a_string.SumOfIntegersInString("The Great Depression lasted from 1929 to 1939."))
	fmt.Println("**********************************")
	fmt.Println(time.Since(start))
}
