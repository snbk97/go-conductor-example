package work

import (
	"ccc/models"
	"fmt"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	log "github.com/sirupsen/logrus"
)

const (
	DECISION_APPROVED = "APPROVED"
	DECISION_REJECTED = "REJECTED"
)

// RunMakerChecker function is wrokflow runner which runs a particular workflow
// onsubmit of a request. the request is the MAKER in this scenario
// CHECKER will either approve or reject the request
// each step has provision of visibility, where email can be sent to the necessary person.
func RunMakerChecker() {
	// input payload: consider this as a user submitting a request
	userSubmitPayload := &models.PublishPayload{
		CreatedBy:          "5",
		Status:             "PUBLISHED",
		WorkflowID:         "null",
		ChangeRequestCount: "0",
	}

	workflowId, err := RunRemoteWorkFlow("MakerCheckerFlow2", map[string]interface{}{
		"userId":     userSubmitPayload.CreatedBy,
		"checkerDL":  "checkers.dl@email.com",
		"makerEmail": "user@email.com",
	})
	if err != nil {
		log.Error("Error starting workflow: ", err)
		return
	}

	// once the workflow is started, the CHECKER is notified of the new request via email
	// the workflow will be pause at a HUMAN task for the checker to check the values, provide decision
	log.Info("\nWorkflow started: ", workflowId)

	// handle the checker decision
	SubmitValueForCheckerDecision(workflowId, DECISION_REJECTED)

	/*
		log.Info("\nTerminate workflow: ", workflowId)
		clients.WorkflowExecutor.Terminate(workflowId, "Terminated by user")
	*/
}

// Sends the decision to the workflow
// this is a manual step, which can be placed behind a UI + http endpoint
func SubmitValueForCheckerDecision(workflowId string, decision string) {
	log.Info("Waiting for 5 sec before submitting value for human task")
	time.Sleep(5 * time.Second)

	wf, err := clients.WorkflowExecutor.GetWorkflow(workflowId, true)
	if err != nil {
		log.Error("Error getting workflow: ", err)
		return
	}
	for _, task := range wf.Tasks {
		fmt.Println("task", task.TaskDefName, task.TaskType, task.Status)

		if task.TaskDefName == "CheckerDecision" {
			clients.WorkflowExecutor.UpdateTask(
				task.TaskId,
				workflowId,
				model.CompletedTask,
				map[string]interface{}{
					"selectedOption": decision,
				},
			)
			break
		}
	}
}
