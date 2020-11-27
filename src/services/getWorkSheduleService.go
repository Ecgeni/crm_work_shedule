package services

import (
	"employee/src/repo"
	"time"
)

type WorkShedulerService struct {
	repo repo.SheduleRepository
}

type SheduleOnDay struct {
	Day 	   string				`json:"day"`
	TimeRanges []map[string]int		`json: "timeRanges"`
}

type SheduleResponse struct {
	Shedule []*SheduleOnDay         `json:"shedule"`
}

func NewWorkShedulerService(repo repo.SheduleRepository) *WorkShedulerService {
	return &WorkShedulerService{repo: repo}
}

func (w *WorkShedulerService) GetShedule(id int, startDate string, endDate string) *SheduleResponse {
	currentShedules := w.calculateShedule(startDate, endDate, w.repo.GetSheduleById(id), w.repo.GetWeekendsById(id));
	shedule := SheduleResponse{Shedule: currentShedules}

	return &shedule
}

func (w *WorkShedulerService) calculateShedule(startDate string, endDate string, shedule []map[string]int, weekends []map[string]string) []*SheduleOnDay {
	var result []*SheduleOnDay
	result = make([]*SheduleOnDay, 1)
	var isWeekendDay bool
	startTime, _ := time.Parse("2006-01-02", startDate)
	endTime, _ := time.Parse("2006-01-02", endDate)

	timeNextDay := startTime.Add(time.Hour * 24)
	
	for timeNextDay.Equal(endTime) {
		timeNextDay := timeNextDay.Add(time.Hour * 24)
		isWeekendDay = false

		for _, period := range weekends {
			startPeriod, _ := time.Parse("2006-01-02", period["start"])
			endPeriod, _ := time.Parse("2006-01-02", period["end"])
			if timeNextDay.After(startPeriod) && timeNextDay.Before(endPeriod) {
				isWeekendDay = true
				break
			} 
		}

		if !isWeekendDay {
			result = append(result, &SheduleOnDay{Day: timeNextDay.Format("2006-01-02"), TimeRanges: shedule})
		}
	}	
	 
	return result
}
