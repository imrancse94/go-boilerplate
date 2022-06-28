package controllers

import (
	"go-boilerplate/Helper"
	"go-boilerplate/constant"
	"go-boilerplate/localize"
	"go-boilerplate/logger"
	"go-boilerplate/models"
	"go-boilerplate/requests"
	"go-boilerplate/response"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	/*req := mux.Vars(r)
	fmt.Println("Test Mux", req["id"], req["name"])*/
	request := requests.LoginRequest{}
	Helper.Request(r, &request)

	logData := logger.LogData{}
	logData.Action = "Test"
	localize.SetLocale("bn")
	res := response.Response{
		StatusCode: constant.Status("SUCCESS"),
		Message:    localize.Trans("I have {{.amount}} Taka", `{"amount":"10"}`),
		Data:       models.GetUserByEmail(request.Email),
	}

	logData.Data = res
	logger.CreateLog(logData)

	response.SuccessRespond(res, w)

	return
}

func User(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>User</h1>"))
}
