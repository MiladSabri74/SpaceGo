package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestSampleProcessInventory_FiltersZeroQuantity(t *testing.T) {
	input := `[{"name": "apple", "quantity": 10, "price": 1.5}, {"name": "banana", "quantity": 0, "price": 2.0}]`
	res, err := ProcessInventory(input)
	if err != nil {
		t.Fatalf("خطای غیرمنتظره در اجرای تابع ProcessInventory: %v", err)
	}

	var resJSON []map[string]interface{}
	if err := json.Unmarshal([]byte(res), &resJSON); err != nil {
		t.Fatalf("خروجی تابع باید یک رشته جیسون معتبر باشد اما دچار خطا شد: %v", err)
	}

	if len(resJSON) != 1 {
		t.Errorf("انتظار می‌رفت آیتم با quantity=0 از لیست حذف شود، اما در خروجی %d آیتم یافت شد.", len(resJSON))
	}
	
	if !strings.Contains(res, "\n") {
		t.Errorf("خروجی شما باید به شکل دندانه‌دار (Indented) قالب‌بندی شده باشد.")
	}
}
