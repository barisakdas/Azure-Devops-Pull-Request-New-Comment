package Log

import (
	"log"
	"strings"
	"time"
)

func CreateLogJson(logType, logFunction, logName, logErrorMessage string) string {
	errorLogData := `{
		"Type" : "` + logType + `",
		"Error_Function" : "` + logFunction + `",
		"Name" : "` + logName + `",
		"Description" : "` + logErrorMessage + `",
		"CreatedDate" : "` + strings.Split(time.Now().String(), "+")[0] + `"
	}`

	log.Fatal(errorLogData)
	// fmt.Println(errorLogData)

	return errorLogData
}

