package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/**
- https://github.com/chai2010/advanced-go-programming-book/blob/master/ch1-basic/ch1-06-goroutine.md
 */
func printHello() {
	fmt.Println("Hello World")
}

func Test() {
	fmt.Println("main execution start")

	go printHello()
	fmt.Println("main execution stop")
}

func Test2() {
	fmt.Println("main execution start")

	go printHello()

	fmt.Println("goroutine num: ", GoRoutineNum())
	time.Sleep(10 * time.Millisecond)
	fmt.Println("main execution stop")
}

func GoRoutineNum() int {
	n := runtime.NumGoroutine()

	return n
}

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func TestWriteRead() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages

	fmt.Println(msg)

}

func sum(s [] int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum
}

func TestC2() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

}

func TestC3() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func TestC4() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return

		}
	}
}

func TestC5() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci2(c, quit)
}

func TestC6() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func TestNormal()  {
	done := make(chan int, 10) // 带 10 个缓存

	// 开N个后台打印线程
	for i := 0; i < cap(done); i++ {
		go func(){
			fmt.Println("你好, 世界", time.Now())
			time.Sleep(5*time.Second)
			done <- 1
		}()
	}

	// 等待N个后台线程完成
	for i := 0; i < cap(done); i++ {
		<-done
	}
}

func TestWait()  {
	var wg sync.WaitGroup

	// 开N个后台打印线程
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			fmt.Println("你好, 世界", time.Now())
			wg.Done()
		}()
	}

	// 等待N个后台线程完成
	wg.Wait()
}



// 生产者: 生成 factor 整数倍的序列
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i*factor
	}
}

// 消费者
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}


func TestConsumerProducer()  {
	
}