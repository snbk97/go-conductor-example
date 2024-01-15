package work

import (
	"os"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/conductor-sdk/conductor-go/sdk/worker"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
	log "github.com/sirupsen/logrus"
)

type ConductorClients struct {
	apiClient        *client.APIClient
	WorkflowExecutor *executor.WorkflowExecutor
	TaskRunner       *worker.TaskRunner
}

var clients = &ConductorClients{}

func InitClients() {
	clients.apiClient = getApiClient()
	clients.WorkflowExecutor = executor.NewWorkflowExecutor(clients.apiClient)
	clients.TaskRunner = worker.NewTaskRunnerWithApiClient(clients.apiClient)
}

func getApiClient() (apiClient *client.APIClient) {
	authSettings := settings.NewAuthenticationSettings(
		os.Getenv("CONDUCTOR_KEY"), os.Getenv("CONDUCTOR_SECRET"),
	)
	log.Info("Auth settings: ", authSettings)

	httpSettings := settings.NewHttpSettings(os.Getenv("CONDUCTOR_URL"))

	apiClient = client.NewAPIClient(
		authSettings,
		httpSettings,
	)

	return
}
