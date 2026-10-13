package main

import (
	"sync"
)

type Inventory struct {
	mu    sync.Mutex
	items map[string]int
}

func NewInventory() *Inventory {
	inventory := &Inventory{
		items: make(map[string]int),
		mu:    sync.Mutex{},
	}
	return inventory
}

func (i *Inventory) AddStock(item string, count int) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.items[item] += count
}

func (i *Inventory) GetStock(item string) int {
	i.mu.Lock()
	defer i.mu.Unlock()
	x := i.items[item]
	return x
}

func (i *Inventory) Purchase(item string, count int) bool {
	i.mu.Lock()
	defer i.mu.Unlock()
	if i.items[item] >= count {
		i.items[item] -= count
		return true
	}
	return false
}
