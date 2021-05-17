package practice

import (
	"fmt"
	"runtime"
	"time"
)

var quit chan bool = make(chan bool)

func timerTest() {
	go timer()

	time.Sleep(time.Second * 6)
	close(quit)
	fmt.Println("done", runtime.NumGoroutine())
}

func timer() {
	t := time.NewTimer(time.Second * 2)
	defer t.Stop()
	interval := time.Second * 2

	for {
		select {
		case <-time.After(time.Second * 3):
			// t.Stop()
			fmt.Println("system stop")
			return
		case <-quit:
			fmt.Println("chan stop")
			return
		case tt := <-t.C:
			fmt.Println("timer running ", tt)
			t.Reset(interval)
			//default:
			//	t.Stop()
			//	fmt.Println("system stop")
		}
	}
}
