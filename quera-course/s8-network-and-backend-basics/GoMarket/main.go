package main

import (
	"encoding/json"
)

type Product struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func ProcessInventory(rawJSON string) (string, error) {
	// TODO: implement JSON processing logic
	var inputProducts, outputProducts []Product

	err := json.Unmarshal([]byte(rawJSON), &inputProducts)
	if err != nil {
		return "", err
	}
	for _, p := range inputProducts {
		if p.Quantity != 0 {
			outputProducts = append(outputProducts, p)
		}
	}
	result, resultErr := json.MarshalIndent(&outputProducts, "", "  ")
	if resultErr != nil {
		return "", resultErr
	}
	return string(result), nil
}
