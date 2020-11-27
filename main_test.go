package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHandleGetWorkShedule(t *testing.T) {
	testCase := []struct {
		id int
		name string
		want string
	}{
		{
			id: 1,
			name: "one",
			want: "{}",
		},
		{
			id: 2,
			name: "two",
			want: "{}",
		},
	}

	handler := http.HandlerFunc(handlerGetSheduleById)
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/id=%d", tc.id), nil)
			handler.ServeHTTP(rec, req)	
			assert.Equal(t, tc.want, rec.Body.Bytes())
		})
	}
}


