package message_queue

import (
	"fmt"
	"github.com/hibiken/asynq"
	"time"
	"work-wechat/config"
)

var r asynq.RedisClientOpt = asynq.RedisClientOpt{
	Addr:     config.App.Redis.Host,
	Password: config.App.Redis.Password,
	DB:       config.App.Redis.DB,
}

// 入队
func Enqueue() {
	client := asynq.NewClient(r)

	// Create a task with typename and payload.
	t1 := asynq.NewTask("email:welcome", map[string]interface{}{"user_id": 42})
	t2 := asynq.NewTask("email:reminder", map[string]interface{}{"user_id": 42})

	// Process the task immediately.
	res, err := client.Enqueue(t1, asynq.Queue("welcome"), asynq.MaxRetry(2))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("welcome result: ", res)

	// Process the task 24 hours later.
	res, err = client.Enqueue(t2, asynq.ProcessIn(5*time.Second), asynq.Queue("reminder"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("reminder result: ", res)
}
