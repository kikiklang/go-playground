package main

import (
	"fmt"
	"time"

	"github.com/kikiklang/go-playground/over_the_road"
)

func main() {
	start := time.Now()
	fmt.Println("**********************************")
	fmt.Println(over_the_road.OverTheRoad(1, 3))
	fmt.Println("**********************************")
	fmt.Println(time.Since(start))
}
