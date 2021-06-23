package Models

// Json model that needs to be dumped on api
type Request struct {
	AzureCollection string `json:"AzureCollection"`
	AzureProject    string `json:"AzureProject"`
	PullRequestID   string `json:"PullRequestId"`
}


