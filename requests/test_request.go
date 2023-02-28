package requests

type TestRequest struct {
	Type string `json:"type"`
	Key  string `json:"key"`
	Data string `json:"data"`
}
