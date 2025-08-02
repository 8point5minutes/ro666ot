package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TileType struct {
	Name   string
	Images []*ebiten.Image
}

func NewTileType(name string) *TileType {
	t := &TileType{Name: name, Images: make([]*ebiten.Image, 0)}
	return t
}

func (t *TileType) AddTileImage(imagePath string) {
	img, _, err := ebitenutil.NewImageFromFile(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	t.Images = append(t.Images, img)
}

type Tile struct {
	X          int
	Y          int
	TileType   *TileType
	ImageIndex int
}

func NewTile(x int, y int, tileType *TileType) *Tile {
	t := &Tile{X: x, Y: y, ImageIndex: 0, TileType: tileType}
	return t
}

func (t *Tile) RandomizeImageIndex() {
	t.ImageIndex = GetRandomInt(len(t.TileType.Images))
}

func (t *Tile) Image() *ebiten.Image {
	return t.TileType.Images[t.ImageIndex]
}
