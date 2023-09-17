package game

import (
	"fmt"
	"math/rand"
	"time"
)

type Coordinates struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
}

type Cell struct {
	Coordindates Coordinates `json:"coordindates,omitempty"`
	Alive        bool        `json:"alive,omitempty"`
}

func (c *Cell) getNeighbor(gameState *GameState, coordinates Coordinates) *Cell {
	x, y := coordinates.X, coordinates.Y

	if x < 0 || x >= len(gameState.Grid) || y < 0 || y >= len(gameState.Grid) {
		return nil
	}

	return &gameState.Grid[y][x]
}

func (c *Cell) LiveNeighbors(gameState *GameState) int {
	neighborCoordinates := []Coordinates{
		{X: c.Coordindates.X - 1, Y: c.Coordindates.Y},
		{X: c.Coordindates.X + 1, Y: c.Coordindates.Y},
		{X: c.Coordindates.X, Y: c.Coordindates.Y - 1},
		{X: c.Coordindates.X, Y: c.Coordindates.Y + 1},
		{X: c.Coordindates.X - 1, Y: c.Coordindates.Y - 1},
		{X: c.Coordindates.X + 1, Y: c.Coordindates.Y - 1},
		{X: c.Coordindates.X - 1, Y: c.Coordindates.Y + 1},
		{X: c.Coordindates.X + 1, Y: c.Coordindates.Y + 1},
	}

	liveNeighbors := 0

	for _, coord := range neighborCoordinates {
		neighbor := c.getNeighbor(gameState, coord)

		if neighbor != nil && neighbor.Alive == true {
			liveNeighbors++
		}
	}

	return liveNeighbors
}

type GameState struct {
	Grid [100][100]Cell `json:"grid,omitempty"`
}

type Game struct {
	State   GameState `json:"state,omitempty"`
	Running bool      `json:"running,omitempty"`
	Ticks   int       `json:"ticks,omitempty"`
}

func Init() *Game {
	var grid [100][100]Cell

	for y, row := range grid {
		for x := range row {
			grid[y][x] = Cell{
				Coordindates: Coordinates{
					X: x,
					Y: y,
				},
				Alive: rand.Intn(99)+1 >= 90,
			}
		}
	}

	initialState := GameState{Grid: grid}
	return &Game{State: initialState}
}

func (g *Game) Start() {
	g.Running = true
	go g.loop()
}

func (g *Game) Stop() {
	g.Running = false
}

func (g *Game) Print() {
	for _, row := range g.State.Grid {
		rowText := ""

		for _, cell := range row {
			if cell.Alive == true {
				rowText += "x"
			} else {
				rowText += "-"
			}
		}

		fmt.Println(rowText)
	}

	fmt.Printf("\n--------------------------------------------------------------------------- %d\n", g.Ticks)
}

func (g *Game) loop() {
	newState := g.State

	for g.Running == true {
		g.Ticks += 1
		g.Print()

		liveCells := 0

		for y, row := range g.State.Grid {
			for x := range row {
				cell := g.State.Grid[y][x]
				newCell := &newState.Grid[y][x]
				neighbors := cell.LiveNeighbors(&g.State)

				if cell.Alive {
					if neighbors < 2 || neighbors > 3 {
						newCell.Alive = false
					}
				} else {
					if neighbors == 3 {
						newCell.Alive = true
					}
				}

				if newCell.Alive {
					liveCells++
				}
			}
		}

		if liveCells == 0 {
			g.Stop()
		}

		g.State = newState
		time.Sleep(time.Second * 1)
	}
}
