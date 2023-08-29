package runner

import (
	"fmt"
	"math"
)

func runInterface() {
	printStart()
	stuff := []shape{
		circle{1},
		rectangle{1, 1},
	}

	for _, val := range stuff {
		printShape(val)
	}
	printEnd()
}

type shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width  float64
	height float64
}

func (c circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return r.width * r.height * 2
}

func printShape(val shape) {
	var whoami string
	switch val.(type) {
	case rectangle:
		whoami = "rectangle"
		break
	case circle:
		whoami = "circle"
		break
	default:
		whoami = "anonymous"
	}

	fmt.Printf("im a %s, ", whoami)
	fmt.Printf("per = %f, ", val.perimeter())
	fmt.Printf("area = %f\n", val.area())
}
