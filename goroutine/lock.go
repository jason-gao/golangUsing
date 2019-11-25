package goroutine

import (
	"fmt"
	"sync"
	"time"
)

// https://www.cnblogs.com/smallleiit/p/10891011.html

// --- 貌似执行结果和预期的不一致

type User struct {
	Name   string
	Locker *sync.Mutex
}

func (u *User) SetName(wati *sync.WaitGroup, name string) {
	defer func() {
		fmt.Println("Unlock set name:", name)
		u.Locker.Unlock()
		wati.Done()
	}()

	u.Locker.Lock()
	fmt.Println("Lock set name:", name)
	time.Sleep(1 * time.Second)
	u.Name = name
}

func (u *User) GetName(wati *sync.WaitGroup) {
	defer func() {
		fmt.Println("Unlock get name:", u.Name)
		u.Locker.Unlock()
		wati.Done()
	}()

	u.Locker.Lock()
	fmt.Println("Lock get name:", u.Name)
	time.Sleep(1 * time.Second)
}

func TestLock() {
	var wg sync.WaitGroup

	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello", time.Now())
	}

	wg.Add(1)
	go sayHello()
	wg.Wait()
}
