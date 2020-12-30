package main

import (
	"fmt"
	"unicode"
)

func (b *board) GetPtr(x byte, y int) *BoardValue {
	x = byte(unicode.ToLower(rune(x)))
	switch x {
	case 'a':
		switch y {
		case 1:
			return &b.NW.NW;
		case 2:
			return &b.NW.W;
		case 3:
			return &b.NW.SW;
		case 4:
			return &b.W.NW;
		case 5:
			return &b.W.W;
		case 6:
			return &b.W.SW;
		case 7:
			return &b.SW.NW;
		case 8:
			return &b.SW.W;
		case 9:
			return &b.SW.SW;
		}
	case 'b':
		switch y {
		case 1:
			return &b.NW.N;
		case 2:
			return &b.NW.C;
		case 3:
			return &b.NW.S;
		case 4:
			return &b.W.N;
		case 5:
			return &b.W.C;
		case 6:
			return &b.W.S;
		case 7:
			return &b.SW.N;
		case 8:
			return &b.SW.C;
		case 9:
			return &b.SW.S;
		}
	case 'c':
		switch y {
		case 1:
			return &b.NW.NE;
		case 2:
			return &b.NW.E;
		case 3:
			return &b.NW.SE;
		case 4:
			return &b.W.NE;
		case 5:
			return &b.W.E;
		case 6:
			return &b.W.SE;
		case 7:
			return &b.SW.NE;
		case 8:
			return &b.SW.E;
		case 9:
			return &b.SW.SE;
		}
	case 'd':
		switch y {
		case 1:
			return &b.N.NW;
		case 2:
			return &b.N.W;
		case 3:
			return &b.N.SW;
		case 4:
			return &b.C.NW;
		case 5:
			return &b.C.W;
		case 6:
			return &b.C.SW;
		case 7:
			return &b.S.NW;
		case 8:
			return &b.S.W;
		case 9:
			return &b.S.SW;
		}
	case 'e':
		switch y {
		case 1:
			return &b.N.N;
		case 2:
			return &b.N.C;
		case 3:
			return &b.N.S;
		case 4:
			return &b.C.N;
		case 5:
			return &b.C.C;
		case 6:
			return &b.C.S;
		case 7:
			return &b.S.N;
		case 8:
			return &b.S.C;
		case 9:
			return &b.S.S;
		}
	case 'f':
		switch y {
		case 1:
			return &b.N.NE;
		case 2:
			return &b.N.E;
		case 3:
			return &b.N.SE;
		case 4:
			return &b.C.NE;
		case 5:
			return &b.C.E;
		case 6:
			return &b.C.SE;
		case 7:
			return &b.S.NE;
		case 8:
			return &b.S.E;
		case 9:
			return &b.S.SE;
		}
	case 'g':
		switch y {
		case 1:
			return &b.NE.NW;
		case 2:
			return &b.NE.W;
		case 3:
			return &b.NE.SW;
		case 4:
			return &b.E.NW;
		case 5:
			return &b.E.W;
		case 6:
			return &b.E.SW;
		case 7:
			return &b.SE.NW;
		case 8:
			return &b.SE.W;
		case 9:
			return &b.SE.SW;
		}
	case 'h':
		switch y {
		case 1:
			return &b.NE.N;
		case 2:
			return &b.NE.C;
		case 3:
			return &b.NE.S;
		case 4:
			return &b.E.N;
		case 5:
			return &b.E.C;
		case 6:
			return &b.E.S;
		case 7:
			return &b.SE.N;
		case 8:
			return &b.SE.C;
		case 9:
			return &b.SE.S;
		}
	case 'i':
		switch y {
		case 1:
			return &b.NE.NE;
		case 2:
			return &b.NE.E;
		case 3:
			return &b.NE.SE;
		case 4:
			return &b.E.NE;
		case 5:
			return &b.E.E;
		case 6:
			return &b.E.SE;
		case 7:
			return &b.SE.NE;
		case 8:
			return &b.SE.E;
		case 9:
			return &b.SE.SE;
		}
	default:
		panic(fmt.Sprintf("Bad coordinates: %c%d", x, y))
	}
	return &b.NW.NW
}

func (b *board) GetCorner(x byte, y int) *corner {
	x = byte(unicode.ToLower(rune(x)))

	// West
	if x <= 'c' {
		if y <= 3 { return &b.NW }
		if y <= 6 { return &b.W }
		return &b.SW
	}

	// East
	if x >= 'g'	{
		if y <= 3 { return &b.NE }
		if y <= 6 { return &b.E }
		return &b.SE
	}

	// North
	if y <= 3 { return &b.N }
	
	// South
	if y >= 7 { return &b.S }

	return &b.C
}


func (c *corner) IsOccupiedBy(v BoardValue) bool {
	// Diagonals:
	if c.NW == v && c.C == v && c.SE == v { return true }
	if c.SW == v && c.C == v && c.NE == v { return true }

	// Horizontals:
	if c.NW == v && c.N == v && c.NE == v { return true }
	if c.W  == v && c.C == v && c.E  == v { return true }
	if c.SW == v && c.S == v && c.SE == v { return true }

	// Verticals:
	if c.NW == v && c.W == v && c.SW == v { return true }
	if c.N  == v && c.C == v && c.S  == v { return true }
	if c.NE == v && c.E == v && c.SE == v { return true }

	return false
}

func (c *corner) Fill(v BoardValue) {
	c.NW = v
	c.N = v
	c.NE = v
	c.SW = v
	c.S = v
	c.SE = v
	c.W = v
	c.C = v
	c.E = v
}

func (b *board) GetAt(x byte, y int) BoardValue {
	return *b.GetPtr(x,y)
}

func (b *board) SetAt(x byte, y int, value BoardValue) {
	*b.GetPtr(x,y) = value
	c := b.GetCorner(x,y)
	if c.IsOccupiedBy(value) {
		c.Captured = value
		c.Fill(value)
	}
}

func (b *board) SetImpl(s string, v BoardValue) {
	for i := 0; i < len(s); i+=2 {
		x := byte(s[i])
		y := int(s[i+1] - '0')
		b.SetAt(x,y,v)
	}
}

func (b *board) SetX(s string) {
	b.SetImpl(s, X)
}

func (b *board) SetO(s string) {
	b.SetImpl(s, O)
}