package main

import (
 "net/http"
 "employee/src/routes"
 "employee/src/services"
 "employee/src/repo"
)

func main() {
	http.HandleFunc("/", handlerGetSheduleById)
	http.HandleFunc("/favicon.ico", func(rw http.ResponseWriter, r *http.Request) {
		return
	})
	http.ListenAndServe(":8080", nil)
}

func handlerGetSheduleById(rw http.ResponseWriter, r *http.Request) {
	repo := repo.NewEmployeeSheduleRepositury()
	service := services.NewWorkShedulerService(repo)
	route := routes.NewWorkShedule(service)
	route.WorkSheduleAction(rw, r)
}