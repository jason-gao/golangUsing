package main

import "jason-gao/golangUsing/goroutine"

func main() {

	// string
	//https://golang.org/pkg/strings/#LastIndex

	//ip := "12.2.111.255"
	//fmt.Println(strings.Index("chicken", "ken"))
	//fmt.Println(strings.Index("chicken", "dmr"))
	//str := "chicken"
	//fmt.Println(str[:4])
	//fmt.Println("=====Index end======")
	//
	//fmt.Println(strings.Index("go gopher", "go"))
	//fmt.Println(strings.LastIndex("go gopher", "go"))
	//fmt.Println(strings.LastIndex("go gopher", "rodent"))
	//pos := strings.LastIndex(ip, ".")
	//fmt.Println(pos)
	//fmt.Print(ip[:pos])
	//fmt.Println("====LastIndex end ====")
	//
	//fmt.Println(strings.LastIndexAny("go gopher", "go"))
	//fmt.Println(strings.LastIndexAny("go gopher", "rodent"))
	//fmt.Println(strings.LastIndexAny("go gopher", "fail"))
	//fmt.Println(strings.LastIndexAny(ip, "."))
	//pos = strings.LastIndex(ip, ".")
	//fmt.Println(ip[:pos])
	//fmt.Println("======LastIndexAny end =====")
	//
	////ip
	//ipLong, err := U.Ip2long(ip)
	//fmt.Println(ipLong, err)
	//
	//ip = U.Long2ip(201486335)
	//U.P(ip)
	//
	////
	//elements := []int{100, 200, 300, 100, 200, 400, 0}
	//fmt.Println(elements)
	//
	//// Test our method.
	//result := U.RemoveDuplicates(elements)
	//fmt.Println(result)

	//goroutine.Test()
	//fmt.Println("goroutine num 1")
	//fmt.Println(goroutine.GoRoutineNum())
	//fmt.Println()
	//goroutine.Test2()
	//
	//fmt.Println("goroutine num 2")
	//fmt.Println(goroutine.GoRoutineNum())

	//lock
	//lock := &sync.Mutex{}
	//
	//for i := 0; i < 10; i++ {
	//	//go goroutine.Add(1,2)
	//	//fmt.Println("xx")
	//	//fmt.Println(goroutine.GoRoutineNum())
	//	go Count(lock)
	//}
	//
	//for {
	//	lock.Lock()
	//	c := counter
	//	lock.Unlock()
	//
	//	runtime.Gosched()
	//
	//	if c >= 10 {
	//		break;
	//	}
	//}
	//
	//fmt.Println("goroutine num 3")
	//fmt.Println(goroutine.GoRoutineNum())
	//
	//fmt.Println("before")

	////str
	//s := string.StrJoin("hello")
	//fmt.Println(s)

	//user := goroutine.User{}
	//user.Locker = new(sync.Mutex)
	//wait := &sync.WaitGroup{}
	//names := []string{"a", "b", "c"}
	//for _, name := range names {
	//	wait.Add(2)
	//	go user.SetName(wait, name)
	//	go user.GetName(wait)
	//}
	//
	//wait.Wait()

	//channel
	//chs := make([] chan int, 10)
	//
	//for i := 0; i < 10; i++ {
	//	chs[i] = make(chan int)
	//	go goroutine.Count(chs[i])
	//}
	//
	//fmt.Println(goroutine.GoRoutineNum())
	//
	//for _, ch := range (chs) {
	//	<-ch
	//}
	//
	//ch1 := make(chan int)
	//ch2 := make(chan int)

	//select {
	//case <-ch1:
	//	fmt.Println(1)
	//case ch2 <- 1:
	//default:
	//	//
	//}

	//channel timeout
	//goroutine.TimeOut()

	//goroutine.RunReq()

	//fmt.Println(goroutine.GoRoutineNum())

	//goroutine.TestLock()

	//goroutine.ChannelRange()

	//fmt.TestPrintf()
	//goroutine.ChannelClose()

	//goroutine.ChannelBuf()
	//goroutine.ChannelSync()
	//goroutine.ChannelAsync()
	//goroutine.ChannelSelect()
	//goroutine.TestWriteRead()
	//goroutine.TestC2()
	//goroutine.TestC3()
	//goroutine.TestC4()
	//goroutine.TestC5()
	//goroutine.TestC6()
	//goroutine.TestNormal()
	//goroutine.TestWait()

	//hosts, _ := hosts("192.168.10.14/28")
	//for _, ip := range hosts {
	//	fmt.Println("sent: " + ip)
	//}
	//
	//


	//cs := make(chan int,2);
	//go printTwo(cs);
	//go printOne(cs);
	//time.Sleep(5 * 1e9);

	//goroutine.TestHelloG();
	//goroutine.TestRun();
	//goroutine.TestRun2();
	//goroutine.Run33();
	//goroutine.Error();
	//goroutine.TimeoutTest()
	//goroutine.LimitTest()
	goroutine.ReqX()




}



