package controllers

import (
	"go-boilerplate/constant"
	"go-boilerplate/localize"
	"go-boilerplate/models"
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
