package phonenumber

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Number(phoneNumber string) (string, error) {
	phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "(", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, ")", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, ".", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "+", "")
	if len(phoneNumber) == 11 {
		if phoneNumber[0] == '1' {
			phoneNumber = phoneNumber[1:]
		} else {
			return "", errors.New("Invalid phone Number")
		}
	}
	if len(phoneNumber) != 10 {
		return "", errors.New("Invalid phone Number")
	}
	_, err := strconv.Atoi(phoneNumber)
	if err != nil {
		return "", err
	}
	firstAreaCode, _ := strconv.Atoi(string(phoneNumber[0]))
	firstLocalNumber, _ := strconv.Atoi(string(phoneNumber[3]))

	if firstAreaCode < 2 || firstLocalNumber < 2 {
		return "", errors.New("Invalid phone Number")
	}
	return phoneNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[:3], nil
}

func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", number[:3], number[3:6], number[6:]), nil
}
