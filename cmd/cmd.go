package Cmd

import (
	. "Sonarqube-Result-Alert/Models"
	. "Sonarqube-Result-Alert/SonarqubeOps"
	"errors"
	"net/http"
)

// Run We transfer the request body that will come from any environment -as we want- to a struct and then send the parameters to the function that will start the process first.
func Run(r *http.Request) error {
	reqModel,err :=Request{}.ParseRequestToModel(r)

	// We start the first step after checking whether the json has data in the post request that will come on the API.
	if reqModel != (Request{}) {
		err = GetSonarQubeResults("Below are the results from the sonarqube check of the current pull request.",reqModel.AzureCollection,reqModel.AzureProject,reqModel.PullRequestID,reqModel.RepositoryName)
		if err != nil {
			return err
		}
	}else {
		return errors.New("Collection name, project name, and pull request ID parameters are required. \n Sample Request: {\n  \"AzureCollection\" : \"ITCollection\",\n  \"AzureProject\" : \"Voltran.SDP\",\n  \"PullRequestId\" : \"405\"\n}")
	}

	return nil
}
