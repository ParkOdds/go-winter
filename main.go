package main

import (
	"fmt"
	"time"
)

func Ping1S(c chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("ping : ", i)
		time.Sleep(1 * time.Second)
	}
	c <- 10
}

func SendNoti5S(ch chan string) {
	fmt.Println("send noti")
	time.Sleep(5 * time.Second)
	fmt.Println("send noti")

	ch <- "success"
}

func main() {

	c := make(chan int)
	ch := make(chan string)

	go Ping1S(c)
	go SendNoti5S(ch)

	pingVal, notMess := <-c, <-ch
	fmt.Println(pingVal, notMess)
	// time.Sleep(10 * time.Second)
	fmt.Println("completed")
}
