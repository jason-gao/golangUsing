package goroutine

import (
	"fmt"
	"os"
	"time"
)

//var chanName chan ElementType

//var ch chan int
//
//ch := make(chan int)
//
//ch <- value
//
//value := <-ch
//
//c := make(chan int, 1024)
//
//for i := range c{
//	...
//}
//close(c)

func Count(ch chan int) {
	ch <- 1
	fmt.Println("counting")
}

func TimeOut() {
	ch1 := make(chan int)
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(3 * time.Second)
		timeout <- true
	}()

	select {
	case <-ch1:
		fmt.Println("reading...")
	case <-timeout:
		fmt.Println("timeout....")
		//default:
		//	fmt.Println("default....")

	}
}

func ChannelRange() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}

}

func ChannelClose() {
	intStream := make(chan int)
	close(intStream)

	interger, ok := <-intStream

	fmt.Println("(%v): 	%v", ok, interger)

	//ok 返回false表示channel关闭
}

func ChannelBuf() {
	ch := make(chan string, 3)
	ch <- "A"
	ch <- "B"
	ch <- "C"
	//ch <- "C"

	// 超过buffer 发送和接收都会阻塞

	fmt.Println(1)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func ChannelSync() {
	data := make(chan int)

	go func() {
		for d := range data {
			fmt.Println(d)
		}
	}()

	data <- 1
	data <- 2
	data <- 3
	data <- 4

	close(data)
}

func ChannelAsync() {
	data := make(chan int, 3)
	canQuit := make(chan bool) //阻塞主进程，防止未处理完就退出

	go func() {
		for d := range data { //如果data的缓冲区为空，这个协程一直会阻塞，除非channel被close
			fmt.Println(d)
		}
		canQuit <- true
	}()
	data <- 1
	data <- 2
	data <- 3
	data <- 4
	data <- 5
	data <- 6
	data <- 7
	data <- 8
	data <- 9
	close(data) //用完关闭，否则死锁
	<-canQuit   //解除阻塞
}

const (
	MAX_REQUEST_NUM = 10
	CMD_USER_POS    = 1
)

var (
	save chan bool
	quit chan bool
	req  chan *Request
)

type Request struct {
	CmdID int16
	Data  interface{}
}

type UserPos struct {
	X int16
	Y int16
}

func ChannelSelect() {

	go handler()

	var i int16
	for i = 1; i <= 5; i++ {
		newReq := Request{
			CmdID: CMD_USER_POS,
			Data: UserPos{
				X: i,
				Y: 20,
			},
		}
		req <- &newReq
	}

	time.Sleep(2000 * time.Millisecond)

	save <- true
	close(req)

	<-quit

}

func handler() {
	for {
		select {
		case <-save:
			saveGame()
		case r, ok := <-req:
			if ok {
				onReq(r)
			} else {
				fmt.Println("req chan closed")
				os.Exit(0)
			}
		}
	}
}

func init() {
	req = make(chan *Request, MAX_REQUEST_NUM)
	save = make(chan bool)
	quit = make(chan bool)

}

func saveGame() {
	fmt.Printf("Do Something with save game. \n")
	quit <- true
}

func onReq(r *Request) {
	pos := r.Data.(UserPos)
	fmt.Println(r.CmdID, pos)
}

func TestC1() {
	c := make(chan int, 2) // 一个容量为2的缓冲通道
	c <- 3
	c <- 5
	close(c)
	fmt.Println(len(c), cap(c)) // 2 2
	x, ok := <-c
	fmt.Println(x, ok)          // 3 true
	fmt.Println(len(c), cap(c)) // 1 2
	x, ok = <-c
	fmt.Println(x, ok)          // 5 true
	fmt.Println(len(c), cap(c)) // 0 2
	x, ok = <-c
	fmt.Println(x, ok) // 0 false
	x, ok = <-c
	fmt.Println(x, ok)          // 0 false
	fmt.Println(len(c), cap(c)) // 0 2
	close(c)                    // 此行将产生一个恐慌
	c <- 7                      // 如果上一行不存在，此行也将产生一个恐慌。
}

func TestFb() {
	var ball = make(chan string)
	kickBall := func(playName string) {
		for {
			fmt.Print(time.Now(), " ", <-ball, "传球", "\n")
			time.Sleep(time.Second)
			ball <- playName
		}
	}

	go kickBall("A")
	go kickBall("B")
	go kickBall("C")
	go kickBall("D")

	ball <- "裁判"
	var c chan bool
	<-c
}

func TestF() {
	fibonacci := func() chan uint64 {
		c := make(chan uint64)
		go func() {
			var x, y uint64 = 0, 1
			for ; y < (1 << 63); c <- y { // 步尾语句
				x, y = y, x+y
			}
			close(c)
		}()
		return c
	}
	c := fibonacci()
	//for x, ok := <-c; ok; x, ok = <-c { // 初始化和步尾语句
	//	time.Sleep(time.Second)
	//	fmt.Println(x)
	//}
	for x := range c {
		time.Sleep(time.Second)
		fmt.Println(x)
	}
}

//https://gfw.go101.org/article/channel.html
func TestSl() {
	var c chan struct{}
	select {
	case <-c: //阻塞
	case c <- struct{}{}: //阻塞
	default:
		fmt.Println("go here.")
	}
}

func TestSl2() {
	select {}
}

func TestSl3() {
	c := make(chan string, 2)
	trySend := func(v string) {
		select {
		case c <- v: //c缓冲满，执行default
		default:
			fmt.Println("go here.")
		}
	}
	tryReceive := func() string {
		select {
		case v := <-c:
			return v
		default:
			return "-"
		}
	}
	trySend("A")
	trySend("B")
	trySend("C") //发送失败，但不会阻塞

	fmt.Println(tryReceive())
	fmt.Println(tryReceive())
	fmt.Println(tryReceive()) //接收失败

}

func TestSl4() {
	c := make(chan struct{})
	close(c)
	select {
	case c <- struct{}{}: // 若此分支被选中，则产生一个恐慌
	case <-c:
	}
}
