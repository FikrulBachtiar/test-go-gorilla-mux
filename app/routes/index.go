package routes

import (
	"test-golang/app/configs"
	"test-golang/app/controllers"
	"test-golang/app/repositories"
	"test-golang/app/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RouteIndex(r *mux.Router, db *gorm.DB, env *configs.EnvironmentVariable) {
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewOtpController(userService)

	r.HandleFunc("/user", userController.GetUser).Methods("GET")
}