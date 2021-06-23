package Config

import (
	. "SonarQubeResultAlert/Models"
	"encoding/json"
	"io/ioutil"
)

func GetConfigDatas() (string, string, string, string, string, string) {
	data, _ := ioutil.ReadFile("Config/config.json")
	//data, _ := ioutil.ReadFile("config.json") // Canlıya geçerken bunu aç. Sebebi kubernetes içinde Config klasörünü değil sadece içindekini kopyaladı.

	var config ConfigDatas
	var azureUserName, azurePathToken, sonarqubeUserName, sonarqubePassword,sonarqubeDomain,azureDevopsDomain string

	json.Unmarshal(data, &config)
	for _, config := range config.ConfigDatas {
		switch config.Key {
		case "azureUserName":
			azureUserName = config.Value
		case "azurePathToken":
			azurePathToken = config.Value
		case "SonarqubeUserName":
			sonarqubeUserName = config.Value
		case "SonarqubePassword":
			sonarqubePassword = config.Value
		case "sonarqubeDomain":
			sonarqubeDomain = config.Value
		case "azureDevopsDomain":
			azureDevopsDomain = config.Value
		}
	}
	return azureUserName, azurePathToken, sonarqubeUserName, sonarqubePassword, sonarqubeDomain, azureDevopsDomain
}
