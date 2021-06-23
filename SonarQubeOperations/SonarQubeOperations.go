package SonarQubeOperations

import (
	. "SonarQubeResultAlert/AzureOperations"
	. "SonarQubeResultAlert/Config"
	. "SonarQubeResultAlert/Log"
	. "SonarQubeResultAlert/Models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

// GetSonarQubeResults We will process the data we receive from the sonarqube api with the valid pull request id and call the steps to be taken for the issues in it.
func GetSonarQubeResults(threadContent, collectionName, projectName,pullRequestId string) error {
	_,_,username,pass,domain,_ := GetConfigDatas()
	sonarqubeUrl := domain + `api/issues/search?componentKeys=`+projectName+`.Content&resolved=false&pullRequest=`+pullRequestId

	var sonarqubeResult SonarQubeResult

	client := &http.Client{}
	req, _ := http.NewRequest("GET", sonarqubeUrl, nil)
	req.SetBasicAuth(username, pass)
	resp, err := client.Do(req)

	if err != nil {
		CreateLogJson("Error","GetSonarQubeResults", "Error when requesting to get last person assigned to task with Jira api.", err.Error())
		return err
	}

	bodyText, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyText)
	json.Unmarshal([]byte(bodyString), &sonarqubeResult)

	if len(sonarqubeResult.Issues) > 0 {
		repositoryId,err := GetRepositoryId(collectionName, projectName, pullRequestId)
		if err != nil {
			return err
		}

		theradId := CreateThread(threadContent, collectionName, projectName, repositoryId, pullRequestId)

		for _, issue := range sonarqubeResult.Issues {
			commentValue := `{
    "component": "`+issue.Component+`",
    "line": "`+strconv.Itoa(issue.Line)+`",S
    "message": "`+issue.Message+`",
    "author": "@<`+issue.Author+`>"
}`
			err = AddComment(commentValue,theradId, collectionName, projectName, repositoryId, pullRequestId)
			if err != nil {
				return err
			}
		}
	}else{
		return errors.New("No sonarqube data found for the current pull request. Please check your pull request id")
	}
	return nil
}
