package main

import (
	"fmt"
	"log"
	"net/http"
	"test-golang/app/configs"
	"test-golang/app/middleware"
	"test-golang/app/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeApplicationJsonMiddleware)
	environment := configs.LoadEnv()
	db := configs.InitDB(environment);
	routes.RouteIndex(r, db, environment);

	port := fmt.Sprintf(":%s", environment.APP_PORT);
	log.Printf("Server started at %s", port)
	log.Fatal(http.ListenAndServe(port, r));
}