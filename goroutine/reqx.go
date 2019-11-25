package goroutine

//https://stackoverflow.com/questions/23582143/golang-using-timeouts-with-channels
import (
	"fmt"
	"time"
)


func check(u string, checked chan<- bool) {
	time.Sleep(4 * time.Second)
	checked <- true
}

func IsReachable(urls []string) bool {

	ch := make(chan bool, 1)
	for _, url := range urls {
		go func(u string) {
			checked := make(chan bool)
			go check(u, checked)
			select {
			case ret := <-checked:
				ch <- ret
			case <-time.After(1 * time.Second):
				ch <- false
			}
		}(url)
	}
	return <-ch
}
func ReqX() {
	fmt.Println(IsReachable([]string{"http://www.baidu.com"}))
}
