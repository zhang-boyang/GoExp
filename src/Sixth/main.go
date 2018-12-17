package main

import (
	"fmt"
	"math"
	"unsafe"
)

type Circle struct {
	x      int
	y      int
	Radius int
}

func (c Circle) GetArea() float32 {
	return float32(c.Radius) * float32(c.Radius) * math.Pi
}
func (c Circle) Expand() {
	c.Radius *= 2
}

func (c *Circle) Expand2() {
	c.Radius *= 2
}

type ArrayStruct struct {
	array [10]int
}
type SliceStruct struct {
	slice []int
}

func main() {
	var c Circle = Circle{
		x:      100,
		y:      100,
		Radius: 2,
	}
	fmt.Println(c)

	var c2 Circle = Circle{
		Radius: 3,
	}
	fmt.Printf("%+v\n", c2)

	var c3 Circle = Circle{3, 3, 5}
	fmt.Printf("%+v\n", c3)

	fmt.Println(c3.GetArea())
	c4 := new(Circle)
	fmt.Printf("%+v\n", c4)

	var c5 *Circle = nil
	fmt.Println(unsafe.Sizeof(c5))
	var c6 Circle
	fmt.Println(unsafe.Sizeof(c6))

	var a ArrayStruct = ArrayStruct{
		[...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	var b SliceStruct = SliceStruct{
		[]int{1, 2, 3, 4, 5},
	}

	fmt.Println(a, unsafe.Sizeof(a))
	fmt.Println(b, unsafe.Sizeof(b))

	c3.Expand()
	fmt.Println(c3.Radius)
	c3.Expand2()
	fmt.Println(c3.Radius)

	var CD Ci = Ci{
		Point: Point{
			x: 10,
			y: 10,
		},
		Radius: 5,
	}

	fmt.Println(CD.Radius, CD.Point.x, CD.y)
	CD.show()

	var F Apple
	F.eat()
	F.enjoy() //there is not something like C++ virtual or polymorphic in Go
}

type Point struct {
	x int
	y int
}

func (p Point) show() {
	fmt.Println(p.x, p.y)
}

type Ci struct {
	Point
	Radius int
}

type Fruit struct {
}

func (f Fruit) eat() {
	fmt.Println("eat Fruit")
}

func (f Fruit) enjoy() {
	fmt.Println("enjoy")
	f.eat()
}

type Apple struct {
	Fruit
	s string
}

func (a Apple) eat() {
	fmt.Println("eat Apple")
}
