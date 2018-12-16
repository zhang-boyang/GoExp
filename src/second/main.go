package main

import (
	"fmt"
)

func AddMap(k, v string) map[string]interface{} {
	var retMap map[string]interface{}
	retMap = make(map[string]interface{})
	retMap[k] = v
	return retMap
}

func AddMapTwo(k, k2, v, v2 string) map[string]interface{} {
	var retMap map[string]interface{}
	retMap = make(map[string]interface{})
	retMap[k] = v
	retMap[k2] = v2
	return retMap
}

//MyExample my example
type MyExample struct {
	Name string
	Age  int
	City string
}

//DeriveMyExample derive from my example
type DeriveMyExample struct {
	MyExample
	Gender string
}

//NoChangeParam no pointer
func NoChangeParam(st MyExample) {
	st.City = "Toukyou"
}

//ChangeParam pointer
func ChangeParam(st *MyExample) {
	st.City = "KyouTo"
}

//SetAttAge func
func (m *MyExample) SetAttAge(n int) {
	m.Age = n
}

func main() {
	var (
		NumberList   [3]string
		NoNumberList []string
	)
	NumberList = [3]string{"ni", "hao", "hello"}
	NoNumberList = []string{"ni", "hao", "hello"}
	fmt.Print(NumberList)
	fmt.Println(NumberList)
	fmt.Print(NoNumberList[1:])
	NoNumberList = append(NoNumberList, "add new")
	fmt.Print(NumberList, " ", len(NumberList), " ", cap(NumberList), "\n")
	fmt.Print(NoNumberList, " ", len(NoNumberList), " ", cap(NoNumberList), "\n")

	for idx, value := range NumberList {
		fmt.Println(idx, value)
	}

	fmt.Println(AddMap("go", "name"))
	fmt.Println(AddMapTwo("go", "go2", "name", "name2"))
	getMap := AddMapTwo("go", "go2", "name", "name2")
	for k, v := range getMap {
		fmt.Println(k, v)
	}

	var MyStruct MyExample
	MyStruct.Age = 10
	MyStruct.City = "Yokohama"
	MyStruct.Name = "Yoshizawa"
	fmt.Println(MyStruct)

	NoChangeParam(MyStruct)
	fmt.Println(MyStruct)
	ChangeParam(&MyStruct)
	fmt.Println(MyStruct)

	MyStruct.SetAttAge(11)
	fmt.Println(MyStruct)

	DMyStuct := DeriveMyExample{MyExample{"Kobe", 10, "AnPai"}, "man"}
	DMyStuct.SetAttAge(22)
	fmt.Println(DMyStuct)

	lambdaFunc := func(a, b int) int { return a + b }
	fmt.Println(lambdaFunc(10, 11))

	var (
		vV   int   = 42
		vVp  *int  = &vV
		vVpp **int = &vVp
	)

	fmt.Println(vV, &vV, *vVp, vVp, &vVp, **vVpp, *vVpp, vVpp)

	var B byte = 'Q'
	fmt.Println(B)

}
