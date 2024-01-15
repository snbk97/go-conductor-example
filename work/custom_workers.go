package work

import (
	"fmt"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	log "github.com/sirupsen/logrus"
)

// RegisterWorkers util funtion to register workers
func RegisterWorkers(taskName string, executeFunction model.ExecuteTaskFunction) {
	const (
		batchSize    = 3
		pollInterval = 10 * time.Millisecond
	)

	clients.TaskRunner.StartWorker(
		taskName,
		executeFunction,
		batchSize,
		pollInterval,
	)
}

const workerPrefix = "dp"

// WorkerE map of worker enums
var WorkersE = struct {
	ProcessUpdateApproved string
	ProcessUpdateRejected string
}{
	ProcessUpdateApproved: fmt.Sprintf("%s_process_update_approved", workerPrefix),
	ProcessUpdateRejected: fmt.Sprintf("%s_process_update_rejected", workerPrefix),
}

func InitWorkers() {
	RegisterWorkers(WorkersE.ProcessUpdateApproved, handleProcessUpdateApproved)
	RegisterWorkers(WorkersE.ProcessUpdateRejected, handleProcessUpdateRejected)

	log.Info("Workers registered")
}

func handleProcessUpdateApproved(task *model.Task) (interface{}, error) {
	log.Info("\nhandleProcessUpdateApproved:Received Task")
	taskResult := model.NewTaskResultFromTask(task)
	taskResult.Status = model.CompletedTask
	taskResult.Logs = []model.TaskExecLog{
		{
			Log:         "Task executed: handleProcessUpdateApproved",
			TaskId:      task.TaskId,
			CreatedTime: time.Now().UnixMilli(),
		},
	}
	taskResult.OutputData = map[string]interface{}{
		"status": "APPROVED",
	}

	return taskResult, nil
}

func handleProcessUpdateRejected(task *model.Task) (interface{}, error) {
	log.Info("\nhandleProcessUpdateRejected:Received Task")
	taskResult := model.NewTaskResultFromTask(task)
	taskResult.Logs = []model.TaskExecLog{
		{
			Log:         "Task executed: handleProcessUpdateRejected",
			TaskId:      task.TaskId,
			CreatedTime: time.Now().UnixMilli(),
		},
	}
	taskResult.OutputData = map[string]interface{}{
		"status": "REJECTED",
	}

	taskResult.Status = model.CompletedTask
	return taskResult, nil
}
