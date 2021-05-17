package message_queue

import (
	"fmt"
	"github.com/hibiken/asynq"
	"golang.org/x/net/context"
)

func handler(ctx context.Context, t *asynq.Task) error {
	fmt.Println("Consumerring: ", t.Type)
	switch t.Type {
	case "email:welcome":
		id, err := t.Payload.GetInt("user_id")
		if err != nil {
			return err
		}
		fmt.Println("Send Welcome Email to User ", id)

	case "email:reminder":
		id, err := t.Payload.GetInt("user_id")
		if err != nil {
			return err
		}
		fmt.Println("Send Reminder Email to User ", id)

	default:
		return fmt.Errorf("unexpected task type: %s", t.Type)
	}
	return nil
}

func Consumer() {
	srv := asynq.NewServer(r, asynq.Config{
		Concurrency: 2,
		Queues: map[string]int{
			"default":  1,
			"welcome":  1,
			"reminder": 1,
		},
	})

	// Use asynq.HandlerFunc adapter for a handler function
	if err := srv.Run(asynq.HandlerFunc(handler)); err != nil {
		fmt.Println("Consumer error：", err)
	}
}
