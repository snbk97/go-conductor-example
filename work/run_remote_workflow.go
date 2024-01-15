package work

import (
	"fmt"
	"os"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	log "github.com/sirupsen/logrus"
)

func RunRemoteWorkFlow(name string, input map[string]interface{}) (string, error) {
	// not considering verison for now
	wf := &model.StartWorkflowRequest{
		Name:          name,
		CorrelationId: "",
		Input:         input,
	}
	workflowId, err := clients.WorkflowExecutor.StartWorkflow(wf)
	if err != nil {
		log.Error("Error starting workflow: ", err)
	}
	if workflowId != "" {
		log.Info(fmt.Sprintf("\nStarted workflow: %s/%s", os.Getenv("WORKFLOW_BASE_URL"), workflowId))
	}
	return workflowId, err
}
