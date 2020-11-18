package repo

type SheduleRepository interface {
	GetById(int) int
}

type employeeSheduleRepository struct {
	shedules map[int]string
}

var storage = map[int]int {
	1: 8,
	2: 14,
}

func NewEmployeeSheduleRepositury() *employeeSheduleRepository {
	return &employeeSheduleRepository{}
}

func (e *employeeSheduleRepository) GetById(id int) int {
	return storage[id]
}