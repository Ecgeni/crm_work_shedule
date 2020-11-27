package services

import (
	"testing"
	"employee/src/repo"
	"github.com/stretchr/testify/assert"
)

func TestGetWorkSheduleSuccess(t *testing.T) {
	service := services.NewWorkShedulerService(repo.NewEmployeeSheduleRepositury())

	result := service.GetShedule(1, "2020-01-01", "2020-02-01")

	assert.Equal(t, result, services.SheduleResponse{Shedule: {&services.SheduleOnDay{Day: "2020-01-01", TimeRanges: {{"start":1000, "end":1300}, {"start":1000, "end":1300}}}}})
}