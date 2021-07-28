package main

import (
	"fmt"

	"github.com/kikiklang/go-playground/hello"
	"github.com/kikiklang/go-playground/quarter"
)

func main() {
	fmt.Println(hello.Hello())
	fmt.Println(quarter.QuarterOf(5))
}
