package controllers

import (
	"encoding/json"
	"net/http"
	"test-golang/app/services"
)

type userController struct {
	userService services.UserService
}

func NewOtpController(userService services.UserService) *userController {
	return &userController{
		userService: userService,
	}
}

func (controller *userController) GetUser(res http.ResponseWriter, req *http.Request) {

	data, err := controller.userService.AllUserFamily()
	if err != nil {
		http.Error(res, "Failed!", http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(data)
	return
	// fmt.Fprint(res, data)
}