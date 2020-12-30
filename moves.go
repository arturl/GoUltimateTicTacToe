package main

import (
	"fmt"
)

func IsValidMove(b *board, prev position, p position) (bool, string) {
	if p.x < 'a' || p.x > 'i' || p.y < 1 || p.y > 9 { return false, "coordinates out of bounds" }
	if b.GetAt(p.x, p.y) != Empty { return false, "position already occupied" }
	if prev.x != 0 {
		prev_corner, _, _ := GetSmallCorner(b, prev)
		if prev_corner.Captured != Empty {
			// The square is captured, so all moves are possible
			return true, ""
		}
		this_corner := b.GetCorner(p.x, p.y)
		if prev_corner != this_corner { return false, fmt.Sprintf("cannot move in this square because previous move was %c%d", prev.x, prev.y) }
	}
	return true, ""
}

func GetSmallCorner(b *board, prev position) (*corner, byte, int) {

	switch(prev) {
	case position{'a',1}, position{'d',1}, position{'g',1},
		 position{'a',4}, position{'d',4}, position{'g',4},
		 position{'a',7}, position{'d',7}, position{'g',7}:
		return &b.NW, 0, 0
	case position{'b',1}, position{'e',1}, position{'h',1},
		 position{'b',4}, position{'e',4}, position{'h',4},
		 position{'b',7}, position{'e',7}, position{'h',7}:
		return &b.N, 3, 0
	case position{'c',1}, position{'f',1}, position{'i',1},
		 position{'c',4}, position{'f',4}, position{'i',4},
		 position{'c',7}, position{'f',7}, position{'i',7}:
		return &b.NE, 6, 0
	case position{'a',2}, position{'d',2}, position{'g',2},
		 position{'a',5}, position{'d',5}, position{'g',5},
		 position{'a',8}, position{'d',8}, position{'g',8}:
		return &b.W, 0, 3
	case position{'b',2}, position{'e',2}, position{'h',2},
		 position{'b',5}, position{'e',5}, position{'h',5},
		 position{'b',8}, position{'e',8}, position{'h',8}:
		return &b.C, 3, 3
	case position{'c',2}, position{'f',2}, position{'i',2},
		 position{'c',5}, position{'f',5}, position{'i',5},
		 position{'c',8}, position{'f',8}, position{'i',8}:
		return &b.E, 6, 3
	case position{'a',3}, position{'d',3}, position{'g',3},
		 position{'a',6}, position{'d',6}, position{'g',6},
		 position{'a',9}, position{'d',9}, position{'g',9}:
		return &b.SW, 0, 6
	case position{'b',3}, position{'e',3}, position{'h',3},
		 position{'b',6}, position{'e',6}, position{'h',6},
		 position{'b',9}, position{'e',9}, position{'h',9}:
		return &b.S, 3, 6
	case position{'c',3}, position{'f',3}, position{'i',3},
		 position{'c',6}, position{'f',6}, position{'i',6},
		 position{'c',9}, position{'f',9}, position{'i',9}:
		return &b.SE, 6, 6
	}
	panic(fmt.Sprintf("prev value bad: %c%d", prev.x, prev.y))
}

type positionDelta struct {
	corner
	xdelta byte
	ydelta int
}

func FindAllMoves(b *board, prev position) []position {
	output := []position{}

	if IsGameWon(b, O) || IsGameWon(b, X) { return output }

	var c *corner;
	var xdelta byte;
	var ydelta int;

	c, xdelta, ydelta = GetSmallCorner(b, prev)

	pd := positionDelta{*c, xdelta, ydelta}

	available_squares := []positionDelta{pd}

	if c.Captured != Empty {
		// The square is captured, so moves to other squares are possible
		if b.NW.Captured == Empty { available_squares = append(available_squares, positionDelta{b.NW, 0, 0}) }
		if b.N.Captured  == Empty { available_squares = append(available_squares, positionDelta{b.N,  3, 0}) }
		if b.NE.Captured == Empty { available_squares = append(available_squares, positionDelta{b.NE, 6, 0}) }
		if b.W.Captured  == Empty { available_squares = append(available_squares, positionDelta{b.W,  0, 3}) }
		if b.C.Captured  == Empty { available_squares = append(available_squares, positionDelta{b.C,  3, 3}) }
		if b.E.Captured  == Empty { available_squares = append(available_squares, positionDelta{b.E,  6, 3}) }
		if b.SW.Captured == Empty { available_squares = append(available_squares, positionDelta{b.SW, 0, 6}) }
		if b.S.Captured  == Empty { available_squares = append(available_squares, positionDelta{b.S,  3, 6}) }
		if b.SE.Captured == Empty { available_squares = append(available_squares, positionDelta{b.SE, 6, 6}) }
	}

	adjustPos := func(p position, xdelta byte, ydelta int) position {
		return position{ p.x + xdelta, p.y + ydelta }
	}

	for _, c := range available_squares {
		if c.NW == Empty {
			pos := position{'a',1}
			output = append(output, adjustPos(pos, c.xdelta, c.ydelta))
		}
		if c.N == Empty {
			pos := position{'b',1}
			output = append(output, adjustPos(pos, c.xdelta, c.ydelta))
		}
		if c.NE == Empty {
			pos := position{'c',1}
			output = append(output, adjustPos(pos, c.xdelta, c.ydelta))
		}
		if c.W == Empty {
			pos := position{'a',2}
			output = append(output, adjustPos(pos, c.xdelta, c.ydelta))
		}
		if c.C == Empty {
			pos := position{'b',2}
			output = append(output, adjustPos(pos, c.xdelta, c.ydelta))
		}
		if c.E == Empty {
			pos := position{'c',2}
			output = append(output, adjustPos(pos, c.xdelta, c.ydelta))
		}
		if c.SW == Empty {
			pos := position{'a',3}
			output = append(output, adjustPos(pos, c.xdelta, c.ydelta))
		}
		if c.S == Empty {
			pos := position{'b',3}
			output = append(output, adjustPos(pos, c.xdelta, c.ydelta))
		}
		if c.SE == Empty {
			pos := position{'c',3}
			output = append(output, adjustPos(pos, c.xdelta, c.ydelta))
		}
	}
	return output
}