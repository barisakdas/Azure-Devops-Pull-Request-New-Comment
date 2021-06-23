package cmd

import (
	. "SonarQubeResultAlert/Log"
	. "SonarQubeResultAlert/Models"
	. "SonarQubeResultAlert/SonarQubeOperations"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Run We transfer the request body that will come from any environment -as we want- to a struct and then send the parameters to the function that will start the process first.
func Run(r *http.Request) error {
	var reqModel Request
	bodyText, err := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal([]byte(bodyText), &reqModel)

	if err != nil {
		CreateLogJson("Error","Run", "An error occurred when converting the data from the request via API from json to a golang struct.", err.Error())
		return err
	}

	// We start the first step after checking whether the json has data in the post request that will come on the API.
	if reqModel != (Request{}) {
		err = GetSonarQubeResults("Below are the results from the sonarqube check of the current pull request.",reqModel.AzureCollection,reqModel.AzureProject,reqModel.PullRequestID)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Collection name, project name, and pull request ID parameters are required. \n Sample Request: {\n  \"AzureCollection\" : \"ITCollection\",\n  \"AzureProject\" : \"Voltran.SDP\",\n  \"PullRequestId\" : \"405\"\n}")
	}

	return nil
}
