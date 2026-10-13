package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

type Game struct {
	players map[string]*Player
	maps    map[int]*Map
	mu      sync.RWMutex
	cancel  context.CancelFunc
	wg      sync.WaitGroup
}

func NewGame(parentCtx context.Context, mapIds []int) (*Game, error) {
	if parentCtx == nil {
		return nil, errors.New("parent context is nill")
	}

	ctx, cancel := context.WithCancel(parentCtx)

	game := &Game{
		players: make(map[string]*Player),
		maps:    make(map[int]*Map),
		cancel:  cancel,
	}

	for _, id := range mapIds {
		if _, exists := game.maps[id]; exists {
			cancel()
			return nil, fmt.Errorf("duplicate map ID: %d", id)
		}

		game.maps[id] = &Map{
			id:        id,
			players:   make(map[string]*Player),
			broadcast: make(chan Message),
		}
	}

	for _, m := range game.maps {
		game.wg.Add(1)
		go m.FanOutMessages(ctx, &game.wg)
	}
	return game, nil
}

func (g *Game) Shutdown() {
	// TODO: Cancel context and wait for all map goroutines to finish
	g.cancel()
	g.wg.Wait()

}

func (g *Game) ConnectPlayer(name string) error {
	// TODO: Safely create and add new player if name is unique (case-insensitive)
	g.mu.Lock()
	defer g.mu.Unlock()

	lowerName := strings.ToLower(name)
	if _, exists := g.players[lowerName]; exists {
		return errors.New("player name exist")
	}
	g.players[lowerName] = &Player{
		name:     name,
		messages: make(chan string),
	}
	return nil
}

func (g *Game) GetPlayer(name string) (*Player, error) {
	// TODO: Safely find and return player by name (case-insensitive)
	g.mu.RLock()
	defer g.mu.RUnlock()

	lowerName := strings.ToLower(name)
	if _, exists := g.players[lowerName]; exists {
		return g.players[lowerName], nil
	}
	return nil, errors.New("Player not found")
}

func (g *Game) GetMap(mapId int) (*Map, error) {
	// TODO: Safely find and return map by map ID
	g.mu.RLock()
	defer g.mu.RUnlock()

	if _, exists := g.maps[mapId]; exists {
		return g.maps[mapId], nil
	}
	return nil, errors.New("map not found")
}

func (g *Game) SwitchPlayerMap(name string, mapId int) error {
	// TODO: Safely move player from current map to target map
	player, PlayerErr := g.GetPlayer(name)
	if PlayerErr != nil {
		return PlayerErr
	}
	newMap, mapErr := g.GetMap(mapId)
	if mapErr != nil {
		return mapErr
	}
	player.currentMap = newMap
	return nil
}
