package main

import (
	"fmt"
	"time"

	"github.com/kikiklang/go-playground/performance_check"
)

func main() {
	start := time.Now()
	fmt.Println("**********************************")
	fmt.Println(performance_check.Format5000Name(5000))
	fmt.Println("**********************************")
	fmt.Println(time.Since(start))
}
