package main

import (
	. "SonarQubeResultAlert/Log"
	. "SonarQubeResultAlert/cmd"
	"net/http"
)

func main()  {
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":8080", nil)	// It allows the API to be opened out of the desired port.

	if err != nil {
		CreateLogJson("Error","Main","ListenAndServe Error",err.Error())
		return
	}
}

// Handler => The function that will run when the api on the server is called.
func Handler(res http.ResponseWriter, req *http.Request) {
	err := Run(req)
	if err != nil {
		res.Write([]byte(`{"Statu": "Error", "IsSucceded": "false", "Message": "`+err.Error()+`"}`))
		return
	}
	res.Write([]byte(`{"Statu": "OK", "IsSucceded": "true"}`))
}