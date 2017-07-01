package main

import "fmt"

func NewPlayer(name string, board Board) Player {
	return Player{name, board}
}

type Player struct {
	name string
	board Board
}

func (player *Player) GetMove() (i int, j int) {
	fmt.Print("Enter position: ")
	fmt.Scanf("%d %d", &i, &j)
	fmt.Println("Your input", i, j)
	return i, j
}