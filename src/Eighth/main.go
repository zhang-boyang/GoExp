package main

import (
	"fmt"
	"os"
)

type error interface {
	Error() string
}

type SomeError struct {
	Reason string
}

func (s SomeError) Error() string {
	fmt.Println("Some Error")
	return s.Reason
}

func main() {
	var erre error = SomeError{"Some Happend"}
	fmt.Println(erre)

	filename := "main.go"
	var f, err = os.Open(filename)

	if err != nil {
		fmt.Printf("%s open error! err:%s\n", filename, err.Error())
		return
	}

	defer f.Close()

	/*
		if change buf to below (slice) Read while
		var buf = make([]byte, 0, 100)
	*/
	var buf = make([]byte, 100)
	var content = []byte{}

	for {
		n, err := f.Read(buf)
		if n > 0 {
			content = append(content, buf[:n]...)
		} else if n == 0 {
			fmt.Println("Done")
			break
		} else if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
	fmt.Printf("%s\n", string(content))

	/*
		goroutine 1 [running]:
		main.JudgeNumber(...)
		/Users/zhangby/Documents/coding/go/GoExp/src/Eighth/main.go:65
	*/
	defer func() {
		errr := recover()
		if errr != nil {
			fmt.Println("error catched", errr)
			if errr == negErr {
				fmt.Println("equal")
			}
		}
	}()
	//fmt.Println(JudgeNumber(-42))
	fmt.Println(JudgeNumber(42))

	defer func() {
		fmt.Println("first defer func1, but second executed")
	}()

	defer func() {
		fmt.Println("second defer func2, but first executed")
	}()

}

var negErr = fmt.Errorf("non neg number")

func JudgeNumber(a int) int {
	if a < 0 {
		panic(negErr)
	} else {
		return 100
	}
}
