package main

import (
 "net/http"
 "./src/routes"
 "./src/services"
 "./src/repo"
)

func main() {
	sheduleRouter := initSheduleRouter()
	http.HandleFunc("/", sheduleRouter.WorkSheduleAction)
	http.HandleFunc("/favicon.ico", func(rw http.ResponseWriter, r *http.Request) {
		return
	})
	http.ListenAndServe(":8080", nil)
}

func initSheduleRouter() *routes.WorkShedule {
	repo := repo.NewEmployeeSheduleRepositury()
	service := services.NewWorkShedulerService(repo)

	return routes.NewWorkShedule(service)
}