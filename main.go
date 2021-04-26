package main

import (
	"strconv"
	. "work-wechat/config"
)

func main() {
	r := initRouter()
	r.Run(App.Server.Host + ":" + strconv.Itoa(App.Server.Port))
}
