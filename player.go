package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	playerSprite *ebiten.Image
}

func NewPlayer(imagePath string) Player {
	img, _, err := ebitenutil.NewImageFromFile(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	t := Player{playerSprite: img}
	return t
}

func (p Player) Sprite() *ebiten.Image {
	return p.playerSprite
}
