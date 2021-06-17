package message_queue

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"net/http"
	"time"
	"work-wechat/config"
)

var r asynq.RedisClientOpt = asynq.RedisClientOpt{
	Addr:     config.App.Redis.Host,
	Password: config.App.Redis.Password,
	DB:       config.App.Redis.DB,
}

// 入队
func Enqueue(ctx *gin.Context) {
	client := asynq.NewClient(r)

	// Create a task with typename and payload.
	t1 := asynq.NewTask("email:welcome", map[string]interface{}{"user_id": 42})
	t2 := asynq.NewTask("email:reminder", map[string]interface{}{"user_id": 42})

	// Process the task immediately.
	res, err := client.Enqueue(t1, asynq.Queue("welcome"), asynq.MaxRetry(2))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed", "error": err.Error()})
	}
	fmt.Println("welcome result: ", res)

	// Process the task 24 hours later.
	res, err = client.Enqueue(t2, asynq.ProcessIn(5*time.Second), asynq.Queue("reminder"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed", "error": err.Error()})
	}

	fmt.Println("reminder result: ", res)
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
