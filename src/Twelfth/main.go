package main

import (
	"fmt"
	"unsafe"
)

type Rect struct {
	width  int
	height int
}

type test interface {
	test()
}

type Stest struct {
	a int
}

func main() {
	var r Rect = Rect{50, 60}
	w := *(*int)(unsafe.Pointer(&r))
	h := *(*int)(unsafe.Pointer((uintptr(unsafe.Pointer(&r)) + uintptr(8))))
	fmt.Println(w, h)

	hOffsizeptr := (*int)(unsafe.Pointer((uintptr(unsafe.Pointer(&r)) + unsafe.Offsetof(r.height))))
	hsizeofptr := (*int)(unsafe.Pointer((uintptr(unsafe.Pointer(&r)) + unsafe.Sizeof(r.width))))
	fmt.Println(h, *hOffsizeptr, *hsizeofptr)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	address := (**int)(unsafe.Pointer(&slice))
	len := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(address)) + uintptr(8)))
	cap := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(address)) + uintptr(16)))
	fmt.Println(*len, *cap)

	list := "1234"
	fmt.Println(string2slice(list))
}

func string2slice(list string) []byte {
	/*
		add := *(*[2]int)(unsafe.Pointer(&list))
		fmt.Println(add)
		sliceHead := [3]int{}
		sliceHead[0] = add[0]
		sliceHead[1] = add[1]
		sliceHead[2] = add[1]
		return *(*[]int)(unsafe.Pointer(&sliceHead))
	*/
	address := *(*[2]int)(unsafe.Pointer(&list))
	//fmt.Println(address)
	sliceHead := [3]int{}
	sliceHead[0] = address[0]
	sliceHead[1] = address[1]
	sliceHead[2] = address[1]
	return *(*[]byte)(unsafe.Pointer(&sliceHead))

	/*
		Ts := Stest{a: 1}
		var T test = Ts
		data must implement interface this interface implement (missing test method)
	*/
}
