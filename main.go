package main

import (
	"fmt"

	"github.com/kikiklang/go-playground/even_or_odd"
	"github.com/kikiklang/go-playground/hello"
	"github.com/kikiklang/go-playground/quarter"
	"github.com/kikiklang/go-playground/reversed_string"
)

func main() {
	fmt.Println("**********************************")
	fmt.Println(hello.Hello())

	fmt.Println("**********************************")
	fmt.Println(quarter.QuarterOf(5))

	fmt.Println("**********************************")
	fmt.Println(even_or_odd.EvenOrOdd(5))

	fmt.Println("**********************************")
	fmt.Println(reversed_string.ReversedString("test that shit"))
}
