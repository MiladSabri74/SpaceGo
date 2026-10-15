package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
	"testing"
	"time"
)

var (
	sampleStartOnce sync.Once
	sampleBaseURL   = "http://localhost:18081"
)

func ensureSampleServerStarted() {
	sampleStartOnce.Do(func() {
		go StartServer("18081")
		time.Sleep(300 * time.Millisecond)
	})
}

func TestSample_PostAndGetTicket(t *testing.T) {
	ensureSampleServerStarted()

	reqBody := map[string]string{
		"name": "Nima",
		"ride": "Bumper Cars",
	}
	bodyBytes, _ := json.Marshal(reqBody)

	// Test POST
	postResp, err := http.Post(sampleBaseURL+"/tickets", "application/json", bytes.NewReader(bodyBytes))
	if err != nil {
		t.Fatalf("خطا در ارسال POST: %v", err)
	}
	defer postResp.Body.Close()

	if postResp.StatusCode != http.StatusCreated {
		t.Errorf("کد وضعیت POST باید 201 باشد، اما %d دریافت شد.", postResp.StatusCode)
	}

	var res map[string]interface{}
	json.NewDecoder(postResp.Body).Decode(&res)
	if res["id"] == nil {
		t.Errorf("پاسخ POST باید شامل فیلد id باشد.")
	}

	// Test GET
	getResp, err := http.Get(sampleBaseURL + "/tickets")
	if err != nil {
		t.Fatalf("خطا در ارسال GET: %v", err)
	}
	defer getResp.Body.Close()

	if getResp.StatusCode != http.StatusOK {
		t.Errorf("کد وضعیت GET باید 200 باشد، اما %d دریافت شد.", getResp.StatusCode)
	}
}
