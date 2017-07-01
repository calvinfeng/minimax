package main

import (
    "fmt"
    "strconv"
)

func NewGame() Game {
	p1 := NewPlayer("Human", "X", false)
	p2 := NewPlayer("Computer", "O", true)
	return Game{
        PlayerOne: p1,
        PlayerTwo: p2,
        CurrentPlayer: p1,
        Board: NewBoard(),
        TurnNumber: 1,
    }
}

type Game struct {
	PlayerOne Player
	PlayerTwo Player
	CurrentPlayer Player
	Board Board
    TurnNumber int
}

func (game *Game) Start() {
    fmt.Println("\n___Welcome to Tic Tac Toe in Go___\n")
	for !game.Board.IsOver() {
        fmt.Println("Turn #" + strconv.Itoa(game.TurnNumber))
        fmt.Println(game.Board)
		fmt.Println("Current player:", game.CurrentPlayer)

		i, j := game.CurrentPlayer.GetMove(&game.Board)
		game.Board.PlaceMark(i, j, game.CurrentPlayer.GetMark())

		game.switchPlayer()
        game.TurnNumber += 1
	}
    fmt.Println(game.Board)
    fmt.Println("Game over!")
}

func (game *Game) switchPlayer() {
	if game.CurrentPlayer == game.PlayerOne {
		game.CurrentPlayer = game.PlayerTwo
	} else {
		game.CurrentPlayer = game.PlayerOne
	}
}
