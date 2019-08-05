package main

import (
	"fmt"
	//"time"
)

func main() {
	// 例 1
	// 同步模式
    // data := make(chan int)
	// 异步模式
	// data := make(chan int, 3)
    // canQuit := make(chan bool) //阻塞主进程，防止未处理完就退出

    // go func() {
    //     for d := range data { //如果data的缓冲区为空，这个协程会一直阻塞，除非被channel被close
    //         fmt.Println(d)
    //     }
    //     canQuit <- true
    // }()

    // data <- 1
    // data <- 2
    // close(data) //用完需要关闭，否则goroutine会被死锁
	// <-canQuit //解除阻塞

	// 例 2
	// s := []int{7, 2, 8, -9, 4, 0}
	// c := make(chan int)
	// go sum(s[:len(s)/2], c)
	// go sum(s[len(s)/2:], c)
	// x, y := <-c, <-c // receive from c，两个都准备好才进行操作，否则一直等待
	// fmt.Println(x, y, x+y)

	// 例 3
	// go func() {
	// 	time.Sleep(1 * time.Hour)
	// }()
	// c := make(chan int)
	// go func() {
	// 	for i := 0; i < 10; i = i + 1 {
	// 		c <- i
	// 	}
	// 	close(c) // 已经关闭的 channel 可以继续读取数据
	// }()
	// // range c产生的迭代值为Channel中发送的值，它会一直迭代直到channel被关闭。上面的例子中如果把close(c)注释掉，程序会一直阻塞在for …… range那一行
	// for i := range c {
	// 	fmt.Println(i)
	// }
	// fmt.Println("Finished")

	// 例 4
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			// <-c 从 c中接受数据
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
			case c <- x: // 如果成功 c 中写入数据
				x, y = y, x+y
			case <-quit: // 如果 quit 成功读取到数据
				fmt.Println("quit")
				return
			default:
				// 代码块，最多只有一个
		}
	}
}