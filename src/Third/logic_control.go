package main

import "fmt"

func main() {

	fmt.Println(sign(max(42, 24), min(24, 42)))
	fmt.Println(Prize1(42), "\n", Prize2(100))

	/*
		for {
			fmt.Println("Hola!!")
		}
	*/
	for i := 10; i > 0; i-- {
		fmt.Print(i)
	}
	fmt.Println()

	var arr [10]int
	/*
		for i := 0; i <= len(arr); i++ {
			fmt.Print(arr[i])
		}
	*/
	for i := len(arr); i > 0; i-- {
		fmt.Print(arr[i-1])
	}
	fmt.Println()

	var a = [10]int{1, 2, 3, 4, 5, 6, 7, 8}
	var b [10]int
	b = a
	a[1] = 1234
	fmt.Println(a, b)

	for idx := range a {
		fmt.Println(idx, a[idx])
	}

	var c []int = make([]int, 5, 8)
	d := make([]int, 5)
	e := []int{1, 2, 3, 4, 5}
	fmt.Println(c, len(c), cap(c))
	fmt.Println(d, len(d), cap(d))
	fmt.Println(e, len(e), cap(e))

	var f []int = []int{1, 2, 3, 4, 5, 6}
	g := f
	f[0] = 42
	fmt.Println(f, len(f), cap(f), &f)
	fmt.Println(g, len(g), cap(g), &g)

	for k, v := range a {
		fmt.Println(k, v)
	}

	for k, v := range f {
		fmt.Println(k, v)
	}

	h := append(f, 9)
	fmt.Println(f, len(f), cap(f))
	fmt.Println(g, len(g), cap(g))
	fmt.Println(h, len(h), cap(h))

	/*
		evaluated but not used
		append(h, 10)
	*/
	_ = append(h, 10)

	fmt.Println(a[2:5])
	fmt.Println(h[2:5])
	iA := h[2:5]
	fmt.Println(h, len(h), cap(h))
	fmt.Println(iA, len(iA), cap(iA))

	h[3] = 10068
	fmt.Println(h, len(h), cap(h))
	fmt.Println(iA, len(iA), cap(iA))

	iB := a[2:5]
	a[3] = 10086
	fmt.Println(iB, len(iB), cap(iB))

	o := []int{1, 2, 3, 4, 5, 6}
	var p []int = make([]int, 2)

	p = append(p, 1)
	fmt.Println(p, o)

	copyN := copy(p, o)
	fmt.Println(copyN, p, o)
	o[0] = 10086
	fmt.Println(p, o)

}

func max(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

func min(a, b uint) uint {
	if a > b {
		return b
	}

	return a
}

func sign(a, b uint) int {
	if a-b > 0 {
		return 1
	} else if a-b == 0 {
		return 0
	}
	return -1
}

func Prize1(score int) string {
	switch score / 10 {
	case 0, 1, 2, 3, 4, 5:
		return "Bad"
	case 6:
		return "Pass"
	case 7, 8:
		return "Good"
	case 9:
		return "Excellent"
	default:
		return "Bravo"
	}
}

func Prize2(score int) string {
	switch {
	case score < 60:
		return "Bad"
	case score < 70:
		return "Pass"
	case score < 80:
		return "Good"
	case score < 100:
		return "Excellent"
	default:
		return "Bravo"
	}
}
