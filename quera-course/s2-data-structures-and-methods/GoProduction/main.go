package main

import (
	"math"
	"strconv"
)

type FilterFunc func(int) bool
type MapperFunc func(int) int

func IsSquare(x int) bool {
	r := math.Sqrt(float64(x))
	return int(r)*int(r) == x
}

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	j := len(str) - 1
	for i := 0; i < len(str); i++ {
		if str[i] != str[j] {
			return false
		}
		j--
	}
	return true
}

func Abs(num int) int {
	// TODO: Return the absolute value of num
	r := math.Abs(float64(num))
	return int(r)
}

func Cube(num int) int {
	// TODO: Return the cube of num
	return num * num * num
}

func Filter(input []int, f FilterFunc) []int {
	// TODO: Filter elements based on predicate function f
	var result []int
	for _, v := range input {
		if f(v) {
			result = append(result, v)
		}

	}
	return result
}

func Map(input []int, m MapperFunc) []int {
	// TODO: Transform elements using mapper function m
	var result []int
	for _, v := range input {
		result = append(result, m(v))
	}
	return result
}
