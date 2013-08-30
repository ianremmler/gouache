package main

import (
	"github.com/ianremmler/gouache"
	"github.com/wsxiaoys/terminal"
	"github.com/wsxiaoys/terminal/color"

	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	maxMoves = 25
	rows     = 14
	cols     = 14
)

var (
	colors = []string{"@{!rR}", "@{!gG}", "@{!bB}", "@{!cC}", "@{!mM}", "@{!yY}"}
	game    *gouache.Game
)

func main() {
	rand.Seed(time.Now().UnixNano())
	game = gouache.New(len(colors), maxMoves, rows, cols)
	play()
}

func play() {
	game.Reset()
	max := game.MaxMoves()
	for {
		cur := game.CurMove()
		terminal.Stdout.Clear()
		terminal.Stdout.Move(0, 0)
		fmt.Println("(#)paint (n)ew (q)uit (u)ndo (r)edo\n")
		printGrid()
		switch {
		case game.Filled():
			fmt.Println("\nYou win!")
		case cur >= max:
			fmt.Println("\nSorry, no more moves.")
		}
		fmt.Printf("\n%02d/%02d> ", cur, max)
		in := ""
		fmt.Scanf("%s", &in)
		in = strings.ToLower(strings.TrimSpace(in))
		switch in {
		case "n":
			game.Reset()
		case "u":
			game.Undo()
		case "r":
			game.Redo()
		case "q":
			os.Exit(0)
		default:
			if val, err := strconv.Atoi(in); err == nil {
				game.Fill(val - 1)
			}
		}
	}
}

func printGrid() {
	g := game.Grid()
	for r := range g {
		for c := range g[r] {
			val := g[r][c] + 1
			color.Print(colors[g[r][c]], fmt.Sprintf("%d ", val))
		}
		fmt.Println()
	}
}
