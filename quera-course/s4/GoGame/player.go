package main

import (
	"errors"
	"strings"
	"sync"
)

type Player struct {
	name       string
	currentMap *Map
	messages   chan string
	mu         sync.RWMutex
}

func (p *Player) GetName() string {
	// TODO: Return player name formatted in title case (first letter capitalized, rest lowercase)
	return strings.ToUpper(p.name[:1]) + strings.ToLower(p.name[1:])
}

func (p *Player) GetChannel() <-chan string {
	return p.messages
}

func (p *Player) SendMessage(text string) error {
	// TODO: Send message to current map's broadcast channel, or return an error if not in
	//  any map
	if p.currentMap == nil {
		return errors.New("Player is not in any map")
	}
	p.currentMap.broadcast <- Message{SenderName: p.GetName(), Text: text}
	return nil
}
