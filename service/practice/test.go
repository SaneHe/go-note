package practice

import (
	"fmt"
	"runtime"
	"sync"
)

//func main() {
//	var Ball int
//	table := make(chan int)
//	go player(table)
//	go player(table)
//
//	table <- Ball
//	time.Sleep(1 * time.Second)
//	<-table
//}
//
//func player(table chan int) {
//	f, err := os.OpenFile("./static/gin.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
//	if err != nil {
//		panic("创建 static 文件失败")
//	}
//	defer f.Close()
//
//	for {
//		ball := <-table
//		ball++
//		time.Sleep(100 * time.Millisecond)
//		table <- ball
//		io.WriteString(f, "一")
//	}
//}

// 定义goroutine 1
//func Echo(out chan<- string) {   // 定义输出通道类型
//	out <- "23423"
//	//close(out)
//}
//
//// 定义goroutine 2
//func Receive(out chan<- string, in <-chan string) { // 定义输出通道类型和输入类型
//	temp := <-in // 阻塞等待echo的通道的返回
//	out <- temp
//	//close(out)
//}
//
//
//func main() {
//	echo := make(chan string)
//	receive := make(chan string)
//
//	go Echo(echo)
//	go Receive(receive, echo)
//
//	getStr := <-receive   // 接收goroutine 2的返回
//
//	fmt.Println(getStr)
//}

type user struct {
	Name string
}

type Admin struct {
	user
}

func Test() {
	var ad Admin
	ad.Name = "张三"
	fmt.Println(ad, ad.user.Name)

	// 将内核处理器设置为1个
	runtime.GOMAXPROCS(1)

	// 计数的信号量
	var wg sync.WaitGroup
	// 使用 Add 方法设设置计算器为2
	wg.Add(2)
	go func() {
		// goroutine 的函数执行完之后，就调用 Done 方法减1
		defer wg.Done()
		for i := 1; i < 10; i++ {
			// 增加耗时，当前 goroutine 被扔进全局队列，从局部队列拿其他 goroutine 执行
			//time.Sleep(time.Second * 1)
			fmt.Println("A:", i)
		}
	}()
	go func() {
		// goroutine 的函数执行完之后，就调用 Done 方法减1
		defer wg.Done()
		for i := 1; i < 10; i++ {
			// 增加耗时，当前 goroutine 被扔进全局队列，从局部队列拿其他 goroutine 执行
			//time.Sleep(time.Second * 1)
			fmt.Println("B:", i)
		}
	}()
	// Wait 方法的意思是如果计数器大于0，就会阻塞
	wg.Wait()
	// 所以 main 函数会一直等待2个 goroutine 完成后，再结束。
}
