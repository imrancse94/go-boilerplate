package response

type ErrorResponseStruct struct {
	StatusCode string      `json:"status_code"`
	Message    string      `json:"message"`
	Error      interface{} `json:"error"`
}
