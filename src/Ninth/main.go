package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Print Start To Run")
	go func() {
		fmt.Println("Run first class in goroutine")
		go func() {
			fmt.Println("Run second class in goroutine")
			go func() {
				fmt.Println("Run Thirs class in goroutine")
			}()
		}()
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Start Parallel")
	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("Parallel 1")
	}()
	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("Parallel 2")
	}()
	time.Sleep(time.Second * 3) //waiting for coruntine, P1 never show
	fmt.Println("main Run end")

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("get panic")
			}
		}()
		panic("I do it on purpose")
	}()
	fmt.Println("Wait Panic")
	time.Sleep(time.Second * 3)

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second / 10)
		fmt.Println("run in main fun", i)
	}

	fmt.Println(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(21)
	fmt.Println(runtime.GOMAXPROCS(0))
	//goR()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second / 10)
		fmt.Println("run in main fun", i)
		fmt.Println(runtime.NumGoroutine())
	}

	// Coroutine is so light that it can be
	// created a lot in a process
	for i := 0; i >= 0; i++ {
		go func() {
			for {
				time.Sleep(time.Second)
			}
		}()
		if i%1000 == 0 {
			fmt.Println("now", runtime.NumGoroutine())
		}
	}
}

func goR() {
	n := 15 //when n > runtime.GOMAXPROCS(0) - 1 "run in main fun" will not be run
	//(expriment on 12 cores machine)
	for ; n > 0; n-- {
		go func() {
			fmt.Println("Run in loop", n)
			for {
			}
		}()
	}
}
