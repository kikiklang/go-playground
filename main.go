package main

import (
	"fmt"
	"time"

	"github.com/kikiklang/go-playground/multiple_of_index"
)

func main() {
	start := time.Now()
	x := []int{22, -6, 32, 82, 9, 25}
	fmt.Println("**********************************")
	fmt.Println(multiple_of_index.MultipleOfIndex(x))
	fmt.Println("**********************************")
	fmt.Println(time.Since(start))
}
