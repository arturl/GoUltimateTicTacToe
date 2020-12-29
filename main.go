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

type position struct {
	x byte
	y int
}

func IsGameWon(b *board, v BoardValue) bool {
	// Diagonals:
	if b.NW.Captured == v && b.C.Captured == v && b.SE.Captured == v { return true }
	if b.SW.Captured == v && b.C.Captured == v && b.NE.Captured == v { return true }

	// Horizontals:
	if b.NW.Captured == v && b.N.Captured == v && b.NE.Captured == v { return true }
	if b.W.Captured  == v && b.C.Captured == v && b.E.Captured  == v { return true }
	if b.SW.Captured == v && b.S.Captured == v && b.SE.Captured == v { return true }

	// Verticals:
	if b.NW.Captured == v && b.W.Captured == v && b.SW.Captured == v { return true }
	if b.N.Captured  == v && b.C.Captured == v && b.S.Captured  == v { return true }
	if b.NE.Captured == v && b.E.Captured == v && b.SE.Captured == v { return true }

	return false
}

func main() {
	b := board{}

	fmt.Println(b.GetAt('D', 4))

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

		fmt.Printf("Computer has %d moves: ", len(computer_moves))
		for _, cmv := range computer_moves {
			fmt.Printf("%c%d ", cmv.x, cmv.y)
		}
		fmt.Printf("\n")

		if len(computer_moves) == 0 {
			DrawBoard(&b)
			fmt.Printf("No more moves, game over\n")
			break
		}

		// No intelligence: just pick the first one
		computer_move = computer_moves[0]
		fmt.Printf("Computer played: %c%d\n", computer_move.x, computer_move.y)
		b.SetAt(computer_move.x, computer_move.y, O)

		if IsGameWon(&b, O) {
			DrawBoard(&b)
			fmt.Printf("Computer won. Game over\n")
			break
		}
	}
}