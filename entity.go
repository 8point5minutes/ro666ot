package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ECS struct {
	Entities  []Entity      //list of entities
	Positions map[int]*Tile //entity index: map position
	PlayerID  int           //index of player's entity for convenience
}

func NewECS() *ECS {
	return &ECS{
		Positions: make(map[int]*Tile),
	}
}

// adds entity at a given position and returns its id
func (es *ECS) AddEntity(e Entity, t *Tile) int {
	i := len(es.Entities)
	es.Entities = append(es.Entities, e)
	es.Positions[i] = t
	return i
}

// moves entity to tile
func (es *ECS) MoveEntity(i int, t *Tile) int {
	es.Positions[i] = t
	return i
}

// Entity represents any sort of creature or object on the map
type Entity interface {
	Sprite() *ebiten.Image
}
