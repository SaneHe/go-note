package practice

import (
	"fmt"
	"sync"
	"time"
)

const loop = 10

var lastNum int

func Exec() {
	// waitGo
	// channelGo()
	//channel1Go()
	//channel2Go()
	//channel3Go()
	//worker()
	deadLock()
}

/**
 * @Description: waitGroup 等待
 */
func waitGo() {

	var wg sync.WaitGroup

	wg.Add(loop)
	for i := 0; i < loop; i++ {
		// wg.Add(1)
		go func(i int, wa *sync.WaitGroup) {
			defer wa.Done()
			fmt.Printf("%d => %d\n", i, i) // 上下位置互换
			lastNum = i
		}(i, &wg)
	}

	wg.Wait()

	fmt.Println("success -- lastNum", lastNum)
}

/**
 * @Description: 阻塞 channel
 */
func channel1Go() {
	var cha = make(chan int)

	// 写入类型 channel
	go func() {
		//defer close(cha)
		cha <- 1001
		fmt.Println("data<-", 1001)
	}()

	//time.Sleep(1*time.Second)
	fmt.Println(<-cha)
	//for v := range cha {
	//	fmt.Println(v)
	//}
	//close(cha)
	fmt.Println("------complete------")
}

/**
 * @Description: 阻塞 channel 传参
 */
func channel2Go() {
	var cha = make(chan int)

	// 写入类型 channel
	go func(data chan<- int) {
		//defer close(cha)
		data <- 1001
		fmt.Println("data<-", 1001)
		//fmt.Println(data)
	}(cha)

	//time.Sleep(1*time.Second)
	fmt.Println(<-cha)
	//for v := range cha {
	//	fmt.Println(v)
	//}
	//close(cha)
	fmt.Println("------complete------")
}

/**
 * @Description: 阻塞循环 channel
 */
func channel3Go() {
	var ch2 = make(chan int)
	go func(chnl chan int) {
		defer close(chnl)
		for i := 0; i < loop; i++ {
			chnl <- i
			fmt.Println("chnl<- ", i)
		}
	}(ch2)

	for v := range ch2 {
		fmt.Println("successs --", v)
	}

	fmt.Println("complete")
}

/**
 * @Description: 定时任务
 */
func worker() {
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case current := <-ticker.C:
			fmt.Println("执行 1s 定时任务: ", current.String())
		case <-time.After(4 * time.Second):
			ticker.Stop()
			//default:
			//	fmt.Println("failed")
		}
	}
}

/**
 * @Description: 死锁
 * https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651444773&idx=2&sn=3073004239c4140feb1055f0422c6045&chksm=80bb08d7b7cc81c1e733bb1fa8aea9496d679b4be277f727a1c500d42cf956bf5bbda965966e&scene=21#wechat_redirect
 */
func deadLock() {
	ch1 := make(chan int)
	// 死锁
	//go fmt.Println(<-ch1)

	// 正常
	go func() {
		fmt.Println(<-ch1)
	}()
	ch1 <- 5
	time.Sleep(1 * time.Second)
}
