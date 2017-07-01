package main

import "fmt"

// Since we defined pointer receiver for GetMove(), we need to return pointers to implement
// Player interface
func NewPlayer(name string, mark string, ai bool) Player {
    if ai {
        return &ComputerPlayer{name, mark}
    } else {
        return &HumanPlayer{name, mark}
    }
}

type Player interface {
    GetMove(b *Board) (int, int)
    GetMark() string
}

type ComputerPlayer struct {
    Name string
    Mark string
}

type HumanPlayer struct {
    Name string
    Mark string
}

func (hp *HumanPlayer) GetMove(b *Board) (i int, j int) {
	fmt.Print("Enter position: ")
	fmt.Scanf("%d %d", &i, &j)
	fmt.Println("Your input", i, j)
	return i, j
}

func (hp *HumanPlayer) GetMark() string {
    return hp.Mark
}

func (cp *ComputerPlayer) GetMove(b *Board) (i int, j int) {
    fmt.Println("Computer is thinking...")
    move := cp.minimax(b, cp.Mark, 1)
    i, j = move["i"], move["j"]
    return i, j
}

func (cp *ComputerPlayer) minimax(board *Board, mark string, depth int) map[string]int{
    if board.IsOver() {
        var score map[string]int
        score = make(map[string]int)
        if board.Winner() == cp.Mark {
            score["value"] = 10 - depth
        } else {
            score["value"] = depth - 10
        }
        return score
    }

    scores := []map[string]int{}
    for _, pos := range board.GetAvailablePos() {
        newBoard := board.GetCopy()
        i, j := pos[0], pos[1]
        newBoard.PlaceMark(i, j, mark)

        var score map[string]int
        if mark == "X" {
            score = cp.minimax(&newBoard, "O", depth + 1)
        } else {
            score = cp.minimax(&newBoard, "X", depth + 1)
        }
        score["i"] = i
        score["j"] = j
        scores = append(scores, score)
    }

    if mark == cp.Mark { // max
        maxScore := scores[0]
        for _, s := range scores {
            if maxScore["value"] < s["value"] {
                maxScore = s
            }
        }
        return maxScore
    } else { // min
        minScore := scores[0]
        for _, s := range scores {
            if minScore["value"] > s["value"] {
                minScore = s
            }
        }
        return minScore
    }
}

func (p *ComputerPlayer) GetMark() string {
    return p.Mark
}
