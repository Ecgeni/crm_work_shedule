package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"employee/src/services"
	"strconv"
)

type WorkShedule struct {
	service *services.WorkShedulerService
}

func NewWorkShedule(service *services.WorkShedulerService) *WorkShedule {
	return &WorkShedule{service: service}
}

func (s *WorkShedule) WorkSheduleAction(w http.ResponseWriter, r *http.Request) {
	
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Print(err)
		return
	}


	shedule := s.service.GetShedule(id, "", "")

	b, err := json.Marshal(shedule)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, " error marshal into json")
		return
	}

	fmt.Fprint(w, string(b))
	log.Print(string(b))
}
