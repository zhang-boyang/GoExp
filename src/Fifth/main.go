package main

import "fmt"

func main() {
	var s = "中文golang"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Printf("\n")

	for codepoint, runeValue := range s {
		fmt.Printf("%d-%x ", codepoint, runeValue)
	}
	fmt.Printf("\n")

	s1 := "hola"
	var s2 string
	for i := 0; i < 10; i++ {
		s2 += s1
	}
	fmt.Println(s1)
	fmt.Println(s2)

	//s1[0] = 'a'  cannot assign to s1[0]
	s2 = s1[1:3]
	fmt.Println(s2)

	var bS3 []byte = []byte(s1)
	var s4 string = string(bS3)
	fmt.Println(bS3)
	fmt.Println(s4)
}
