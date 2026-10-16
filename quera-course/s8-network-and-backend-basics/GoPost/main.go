package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Response struct {
	Status string `json:"status"`
}

func GetTrafficStatus(city string) (string, error) {
	// TODO: fetch traffic status using the third-party API
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return "", errors.New("api key is invalid")
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	url := "http://localhost:8080/trafics"

	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		return "", errors.New("can not create request")
	}

	req.Header.Set("User-Agent", "Go-HTTP-Client/1.0")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("error in sending request")
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status is : %s", http.StatusText(resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var result Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}
	return result.Status, nil
}
