package main

import (
	"github.com/ianremmler/gouache"
	"github.com/wsxiaoys/terminal"
	"github.com/wsxiaoys/terminal/color"

	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	colorStr = [...]string{
		"@{!rR}", "@{!gG}", "@{!bB}", "@{!cC}", "@{!mM}", "@{!yY}", "@{!kK}", "@{!wW}",
	}
	maxMoves int
	rows     int
	cols     int
	colors   int
	game     *gouache.Game
)

func init() {
	flag.IntVar(&maxMoves, "moves", 25, "maximum number of moves")
	flag.IntVar(&rows, "rows", 14, "number of rows")
	flag.IntVar(&cols, "cols", 14, "number of columns")
	flag.IntVar(&colors, "colors", 6, "number of colors, max " + strconv.Itoa(len(colorStr)))
}

func main() {
	flag.Parse()
	if colors > len(colorStr) {
		colors = len(colorStr)
	}

	rand.Seed(time.Now().UnixNano())
	game = gouache.New(colors, maxMoves, rows, cols)
	play()
}

func play() {
	game.Reset()
	max := game.MaxMoves()
	for {
		cur := game.CurMove()
		terminal.Stdout.Clear()
		terminal.Stdout.Move(0, 0)
		fmt.Println("(#)paint (n)ew (q)uit (u)ndo (r)edo (0)rewind\n")
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
		case "0":
			game.Rewind()
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
			color.Print(colorStr[g[r][c]], fmt.Sprintf("%d ", val))
		}
		fmt.Println()
	}
}
