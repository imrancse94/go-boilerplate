package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/asaskevich/govalidator"
	"go-boilerplate/gate"
	"go-boilerplate/response"
	"io"
	"io/ioutil"
	"net/http"
)

func DecodeJsonRequest(r *http.Request, interfaceRef interface{}) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, interfaceRef); err != nil {
		if err == io.EOF {
			return errors.New("missing request body")
		}

		return err
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	return nil
}

func GoValidate(data interface{}) interface{} {
	if valid, err := govalidator.ValidateStruct(data); valid == false {
		return govalidator.ErrorsByField(err)
	}

	return nil
}

func ValidateRequest(data interface{}, r *http.Request) interface{} {
	if data != nil {
		if err := DecodeJsonRequest(r, &data); err != nil {
			return err
		}

		if err := GoValidate(data); err != nil {
			return err
		}
	}

	return nil
}

func Validation(data interface{}) gate.Middleware {

	// Create a new gate
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			err := ValidateRequest(data, r)
			// Do middleware things
			if err != nil {
				//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
				//fmt.Fprint(w, err.(error).Error())
				response.ErrorResponse(response.ErrorResponseStruct{
					StatusCode: "E001",
					Message:    "Validation Error",
					Error:      err,
				}, w)
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}
