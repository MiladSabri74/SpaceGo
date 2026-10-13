package main

import "errors"

// CalculateShipping returns the shipping cost and an error if inputs are invalid.
func CalculateShipping(weight float64, distance int, isVIP bool) (int, error) {
    if weight < 0 || distance < 0 {
        return 0, errors.New("invalid input")
    }

    if isVIP {
        return 0, nil
    }

    cost := int(weight*1000) + (distance * 500)
    return cost, nil
}
