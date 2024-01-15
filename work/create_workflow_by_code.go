// Work package contains all the workflow definitions
// and the workflow execution code.
package work

import (
	"fmt"
	"os"

	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	log "github.com/sirupsen/logrus"
)

const _webhook_url_ = "https://webhook.site/e587c1b7-5431-47c4-ad0e-8c306a05729d"

func CreateNewWorkflow() *workflow.ConductorWorkflow {
	// create & register workflow
	wf := workflow.NewConductorWorkflow(clients.WorkflowExecutor).
		Name("go-workflow").
		Description("Create a new workflow via go code").
		InputParameters("userId", "notificationChannel").
		Version(1).
		Add(FetchUserTask()).
		Add(SetupSwitchTask())

	clients.WorkflowExecutor.RegisterWorkflow(true, wf.ToWorkflowDef())
	return wf
}

func FetchUserTask() *workflow.HttpTask {
	httpInput := &workflow.HttpInput{
		Method: "GET",
		Uri:    "https://reqres.in/api/users/" + "${workflow.input.userId}",
	}
	wf := workflow.NewHttpTask("go-workflow-http-task-fetch-user", httpInput)

	return wf
}

func SetupSwitchTask() *workflow.SwitchTask {

	return workflow.NewSwitchTask("go-workflow-switch-notify-user", "${workflow.input.notificationChannel}").
		// SwitchCase("6", HandleEmailNotification2()).
		SwitchCase("Email", HandleEmailNotification()).
		SwitchCase("Sms", HandleSMSNotification())

}

func HandleSMSNotification() workflow.TaskInterface {
	fakenumber := "555-555-5555"
	return workflow.NewHttpTask("go-workflow-http-task-send-sms", &workflow.HttpInput{
		Method: "POST",
		Uri:    _webhook_url_,
		Body: map[string]interface{}{
			// "email":   "${workflow.input.email}",
			"phone":   fakenumber,
			"message": "Hello from go workflow",
		},
	})
}
func HandleEmailNotification2() workflow.TaskInterface {

	return workflow.NewHttpTask("go-workflow-http-task-send-email2", &workflow.HttpInput{
		Method: "POST",
		Uri:    _webhook_url_,
		Body: map[string]interface{}{
			"email":   "${workflow.input.userId}",
			"message": "Hello from go workflow",
		},
	})
}

func HandleEmailNotification() workflow.TaskInterface {

	return workflow.NewHttpTask("go-workflow-http-task-send-email", &workflow.HttpInput{
		Method: "POST",
		Uri:    _webhook_url_,
		Body: map[string]interface{}{
			"email":   "${workflow.input.userId}",
			"message": "Hello from go workflow",
		},
	})
}

func RunCreatedWorkflow(userId string, notificationChannel string) {
	wf := CreateNewWorkflow()
	workFlowRun, err := wf.ExecuteWorkflowWithInput(map[string]interface{}{
		"userId":              userId,
		"notificationChannel": notificationChannel,
	}, "")
	if err != nil {
		log.Error("Error starting workflow: ", err)
	}
	if workFlowRun != nil {
		workflowUrl := fmt.Sprintf("%s/%s", os.Getenv("WORKFLOW_BASE_URL"), workFlowRun.WorkflowId)
		log.Info("\n\nGO workflow run complete: ", workflowUrl)

	}
}
