package main

import (
	"strconv"
	. "work-wechat/config"
)

func main() {
	// 监听队列并消费
	// go message_queue.Consumer()
	r := initRouter()
	r.Run(App.Server.Host + ":" + strconv.Itoa(App.Server.Port))
}
