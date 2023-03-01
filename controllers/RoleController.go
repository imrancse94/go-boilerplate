package controllers

import (
	"go-boilerplate/Helper"
	"go-boilerplate/constant"
	"go-boilerplate/localize"
	"go-boilerplate/models"
	"go-boilerplate/requests"
	"go-boilerplate/response"
	"net/http"
)

func GetRoles(w http.ResponseWriter, r *http.Request) {

	roles := models.GetRoles()

	res := response.Response{
		StatusCode: constant.Status("SUCCESS"),
		Message:    localize.Trans("success", ""),
		Data:       roles,
	}
	response.SuccessRespond(res, w)
}

func AddRole(w http.ResponseWriter, r *http.Request) {
	request := requests.AddRoleRequest{}
	Helper.Request(r, &request)

	statusCode := constant.Status("FAILED")

	if true {
		statusCode = constant.Status("SUCCESS")
	}

	res := response.Response{
		StatusCode: statusCode,
		Message:    localize.Trans("success", ""),
		Data:       request,
	}
	response.SuccessRespond(res, w)
}
