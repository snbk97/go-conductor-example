package models

type PublishPayload struct {
	// add more fields here
	CreatedBy          string `json:"createdBy"`
	Status             string `json:"status"`
	WorkflowID         string `json:"workflowId"`
	ChangeRequestCount string `json:"changeRequestCount"`
}
