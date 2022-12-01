package cloudafrica

import (
	"context"
	"fmt"
	"time"
)

func WaitForTask(client APIClient, auth context.Context, taskId int64, attempts int, sleep time.Duration) (err error) {
	for i := 0; i < attempts; i++ {
		res, _, err := client.DefaultApi.TasksIdGet(auth, taskId).Execute()
		if err == nil {
			if *res.Task.Status == "complete" {
				return nil
			}
		}
		if i > 0 {
			time.Sleep(sleep)
		}
		//err = f()
		//if err == nil {
		//	return nil
		//}
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}
