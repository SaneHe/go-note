package practice

import (
	"time"
)

func func1(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "func1 ok"
}

func func2(ch chan string) {
	time.Sleep(time.Second * 10)
	ch <- "func2 ok"
}

func run(f func(ch chan string), cOut chan string) {
	chRun := make(chan string)
	go f(chRun)

	select {
	case r := <-chRun:
		cOut <- r
	case <-time.After(time.Second * 3): // 超时控制
		cOut <- "time out"
	}
}

func SyncCron(functions ...func(ch chan string)) []string {
	// 创建n个chan用来接受任务结果
	chs := make([]chan string, len(functions))

	// 关闭 chan
	defer func() {
		for _, c := range chs {
			if c != nil {
				close(c)
			}
		}
	}()

	// 执行 func
	for i, f := range functions {
		chs[i] = make(chan string)
		go run(f, chs[i])
	}

	// 取结果
	result := make([]string, len(functions))
	for i, ch := range chs {
		result[i] = <-ch
	}

	// 返回结果
	return result
}
