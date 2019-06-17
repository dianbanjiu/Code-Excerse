package main

import (
	"fmt"
	"math"
	"reflect"
)

type Rectangle struct {
	height, wight float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.height * r.wight
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r := Rectangle{12, 13}
	c := Circle{4}
	fmt.Println(r.area())
	fmt.Println(c.area())

	i := 3
	t := reflect.TypeOf(i)
	fmt.Println(t)
}
