package main

import "fmt"

type Shape interface{ Area() float64 }
type Rect struct{ W, H float64 }

func (r Rect) Area() float64 { return r.W * r.H }

type Circle struct{ R float64 }

func (c Circle) Area() float64 { return 3.14159 * c.R * c.R }
func Total(a []Shape) float64 {
	t := 0.0
	for _, s := range a {
		t += s.Area()
	}
	return t
}
func main() { fmt.Println(Total([]Shape{Rect{2, 3}, Circle{1}})) }
