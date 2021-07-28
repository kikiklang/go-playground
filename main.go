package main

import (
	"fmt"
	"time"

	"github.com/kikiklang/go-playground/maximum_length_difference"
)

func main() {
	start := time.Now()
	a1 := []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	a2 := []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
	fmt.Println("**********************************")
	fmt.Println(maximum_length_difference.MxDifLg(a1, a2))
	fmt.Println("**********************************")
	fmt.Println(time.Since(start))
}
