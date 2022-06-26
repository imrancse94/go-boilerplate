package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode string      `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func SuccessRespond(fields Response, writer http.ResponseWriter) {
	if fields.Data == "" {
		fields.Data = []string{}
	}
	message, err := json.MarshalIndent(fields, "", " ")

	if err != nil {
		//An error occurred processing the json
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("An error occurred internally"))
		return
	}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(message)
}

func ErrorResponse(fields ErrorResponseStruct, writer http.ResponseWriter) {
	//Create a new map and fill it
	if fields.Error == "" {
		fields.Error = []string{}
	}
	data, err := json.MarshalIndent(fields, "", " ")

	if err != nil {
		//An error occurred processing the json
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("An error occured internally"))
	}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(data)
	return
}
