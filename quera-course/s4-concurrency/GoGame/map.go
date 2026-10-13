package main

import (
	"context"
	"fmt"
	"sync"
)

type Map struct {
	id        int
	players   map[string]*Player
	broadcast chan Message
	mu        sync.RWMutex
}

func (m *Map) FanOutMessages(ctx context.Context, wg *sync.WaitGroup) {
	// TODO: Process broadcast channel, format messages, and non-blockingly forward to all other players in map until ctx.Done()
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-m.broadcast:
			playerMsg := fmt.Sprintf("%s says: %s", msg.SenderName, msg.Text)
			m.mu.RLock()
			for _, player := range m.players {
				if player.GetName() != msg.SenderName {
					select {
					case player.messages <- playerMsg:
					default:
						// skip player
					}
				}
			}
			m.mu.RUnlock()
		}
	}
}
