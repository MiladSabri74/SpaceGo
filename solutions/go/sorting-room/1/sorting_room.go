package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch fnb.(type) {
	case FancyNumber:
		x, _ := strconv.Atoi(fnb.Value())
		return x
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	x := 0
	switch fnb.(type) {
	case FancyNumber:
		x = ExtractFancyNumber(fnb)
	}
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(x))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	switch i.(type) {
	case int:
		x := i.(int)
		return DescribeNumber(float64(x))
	case float64:
		x := i.(float64)
		return DescribeNumber(x)
	case NumberBox:
		x := i.(NumberBox)
		return DescribeNumberBox(x)
	case FancyNumberBox:
		x := i.(FancyNumberBox)
		return DescribeFancyNumberBox(x)
	}
	return "Return to sender"
}
