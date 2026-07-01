package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	h int
	m int
}

func New(h, m int) Clock {
	var c Clock
	for m >= 60 {
		m -= 60
		h++
	}
	for h >= 24 {
		h -= 24
	}
	for m < 0 {
		m += 60
		h--
	}
	for h < 0 {
		h += 24
	}
	c.h = h
	c.m = m
	return c
}

func (c Clock) Add(m int) Clock {
	c.m = c.m + m
	for c.m >= 60 {
		c.m -= 60
		c.h++
	}
	for c.h >= 24 {
		c.h -= 24
	}
	return c
}

func (c Clock) Subtract(m int) Clock {
	c.m = c.m - m
	for c.m < 0 {
		c.m += 60
		c.h--
	}
	for c.h < 0 {
		c.h += 24
	}
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
