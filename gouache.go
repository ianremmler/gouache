package gouache

import "math/rand"

type Grid [][]int

type Game struct {
	numVals  int
	maxMoves int
	numMoves int
	curMove  int
	grid     []Grid
}

func New(numVals, maxMoves, cols, rows int) *Game {
	g := &Game{
		numVals:  numVals,
		maxMoves: maxMoves,
		grid:     make([]Grid, maxMoves+1),
	}
	for r := range g.grid {
		g.grid[r] = make(Grid, rows)
		for c := range g.grid[r] {
			g.grid[r][c] = make([]int, rows)
		}
	}
	g.Reset()
	return g
}

func (g *Game) NumVals() int {
	return g.numVals
}

func (g *Game) MaxMoves() int {
	return g.maxMoves
}

func (g *Game) CurMove() int {
	return g.curMove
}

func (g *Game) Reset() {
	g.curMove = 0
	for r := range g.grid[0] {
		for c := range g.grid[0][0] {
			g.grid[0][r][c] = rand.Int() % g.numVals
		}
	}
}

func (g *Game) fill(r, c, old, val int) {
	grid := g.grid[g.curMove]
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] != old {
		return
	}
	grid[r][c] = val
	g.fill(r-1, c, old, val)
	g.fill(r+1, c, old, val)
	g.fill(r, c-1, old, val)
	g.fill(r, c+1, old, val)
}

func (g *Game) Fill(val int) bool {
	if g.curMove >= g.maxMoves {
		return false
	}
	cur := g.grid[g.curMove]
	old := cur[0][0]
	if old == val {
		return false
	}
	g.curMove++
	g.numMoves = g.curMove
	nxt := g.grid[g.curMove]
	for r := range cur {
		for c := range cur[0] {
			nxt[r][c] = cur[r][c]
		}
	}
	g.fill(0, 0, old, val)
	return true
}

func (g *Game) Filled() bool {
	grid := g.grid[g.curMove]
	val := grid[0][0]
	for r := range grid {
		for c := range grid[0] {
			if grid[r][c] != val {
				return false
			}
		}
	}
	return true
}

func (g *Game) Grid() Grid {
	return g.grid[g.curMove]
}

func (g *Game) Undo() bool {
	if g.curMove <= 0 {
		return false
	}
	g.curMove--
	return true
}

func (g *Game) Redo() bool {
	if g.curMove >= g.numMoves {
		return false
	}
	g.curMove++
	return true
}
