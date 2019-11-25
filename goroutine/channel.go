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





