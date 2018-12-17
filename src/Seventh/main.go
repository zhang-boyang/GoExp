package main

import (
	"fmt"
)

type Smellable interface {
	Smell()
}
type Eatable interface {
	Eat()
}

type Flower struct{}
type Apple struct{}

func (f Flower) Smell() {
	fmt.Println("flower smell good!")
}

func (a Apple) Smell() {
	fmt.Println("apple smell good!")
}

func (a Apple) Eat() {
	fmt.Println("apple eat good")
}

func main() {
	var S1 Smellable
	var S2 Eatable

	var a Apple = Apple{}
	var f Flower = Flower{}

	S1 = f
	S2 = a

	S1.Smell()
	S2.Eat()

	S1 = a
	S1.Smell()
	/* interface is like virtual function in C++, but C++'s interface is in Class,
	Go is a individual ont in Class(or struct)*/

	//empty interface

	BigMap := map[string]interface{}{
		"age":   12,
		"city":  "Yokohama",
		"Adult": false,
	}

	fmt.Println(BigMap)

	{
		var age = BigMap["age"].(int)
		var city = BigMap["city"].(string)
		var adult = BigMap["Adult"].(bool)
		fmt.Println(age, city, adult)
	}

	var p F
	p.Name = "peach"
	p.Fruitable = Peach{}

	var b Banana
	b.Name = "banana"
	b.Fruitable = Banana{}

	fmt.Println(p, b)
	p.want()
	b.want()

	var Ia []Fruitable
	Ia = append(Ia, b)
	Ia = append(Ia, p)

	for i := 0; i < len(Ia); i++ {
		Ia[i].eat()
	}

}

type Fruitable interface {
	eat()
}

type Fruit struct {
	Name string
	Fruitable
}

func (f Fruit) want() {
	fmt.Print("I want to")
	f.eat()
}

type Peach struct {
	Fruit
}

type Banana struct {
	Fruit
}

func (p Peach) eat() {
	fmt.Println(" eat peach")
}

func (b Banana) eat() {
	fmt.Println(" eat banana")
}
func (b Banana) peal() {
	fmt.Println("peal banana")
}
