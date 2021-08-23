package Helpers

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func HttpPost(url, body, username, pass string) (string,error) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))

	req.SetBasicAuth(username, pass)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return err.Error(),err
	}

	bodyText, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyText)

	return bodyString,nil
}
