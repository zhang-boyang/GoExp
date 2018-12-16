package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var m map[int]string = make(map[int]string, 10)
	var m2 map[int]string = make(map[int]string)
	println(m, len(m))
	println(m2, len(m2))

	m3 := map[string]int{
		"banana": 1,
		"apple":  2,
		"orange": 3,
	}
	fmt.Println(m3, len(m3))

	var num = m3["banana"]
	fmt.Println(num)

	m3["pear"] = 4
	fmt.Println(m3, len(m3))

	delete(m3, "pear")
	fmt.Println(m3, len(m3))

	var score, ok = m3["pear"]
	fmt.Println(m3, score, ok)

	for name, score := range m3 {
		fmt.Println(name, score)
	}

	for name := range m3 {
		fmt.Println(name)
	}

	t := make([]string, 0)
	fmt.Println(t)
	println(t)
	for name := range m3 {
		println(t)
		t = append(t, name)

	}
	m4 := make(map[string]int)
	println(m4, m3)
	m4 = m3
	println(m4, m3)
	fmt.Println(m4)
	m3["peach"] = 9
	fmt.Println(m4)

	fmt.Println(unsafe.Sizeof(m4))

}
