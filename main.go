package main

import (
	"fmt"
	"time"

	"github.com/kikiklang/go-playground/equable_triangle"
)

func main() {
	start := time.Now()
	fmt.Println("**********************************")
	fmt.Println(equable_triangle.EquableTriangle(5, 12, 13))
	fmt.Println("**********************************")
	fmt.Println(time.Since(start))
}
