package main

import (
	"fmt"
	"unicode"
	"strconv"
)

type BoardValue int

const(
	Empty BoardValue = 0
	X = 1
	O = 2
)

type corner struct {
	Captured BoardValue
	NW BoardValue
	N  BoardValue
	NE BoardValue
	W  BoardValue
	C  BoardValue
	E  BoardValue
	SW BoardValue
	S  BoardValue
	SE BoardValue
}

type board struct {
	NW corner
	N  corner
	NE corner
	W  corner
	C  corner
	E  corner
	SW corner
	S  corner
	SE corner
}

func (b *board) clone () board {
	b2 := board{}
	b2 = *b
	return b2
}

type position struct {
	x byte
	y int
}

func main() {
	b := board{}
/*
	b.SetX("a1b1c1d1e1f1d4e4f4c4a5g4h6i7")
	b.SetO("g1h1i1a4b4c5h4a7b7c7d7e7f7g7h7")
*/
	computer_move := position{}
	for {
		DrawBoard(&b)
		valid := false;
		user_move := position{}
		for !valid {
			// Get user's move
			fmt.Printf("Enter you move (or 'q' to quit the game):")
			user_input := ""
			fmt.Scanln(&user_input)
			if user_input == "q" { return }
			if(len(user_input) == 2) {
				user_move.x = byte(user_input[0])
				user_move.y, _ = strconv.Atoi(user_input[1:])
				fmt.Printf("You entered: %c%d\n", user_move.x, user_move.y)
				user_move.x = byte(unicode.ToLower(rune(user_move.x)))
				valid, err := IsValidMove(&b, computer_move, user_move )
				if valid {
					break
				} else {
					fmt.Printf("Error: %s\n", err)
				}
			} else {
				fmt.Printf("Please enter your move such as e5 or a1\n")
			}
		}

		b.SetAt(user_move.x, user_move.y, X)

		if IsGameWon(&b, X) {
			DrawBoard(&b)
			fmt.Printf("Human won. Game over\n")
			break
		}

		// Get computer move

		computer_moves := FindAllMoves(&b, user_move)

		fmt.Printf("Computer has %d moves\n", len(computer_moves))
/*
		for _, mv := range computer_moves {
			fmt.Printf("%c%d ", mv.x, mv.y)
		}
		fmt.Printf("\n")
		return
*/
		for range computer_moves {
			fmt.Printf("x")
		}
		fmt.Printf("\n")

		if len(computer_moves) == 0 {
			DrawBoard(&b)
			fmt.Printf("No more moves, game over\n")
			break
		}

//		computer_move = FindBestMove(&b, computer_moves, O)
		computer_move, computer_score := negamax(&b, O, user_move, /*5*/4, 52)

		b.SetAt(computer_move.x, computer_move.y, O)
		fmt.Printf("Computer played: %c%d. Score:%d\n", computer_move.x, computer_move.y, computer_score)

		if IsGameWon(&b, O) {
			DrawBoard(&b)
			fmt.Printf("Computer won. Game over\n")
			break
		}
	}
}