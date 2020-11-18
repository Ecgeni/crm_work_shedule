package services

import "../repo"

type WorkShedulerService struct {
	repo repo.SheduleRepository
}

type sheduleOnDay struct {
	Day string                   `json:"day"`
	TimeRanges []map[string]int  `json:"timeRanges"`
}

type sheduleResponse struct {
	Shedule []sheduleOnDay       `json:"shedule"`
}

func NewWorkShedulerService(repo repo.SheduleRepository) *WorkShedulerService {
	return &WorkShedulerService{repo: repo}
}

func (w *WorkShedulerService) GetShedule(id int, startDate string, endDate string) *sheduleResponse {
	sheduleById := w.repo.GetById(id)
	startEnd := map[string]int{"start": sheduleById, "end":sheduleById}
	timeRange := []map[string]int{startEnd}
	day := sheduleOnDay{Day: "2020-01-01", TimeRanges: timeRange}
	dday := []sheduleOnDay{day}
	shedule := sheduleResponse{Shedule: dday}

	return &shedule
}
