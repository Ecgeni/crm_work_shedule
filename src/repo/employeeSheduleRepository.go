package repo

type SheduleRepository interface {
	GetSheduleById(int) []map[string]int
	GetWeekendsById(id int) []map[string]string
}

type employeeSheduleRepository struct {
	shedules map[int]string
}

var storage = map[int]int {
	1: 8,
	2: 14,
}

var shedules = map[int][]map[string]int {
	1: {{"start":1000, "end":1300}, {"start":1000, "end":1300}},
	2: {{"start":900, "end":1200}, {"start":1300, "end":1800}},
}

var weekends = map[int][]map[string]string {
	1: {{"start": "10.01.2020", "end": "11.01.2020"}, {"start": "11.01.2020", "end": "25.01.2020"}, {"start": "01.02.2020", "end": "15.02.2020"}},
	2: {{"start": "10.01.2020", "end": "11.01.2020"}, {"start": "01.02.2020", "end": "01.03.2020"}},
}

func NewEmployeeSheduleRepositury() *employeeSheduleRepository {
	return &employeeSheduleRepository{}
}

func (e *employeeSheduleRepository) GetSheduleById(id int) []map[string]int {
	return shedules[id]
}

func (e *employeeSheduleRepository) GetWeekendsById(id int) []map[string]string {
	return weekends[id]
}