package main

import (
	"fmt"
)

func Opponent(v BoardValue) BoardValue {
	if v == O { return X }
	return O
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

func ScoreFromCapture(c *corner, v BoardValue) int {
	if c.Captured == v { return 1 }
	if c.Captured == Opponent(v) { return -1 }
	return 0
}

func eval(b *board, v BoardValue, occupied int) int {

	score := 0

	if occupied < 40 {

		if IsGameWon(b, v) { return 200 }
		if IsGameWon(b, Opponent(v)) { return -200 }

		score += ScoreFromCapture(&b.C, v) * 4

		score += ScoreFromCapture(&b.NW, v) * 3
		score += ScoreFromCapture(&b.NE, v) * 3
		score += ScoreFromCapture(&b.SW, v) * 3
		score += ScoreFromCapture(&b.SE, v) * 3

		score += ScoreFromCapture(&b.N, v) * 2
		score += ScoreFromCapture(&b.S, v) * 2
		score += ScoreFromCapture(&b.W, v) * 2
		score += ScoreFromCapture(&b.E, v) * 2
	} else {

		if IsGameWon(b, v) { score += 200 - occupied }
		if IsGameWon(b, Opponent(v)) { score -= 200 - occupied }
		//score += 81 - occupied
	}

	return score
}

// Max returns the larger of x or y.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func count_occupied_corner_positions(c* corner) int {
	output := 0
	if c.NW != Empty { output++ }
	if c.N != Empty { output++ }
	if c.NE != Empty { output++ }
	if c.W != Empty { output++ }
	if c.C != Empty { output++ }
	if c.E != Empty { output++ }
	if c.SW != Empty { output++ }
	if c.S != Empty { output++ }
	if c.SE != Empty { output++ }
	return output
}

func count_occupied_board_positions(b* board) int {
	output := 0
	output += count_occupied_corner_positions(&b.NW)
	output += count_occupied_corner_positions(&b.N)
	output += count_occupied_corner_positions(&b.NE)
	output += count_occupied_corner_positions(&b.W)
	output += count_occupied_corner_positions(&b.C)
	output += count_occupied_corner_positions(&b.E)
	output += count_occupied_corner_positions(&b.SW)
	output += count_occupied_corner_positions(&b.S)
	output += count_occupied_corner_positions(&b.SE)
	return output
}

func negamax(b* board, v BoardValue, prev position, depth int) (position, int) {
	occupied := count_occupied_board_positions(b)
	nm_score, nm_pos := negamax_worker(b, v, prev, depth*2, true, occupied, -999, 999);
	fmt.Printf("\n")
	return nm_pos, nm_score
}


func negamax_worker(b* board, v BoardValue, prev position, depth int, log bool, occupied int, alpha int, beta int) (int, position){
	var possible_moves []position
	early_out := depth == 0
	debug_level := 99; // 99 for never
	if !early_out {
		possible_moves = FindAllMoves(b, prev)
		if len(possible_moves) == 0 {
			if depth >= debug_level { fmt.Printf("[%d] early out = true\n", depth) }
			early_out = true;
		}
	}

	if early_out {
		score := eval(b, v, occupied);
		if depth >= debug_level { fmt.Printf("[%d] early out, score = %d\n", depth, score) }
		return score, position{};
	}

	if depth >= debug_level {
		fmt.Printf("[%d] possible_moves: ", depth)
		for _, mv := range possible_moves {
			fmt.Printf("%c%d ", mv.x, mv.y)
		}
		fmt.Printf("\n")
	}

	value := -999
	best_move := position{}
	for _, mv := range possible_moves {
		if log { fmt.Printf(".") }
		if depth >= debug_level { fmt.Printf("[%d] considering %c%d START\n", depth, mv.x, mv.y) }
		child := b.clone()
		child.SetAt(mv.x, mv.y, v)

		nm_score, _ := negamax_worker(&child, Opponent(v), mv, depth-1, false, occupied+1, -beta, -alpha)
		score := -nm_score

		if score > value {
			if depth >= debug_level { fmt.Printf("[%d] considering %c%d, adjusting best score from %d to %d\n", depth, mv.x, mv.y, value, score) }
			value = score
			best_move = mv
		}

		alpha = max(alpha, value);
		if alpha >= beta {
			if depth >= debug_level { fmt.Printf("[%d] considering %c%d PRUNED: %d >= %d\n", depth, mv.x, mv.y, alpha, beta) }
			break // cut-off
		}
		if depth >= debug_level { fmt.Printf("[%d] considering %c%d END. score = %d\n", depth, mv.x, mv.y, score) }
	}

	if depth >= debug_level { fmt.Printf("[%d] return %d, %c%d\n", depth, value, best_move.x, best_move.y) }
	return value, best_move
}

func FindBestMove(b *board, ps []position, v BoardValue) position {
	best_score := -999
	best_move := 0
	for i, mv := range ps {
		board_copy := b.clone()
		board_copy.SetAt(mv.x, mv.y, v)
		score := eval(&board_copy, v, 0)
		if score > best_score {
			best_score = score
			best_move = i
		}
	}
	return ps[best_move]
}