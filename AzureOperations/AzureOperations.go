package AzureOperations

import (
	. "SonarQubeResultAlert/Config"
	. "SonarQubeResultAlert/Log"
	. "SonarQubeResultAlert/Models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// GetRepositoryId To create a thread on Azure devops and to add comments, we need to know the id of the repository where the pr is located.
func GetRepositoryId(collectionName, projectName, pullRequestId string) (string,error) {
	username,pass,_,_,_,domain := GetConfigDatas()
	url := domain+collectionName+`/`+projectName+`/_apis/git/pullrequests/`+pullRequestId+`?api-version=6.0`

	var pr PullRequest

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, pass)
	resp, err := client.Do(req)

	if err != nil {
		CreateLogJson("Error","GetRepositoryId", "Error during api request to fetch boards on Azure Devops", err.Error())
		return "",err
	}

	bodyText, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyText)
	json.Unmarshal([]byte(bodyString), &pr)

	return pr.Repository.ID,nil
}

// CreateThread In order to add a comment to a pr on Azure deveops, it is necessary to create a thread first. So first we go and create a thread.
func CreateThread(threadContent, collectionName, projectName, repositoryId, pullRequestId string) string {
	var response CreateTheradResponse
	username,pass,_,_,_,domain := GetConfigDatas()
	url:= domain+collectionName+`/`+projectName+`/_apis/git/repositories/`+repositoryId+`/pullRequests/`+pullRequestId+`/threads?api-version=6.0`

	body:=`{
    "comments": [
        {
            "parentCommentId": 0,
            "content": "`+threadContent+`",
            "commentType": 1
        }
    ]
}`
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))

	req.SetBasicAuth(username, pass)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		CreateLogJson("Error","CreateThread", "Error while requesting to update tasks on Azure DevOps.", err.Error())
	}

	bodyText, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyText)

	json.Unmarshal([]byte(bodyString), &response)

	CreateLogJson("Info","CreateThread","Creating pull request thread for comments.","AzureDevops thread is created =>"+bodyString)
	return strconv.Itoa(response.ID)
}

// AddComment Function that allows us to add a comment under the created thread
func AddComment(commentValue ,threadId, collectionName, projectName, repositoryId, pullRequestId string) error {
	username,pass,_,_,_,domain := GetConfigDatas()
	url:= domain+collectionName+`/`+projectName+`/_apis/git/repositories/`+repositoryId+`/pullRequests/`+pullRequestId+`/threads/`+threadId+`/comments?api-version=4.1`

	body:=`{
			  "content": "`+commentValue+`",
			  "parentCommentId": 1,
			  "commentType": 1
			}`
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))

	req.SetBasicAuth(username, pass)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		CreateLogJson("Error","CreateThread", "Error while requesting to update tasks on Azure DevOps.", err.Error())
		return err
	}

	bodyText, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyText)

	CreateLogJson("Info","AddComment","Adding comment on pull request thread.","AzureDevops thread comment is added. =>"+bodyString)
	return nil
}

