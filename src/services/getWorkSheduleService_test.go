package services

import (
	"employee/src/repo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWorkSheduleSuccess(t *testing.T) {
	service := NewWorkShedulerService(repo.NewEmployeeSheduleRepositury())

	result := service.GetShedule(1, "2020-01-01", "2020-02-01")

	expectedResult := &SheduleResponse{
		Shedule: []*SheduleOnDay{
			{
				Day: "2020-01-01",
				TimeRanges: []map[string]int{
					{"start": 1000, "end": 1300}, {"start": 1000, "end": 1300},
				},
			},
		},
	}

	assert.Equal(t, result, expectedResult)
}
