package main

import (
	"fmt"
	"time"

	"github.com/kikiklang/go-playground/frequency_sequence"
)

func main() {
	start := time.Now()
	fmt.Println("**********************************")
	fmt.Println(frequency_sequence.FreqSeq("hello world", "-"))
	fmt.Println("**********************************")
	fmt.Println(time.Since(start))
}
