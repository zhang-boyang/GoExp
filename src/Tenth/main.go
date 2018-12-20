package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 10)
	for i := 0; i < cap(ch); i++ {
		ch <- i
	}

	var list [10]int
	for i := 0; i < cap(ch); i++ {
		list[i] = <-ch
	}
	fmt.Println(list)

	NonBufCh := make(chan int)
	go func() {
		var Rnum int = rand.Intn(100)
		NonBufCh <- Rnum
		fmt.Println("send non buf channel")
	}()

	time.Sleep(time.Second)
	GetInt := <-NonBufCh
	fmt.Println("Get Non Buf channel", GetInt)
	time.Sleep(time.Second / 10)

	CloseCh := make(chan int, 4)
	CloseCh <- 42
	CloseCh <- 24
	close(CloseCh)

	va1 := <-CloseCh
	va2 := <-CloseCh
	fmt.Println(va1, va2)
	//CloseCh <- 33   panic: send on closed channe

	ch <- 1
	ch <- 2
	close(ch)
	for v := range ch {
		fmt.Println(v)
		// if not close the ch, main routine will be sleep, it becomes a deadlock
	}

	/*
		Because there is not a method to test whether ch is closed,
		so the reader do not close channel, it should be closed by writer
		If writer do not use this channel anymore, close it!!
	*/

	ShutCh := make(chan int, 4)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go WriteWithWg(ShutCh, wg)
	go WriteWithWg(ShutCh, wg)

	fmt.Println("Wait for shutting")
	go func() {
		for v := range ShutCh {
			println(v)
		}
	}()
	wg.Wait()
	close(ShutCh)

	ch1 := make(chan int, 10)
	go send(ch1, 1)
	ch2 := make(chan int, 10)
	go send(ch2, 2)
	ch3 := make(chan int, 10)
	go send(ch3, 3)

	chclt := make(chan int, 10)
	go collect(chclt, ch1)
	go collect(chclt, ch2)
	go collect(chclt, ch3)

	go recv(chclt)

	time.Sleep(time.Second * 5)
	fmt.Printf("\n")

	Ch1 := make(chan int, 10)
	go send(Ch1, 1)
	Ch2 := make(chan int, 10)
	go send(Ch2, 2)
	Ch3 := make(chan int, 10)
	go send(Ch3, 3)

	//multiRecv(Ch1, Ch2, Ch3)
	time.Sleep(time.Second * 5)
	fmt.Printf("\n")

	CH1 := make(chan int, 10)
	CH2 := make(chan int, 10)
	CH3 := make(chan int, 10)
	/*
		with default:
		0 1 3 11 12 16 4 5 2 9 13 19 15 6 17 18 25 21 22 27
		826505 23 29 2492113 7 30 8 2865897 458466 1334731
		10 14 1704865 20 3606222 24 39

		without default:
		0 1 7 9 2 3 10 11 12 15 4 18 5 13
		19 6 22 8 23 20 14 26 27 25 30 32
		33 39 41 16 28 29 44 17 31 21 46 24 35 40 34
		never miss single one
	*/
	go SendWithoutBlocking(CH1, CH2, CH3)

	sleepRecv(CH1, CH2, CH3)
	time.Sleep(time.Second * 5)
	fmt.Printf("\n")

}

func WriteWithWg(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() //make sure it will be execute in the last
	for i := 0; i < 4; i++ {
		ch <- i
		time.Sleep(time.Second / 10)
	}
}

func send(ch chan int, no int) {
	defer close(ch)
	for i := 0; i < 20; i++ {
		ch <- (no*100 + i)
		time.Sleep(time.Second / 10)
	}
}
func collect(clt chan int, ch chan int) {
	defer close(ch)
	for {
		for v := range ch {
			clt <- v
		}
	}

}

func recv(ch chan int) {
	for v := range ch {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n") // I do not know why the line break doesnot show
}

func multiRecv(ch1, ch2, ch3 chan int) {
	for {
		select { //It won't stop when channels are already closed
		case v := <-ch1:
			fmt.Printf("%d ", v)
		case v := <-ch2:
			fmt.Printf("%d ", v)
		case v := <-ch3:
			fmt.Printf("%d ", v)
		default:
			// without blocking
			fmt.Println("I am so boring")
		}
	}
}

func SendWithoutBlocking(ch1, ch2, ch3 chan int) {
	for i := 0; i >= 0; i++ {
		select {
		case ch1 <- i:
		case ch2 <- i:
		case ch3 <- i:
			//default:
		}
	}
}

func sleepRecv(ch1, ch2, ch3 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("%d ", v)
			time.Sleep(time.Second / 10)
		case v := <-ch2:
			fmt.Printf("%d ", v)
			time.Sleep(time.Second / 10)
		case v := <-ch3:
			fmt.Printf("%d ", v)
			time.Sleep(time.Second / 10)
		}
	}
}
