package main

import (
	"fmt"
	"time"

	"github.com/kikiklang/go-playground/sum_of_cubes"
)

func main() {
	start := time.Now()
	fmt.Println("**********************************")
	fmt.Println(sum_of_cubes.SumCubes(10000))
	fmt.Println("**********************************")
	fmt.Println(time.Since(start))
}
