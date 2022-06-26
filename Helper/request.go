package Helper

import (
	"encoding/json"
	"net/http"
)

func Request(r *http.Request, p interface{}) interface{} {
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(p)
	return p
}
