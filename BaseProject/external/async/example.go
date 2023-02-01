package async

import (
	"context"
	"log"
	"time"
)

func Ex() {
	job1 := NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second)
		log.Println("I am job 1")
		return nil
	})

	if err := job1.Excute(context.Background()); err != nil {
		log.Println(job1.state, err)
	}

	job1.SetRetryDurations([]time.Duration{time.Second * 3})
	for {
		if err := job1.Retry(context.Background()); err != nil {
		}

		if job1.State() == StateRetryFailed || job1.State() == StateCompleted {
			log.Println(job1.state)
			break
		}

	}
}
