package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Tiles   []*Tile
	GameECS *ECS
}

func NewGame() *Game {
	g := &Game{Tiles: make([]*Tile, 0), GameECS: NewECS()}
	g.Tiles = CreateTiles()
	player := NewPlayer("assets/doombot.png")
	g.GameECS.AddEntity(player, GetTileFromIndex(16, 16, g.Tiles))
	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	gd := NewGameData()
	//rendering the map
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := GetTileFromIndex(x, y, g.Tiles)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.X), float64(tile.Y))
			screen.DrawImage(tile.Image(), op)
		}
	}
	//rendering entities
	for i, entity := range g.GameECS.Entities {
		op := &ebiten.DrawImageOptions{}
		tile := g.GameECS.Positions[i]
		op.GeoM.Translate(float64(tile.X), float64(tile.Y))
		screen.DrawImage(entity.Sprite(), op)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	gd := NewGameData()
	return gd.ScreenWidth * gd.TileWidth, gd.ScreenHeight * gd.TileHeight
}

func main() {
	g := NewGame()
	gd := NewGameData()
	ebiten.SetWindowSize(gd.ScreenWidth*gd.TileWidth, gd.ScreenHeight*gd.TileHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Ro666ot")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
