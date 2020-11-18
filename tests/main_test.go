package tests

import ("testing"
		"net/http"
		"fmt"
	)

func TestGetWorkShedule(t *testing.T) {
	resp, err := http.Get("localhost:8080/")
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	} 
	body := make([]byte, 128)
	n, err := resp.Body.Read(body)
		
	if n == 0 || err != nil{
		return
	}
	result := string(body[:n])

	if result != "Hello World" {
		t.Error("Response", result)
	}
}


