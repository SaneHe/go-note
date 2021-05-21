package rpc

import (
	"fmt"
	"net/rpc"
)

/**
 * @Description: 同步调用
 */
func CallRpcFunc() {
	client, _ := rpc.DialHTTP("tcp", "localhost:1234")
	var response Result
	if err := client.Call("Call.Square", 12, &response); err != nil {
		fmt.Println("Failed to call Cal.Square. ", err)
	}

	fmt.Printf("%d^2 = %d", response.Num, response.Ans)
}

/**
 * @Description: 异步调用
 */
func GoFunc() {
	client, _ := rpc.DialHTTP("tcp", "localhost:1234")
	var response Result
	asyncCall := client.Go("Call.Square", 12, &response, nil)
	fmt.Printf("%d^2 = %d", response.Num, response.Ans)

	<-asyncCall.Done
	fmt.Printf("%d^2 = %d", response.Num, response.Ans)
}
