package struct_method_interfaces

import (
	"math"
)


type Rectangle struct {

	Height float64
	Width float64

}

type Shape interface {
    Area() float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rectangle Rectangle) float64{

	return 2*(rectangle.Height+rectangle.Width)
}


func (rectangle *Rectangle) Area() float64{

	return rectangle.Height*rectangle.Width
}


func (c *Circle) Area() float64{
    return math.Pi * c.Radius * c.Radius
}