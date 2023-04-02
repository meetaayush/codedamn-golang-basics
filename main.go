package main

import "fmt"

type Circle struct {
	radius float64
}

func main() {
	c := Circle{radius: 5.365}
	fmt.Println(c.area())
}
