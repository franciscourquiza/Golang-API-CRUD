package main

import (
	"net/http"

	"github.com/franciscourquiza/go-gorm-restapi/db"
	"github.com/franciscourquiza/go-gorm-restapi/models"
	"github.com/franciscourquiza/go-gorm-restapi/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/user/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/user", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", routes.DeleteUserHandler).Methods("DELETE")

	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/task/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/task", routes.PostTaskHandler).Methods("POST")
	r.HandleFunc("/task/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)

}
