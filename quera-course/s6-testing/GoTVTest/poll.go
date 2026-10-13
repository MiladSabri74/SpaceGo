package main

import (
	"errors"
	"sync"
)

type Poll struct {
	// TODO: Add necessary fields
	votes map[string]int
	mu    sync.RWMutex
}

func NewPoll() *Poll {
	poll := &Poll{
		votes: map[string]int{},
	}
	return poll
}

func (p *Poll) CastVote(option string) error {
	// TODO: Implement voting logic
	if option == "" {
		return errors.New("empty option")
	}
	p.mu.Lock()
	defer p.mu.Unlock()

	p.votes[option] += 1
	return nil
}

func (p *Poll) GetResults() map[string]int {
	// TODO: Return voting results safely
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.votes
}
