package controllers

import (
	"go-boilerplate/Helper"
	"go-boilerplate/constant"
	"go-boilerplate/localize"
	"go-boilerplate/logger"
	"go-boilerplate/models"
	"go-boilerplate/requests"
	"go-boilerplate/response"
	Services "go-boilerplate/services"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	/*req := mux.Vars(r)
	fmt.Println("Test Mux", req["id"], req["name"])*/
	request := requests.LoginRequest{}
	Helper.Request(r, &request)
	userData, message := Services.Login(request)
	logData := logger.LogData{}
	logData.Action = "Test"
	localize.SetLocale("en")
	res := response.Response{
		StatusCode: constant.Status("SUCCESS"),
		Message:    localize.Trans(message, ""),
		Data:       userData,
	}

	logData.Data = res
	logger.CreateLog(logData)

	response.SuccessRespond(res, w)

	return
}

func AuthData(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.Header.Get("auth_id"))
	user := models.GetUserById(id)
	res := response.Response{
		StatusCode: constant.Status("SUCCESS"),
		Message:    localize.Trans("Auth data fetched successfully", ""),
		Data:       user,
	}
	response.SuccessRespond(res, w)
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.Header.Get("auth_id"))
	user := models.GetUserById(id)

	data, _ := Services.CreateTokenDataByUser(user)

	res := response.Response{
		StatusCode: constant.Status("SUCCESS"),
		Message:    localize.Trans("Token Refreshed successfully", ""),
		Data:       data,
	}
	response.SuccessRespond(res, w)
}
