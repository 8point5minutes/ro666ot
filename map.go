package main

func GetIndexFromXY(x int, y int) int {
	gd := NewGameData()
	return (x * gd.ScreenHeight) + y
}

func GetTileFromIndex(x int, y int, tiles []*Tile) *Tile {
	return tiles[GetIndexFromXY(x, y)]
}

func CreateTiles() []*Tile {
	gd := NewGameData()
	tiles := make([]*Tile, 0)
	ground := NewTileType("ground")
	ground.AddTileImage("assets/ground_0.png")
	ground.AddTileImage("assets/ground_1.png")
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := NewTile(x*gd.TileWidth, y*gd.TileHeight, ground)
			tile.RandomizeImageIndex()
			tiles = append(tiles, tile)
		}
	}

	return tiles
}
