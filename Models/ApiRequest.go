package Models

import (
	. "Sonarqube-Result-Alert/Log"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Request Json model that needs to be dumped on api
type Request struct {
	AzureCollection 	string `json:"AzureCollection"`
	AzureProject    	string `json:"AzureProject"`
	PullRequestID   	string `json:"PullRequestId"`
	AzureDevopsDomain   string `json:"AzureDevopsDomain"`
	RepositoryName   	string `json:"RepositoryName"`
}

func (req Request) ParseRequestToModel(r *http.Request) (Request,error){
	var reqModel Request
	bodyText, err := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal([]byte(bodyText), &reqModel)

	if err != nil {
		CreateLogJson("Error","Run", "An error occurred when converting the data from the request via API from json to a golang struct.", err.Error())
		return Request{},err
	}
	return reqModel,nil
}