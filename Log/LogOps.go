package Log

import (
	"log"
	"strings"
	"time"
)

func CreateLogJson(logType, logFunction, logName, logErrorMessage string) {
	errorLogData := `{
		"Type" : "` + logType + `",
		"Function" : "` + logFunction + `",
		"Name" : "` + logName + `",
		"Description" : "` + logErrorMessage + `",
		"CreatedDate" : "` + strings.Split(time.Now().String(), "+")[0] + `"
	}`

	log.Fatal(errorLogData)
	// fmt.Println(errorLogData)
}