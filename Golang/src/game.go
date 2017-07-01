package main

import "fmt"

func NewGame() Game {
	board := NewBoard()
	p1 := NewPlayer("Human", board)
	p2 := NewPlayer("Computer", board)
	return Game{p1, p2, p1, board}
}

type Game struct {
	PlayerOne Player
	PlayerTwo Player
	CurrentPlayer Player
	Board Board
}

func (game *Game) Start() {
	for !game.Board.IsOver() {
		fmt.Println("Current player:", game.CurrentPlayer)
		i, j := game.CurrentPlayer.GetMove()

		if (game.CurrentPlayer == game.PlayerOne) {
			game.Board.PlaceMark(i, j, "X")
		} else {
			game.Board.PlaceMark(i, j, "O")
		}

		fmt.Println(game.Board)
		game.switchPlayer()
	}
}

func (game *Game) switchPlayer() {
	if game.CurrentPlayer == game.PlayerOne {
		game.CurrentPlayer = game.PlayerTwo
	} else {
		game.CurrentPlayer = game.PlayerOne
	}
}
