package goroutine

import (
	"fmt"
	"time"
)

//go lang 并发超时控制

//https://www.jianshu.com/p/42e89de33065

func HelloG() {
	time.Sleep(2 * time.Second)
	fmt.Println(time.Now())
}

func TestHelloG() {
	//模拟并发
	for i := 0; i < 10; i++ {
		go HelloG();
	}

	c := make(chan int)
	<-c
}

func run(task_id, sleeptime int, ch chan string) {

	time.Sleep(time.Duration(sleeptime) * time.Second)
	ch <- fmt.Sprintf("task id %d , sleep %d second", task_id, sleeptime)
	return
}

func TestRun() {
	input := []int{3, 2, 1}
	ch := make(chan string)
	startTime := time.Now()
	fmt.Println("Multirun start")
	for i, sleeptime := range input {
		go run(i, sleeptime, ch)
	}

	for range input {
		fmt.Println(<-ch)
	}

	endTime := time.Now()
	fmt.Printf("Multissh finished. Process time %s. Number of tasks is %d", endTime.Sub(startTime), len(input))
}

func TestRun2() {
	c := make(chan int)

	go func() {

		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}

func Run3(task_id, sleeptime, timeout int, ch chan string) {
	ch_run := make(chan string)
	go run3(task_id, sleeptime, ch_run)
	select {
	case re := <-ch_run:
		ch <- re
	case <-time.After(time.Duration(timeout) * time.Second):
		re := fmt.Sprintf("task id %d , timeout", task_id)
		ch <- re
	}
}

func run3(task_id, sleeptime int, ch chan string) {

	time.Sleep(time.Duration(sleeptime) * time.Second)
	ch <- fmt.Sprintf("task id %d , sleep %d second", task_id, sleeptime)
	return
}

func Run33() {
	input := []int{3, 2, 1}
	timeout := 2
	chs := make([]chan string, len(input))
	startTime := time.Now()
	fmt.Println("Multirun start")
	for i, sleeptime := range input {
		chs[i] = make(chan string)
		go Run3(i, sleeptime, timeout, chs[i])
	}

	for _, ch := range chs {
		fmt.Println(<-ch)
	}
	endTime := time.Now()
	fmt.Printf("Multissh finished. Process time %s. Number of task is %d", endTime.Sub(startTime), len(input))
}

func Error() {
	ch := make(chan string)
	ch <- "123"
	fmt.Println(<-ch)

	//ch := make(chan string, 1)
	//ch <- "123"
	//ch <- "123"
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
}

func TimeoutTest() {
	input := []int{3, 2, 1, 5}
	timeout := 2
	chs := make([]chan string, len(input))
	startTime := time.Now()
	fmt.Println("start")
	for i, sleeptime := range input {
		chs[i] = make(chan string)
		go timeoutTestRun(i, sleeptime, timeout, chs[i])
	}
	for _, ch := range chs {
		fmt.Println(<-ch)
	}
	endTime := time.Now()
	fmt.Printf("Multissh finished. Process time %s. Number of task is %d", endTime.Sub(startTime), len(input))
}

func timeoutTestRun(task_id, sleeptime, timeout int, ch chan string) {
	ch_run := make(chan string)
	go timeoutTestrun(task_id, sleeptime, ch_run)
	select {
	case re := <-ch_run:
		ch <- re
	case <-time.After(time.Duration(timeout) * time.Second):
		re := fmt.Sprintf("task id %d, timeout", task_id)
		ch <- re
	}
}

func timeoutTestrun(task_id, sleeptime int, ch chan string) {
	time.Sleep(time.Duration(sleeptime) * time.Second)
	ch <- fmt.Sprintf("task id %d, sleep %d	second", task_id, sleeptime)
	return
}



func limitRun(task_id, sleeptime, timeout int, ch chan string) {
	ch_run := make(chan string)
	go limitrun(task_id, sleeptime, ch_run)
	select {
	case re := <-ch_run:
		ch <- re
	case <-time.After(time.Duration(timeout) * time.Second):
		re := fmt.Sprintf("task id %d , timeout", task_id)
		ch <- re
	}
}

func limitrun(task_id, sleeptime int, ch chan string) {

	time.Sleep(time.Duration(sleeptime) * time.Second)
	ch <- fmt.Sprintf("task id %d , sleep %d second", task_id, sleeptime)
	return
}


func LimitTest() {
	input := []int{3, 2, 1, 5}
	timeout := 2
	chLimit := make(chan bool, 4)
	chs := make([]chan string, len(input))
	limitFunc := func(chLimit chan bool, ch chan string, task_id, sleeptime, timeout int) {
		limitRun(task_id, sleeptime, timeout, ch)
		<-chLimit
	}
	startTime := time.Now()
	fmt.Println("Multirun start")
	for i, sleeptime := range input {
		chs[i] = make(chan string, 1)
		chLimit <- true
		go limitFunc(chLimit, chs[i], i, sleeptime, timeout)
	}

	for _, ch := range chs {
		fmt.Println(<-ch)
	}
	endTime := time.Now()
	fmt.Printf("Multissh finished. Process time %s. Number of task is %d", endTime.Sub(startTime), len(input))
}
