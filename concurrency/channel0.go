package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		fmt.Println(v);
		time.Sleep(1000 * time.Millisecond)
	}
	c <- sum // 把 sum 发送到通道 c
}

// goroutine, channel system
// channel: 1:1, 1:n, n:n
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan int)

	fmt.Println("main: 1");
	go sum(s[:len(s)/2], c)
	
	fmt.Println("main: 2");
	go sum(s[len(s)/2:], c)
	
	fmt.Println("main: 3");
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}
