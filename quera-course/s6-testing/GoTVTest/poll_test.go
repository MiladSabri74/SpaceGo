package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoll(t *testing.T) {
	// TODO: Implement test cases for the Poll
	poll := NewPoll()
	expected := make(map[string]int)

	//test empty option
	poll.CastVote("")
	assert.Equal(t, expected, poll.GetResults(), "empty option")

	var wg sync.WaitGroup
	// test 1000 votes for someone concurrent use gorutine
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			poll.CastVote("Ali")
		}()
	}

	wg.Wait()

	expected["Ali"] = 1000
	assert.Equal(t, expected, poll.GetResults(), "correct")

}
