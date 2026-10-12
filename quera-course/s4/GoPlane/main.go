package main

import (
	"errors"
	"time"
)

func FastPayment(g1, g2, g3 <-chan string) (string, error) {
	select {
	case res := <-g1:
		return res, nil
	case res := <-g2:
		return res, nil
	case res := <-g3:
		return res, nil
	case <-time.After(2 * time.Second):
		return "", errors.New("Timeout exceeded")
	}
}
