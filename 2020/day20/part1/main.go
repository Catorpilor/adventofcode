package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// type tile struct {
// 	id                    int
// 	matix                 [][]byte
// 	left, right, up, down stats
// }

// type stats struct {
// 	dotCount, pondCount int
// }

var monster = [3][20]bool{
	{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false},
	{true, false, false, false, false, true, true, false, false, false, false, true, true, false, false, false, false, true, true, true},
	{false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, false},
}

type Dir byte

const (
	Top Dir = iota
	Right
	Bottom
	Left
)

type Border [10]bool

func (b Border) Reverse() Border {
	var rb Border
	for i, v := range b {
		rb[len(b)-1-i] = v
	}
	return rb
}

type TileType byte

const (
	CornerT TileType = iota
	BorderT
	CenterT
)

type TileBorder struct {
	T *Tile
	D Dir
}

type Tile struct {
	ID   int
	D    [10][10]bool
	Type TileType
}

func NewTile(lines []string) (*Tile, error) {
	var t Tile
	var err error
	t.ID, err = strconv.Atoi(strings.TrimPrefix(strings.TrimSuffix(lines[0], ":"), "Tile "))
	if err != nil {
		return nil, err
	}

	for i, l := range lines[1:] {
		for j, c := range l {
			t.D[i][j] = c == '#'
		}
	}
	return &t, nil
}

func (t *Tile) Flip(d Dir) *Tile {
	var newData [10][10]bool
	for i, l := range t.D {
		for j, v := range l {
			switch d {
			case Left, Right:
				newData[i][len(newData)-1-j] = v
			case Top, Bottom:
				newData[len(newData)-1-i][j] = v
			}
		}
	}
	t.D = newData
	return t
}

func (t *Tile) Rotate(d Dir) *Tile {
	var newData [10][10]bool
	for i, l := range t.D {
		for j, v := range l {
			switch d {
			case Left:
				newData[len(newData)-1-j][i] = v
			case Right:
				newData[j][len(newData)-1-i] = v
			}
		}
	}
	t.D = newData
	return t
}

func (t *Tile) BoarderInt(d Dir) Border {
	var b Border
	l := len(t.D) - 1

	for i := 0; i <= l; i++ {
		var condition bool
		switch d {
		case Top:
			condition = t.D[0][i]
		case Bottom:
			condition = t.D[l][l-i]
		case Left:
			condition = t.D[l-i][0]
		case Right:
			condition = t.D[i][l]
		}
		if condition {
			b[i] = true
		}
	}
	return b
}

func (t *Tile) TileOn(d Dir, lookup map[Border][]TileBorder) (*TileBorder, Border) {
	v := t.BoarderInt(d)
	tbs := lookup[v]
	if len(tbs) == 1 {
		return nil, Border{}
	}
	tb := tbs[0]
	if tb.T == t {
		tb = tbs[1]
	}
	return &tb, v.Reverse()
}

func (t *Tile) Transform(tb *TileBorder, targetDir Dir, targetBorder Border) *Tile {
	from := tb.D
	to := targetDir
	if from == to {
	} else if (from+to)%2 == 0 {
		t.Flip(from)
	} else {
		rotateDir := Dir(posMod(byte(to-from), 4))
		t.Rotate(rotateDir)
	}

	if t.BoarderInt(to) != targetBorder {
		t.Flip((to + 1) % 2)
	}

	return t
}

func (t *Tile) String() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("Tile %d:\n", t.ID))

	for _, l := range t.D {
		for _, v := range l {
			c := '.'
			if v {
				c = '#'
			}
			s.WriteRune(c)
		}
		s.WriteRune('\n')
	}
	return s.String()
}

func ScanGroup(input io.Reader, f func([]string) error) error {
	scanner := bufio.NewScanner(input)
	var groupLines []string
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			if groupLines != nil {
				if err := f(groupLines); err != nil {
					return err
				}
			}
			groupLines = nil
		} else {
			groupLines = append(groupLines, s)
		}
	}
	if groupLines != nil {
		if err := f(groupLines); err != nil {
			return err
		}
	}

	return scanner.Err()
}

func main() {
	var part1, part2 int
	tiles := make([]*Tile, 0, 200)
	err := ScanGroup(os.Stdin, func(s []string) error {
		t, err := NewTile(s)
		if err != nil {
			return err
		}
		tiles = append(tiles, t)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	var sea [][]bool
	part1, sea = buildSea(tiles)
	part2 = moveAndFindMonsters(sea)
	fmt.Println(part1, part2)
	// scanner := bufio.NewScanner(os.Stdin)
	// store := make(map[int]*tile)
	// var tmp [][]byte
	// var id int
	// var l0, r0, up0, down0, n int
	// for scanner.Scan() {
	// 	txt := scanner.Text()
	// 	if txt == "" {
	// 		// copy
	// 		local := make([][]byte, len(tmp))
	// 		for i := range tmp {
	// 			local[i] = make([]byte, len(tmp[i]))
	// 			copy(local[i], tmp[i])
	// 		}
	// 		store[id].matix = local
	// 		store[id].left = stats{dotCount: l0, pondCount: n - l0}
	// 		store[id].right = stats{dotCount: r0, pondCount: n - r0}
	// 		store[id].up = stats{dotCount: up0, pondCount: n - up0}
	// 		store[id].down = stats{dotCount: down0, pondCount: n - down0}
	// 		l0, r0, up0, down0 = 0, 0, 0, 0
	// 		continue
	// 	}
	// 	n = len(txt)
	// 	if txt[n-1] == ':' {
	// 		// tile head
	// 		// extract no.
	// 		for i := n - 2; i >= 0; i-- {
	// 			if txt[i] == ' ' {
	// 				id, _ = strconv.Atoi(txt[i+1 : n-1])
	// 				store[id] = &tile{id: id}
	// 				break
	// 			}
	// 		}
	// 		tmp = nil
	// 	} else {
	// 		if len(tmp) == 0 {
	// 			for i := range txt {
	// 				if txt[i] == '.' {
	// 					up0++
	// 				}
	// 			}
	// 		}
	// 		if len(tmp) == n-1 {
	// 			for i := range txt {
	// 				if txt[i] == '.' {
	// 					down0++
	// 				}
	// 			}
	// 		}
	// 		tmp = append(tmp, []byte(txt))
	// 		if txt[0] == '.' {
	// 			l0++
	// 		}
	// 		if txt[n-1] == '.' {
	// 			r0++
	// 		}

	// 	}

	// }
	// // store[id].matix = tmp
	// // store[id].left = stats{dotCount: l0, pondCount: n - l0}
	// // store[id].right = stats{dotCount: r0, pondCount: n - r0}
	// // store[id].up = stats{dotCount: up0, pondCount: n - up0}
	// // store[id].down = stats{dotCount: down0, pondCount: n - down0}
	// if err := scanner.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "reading standard input:", err)
	// }
	// for k := range store {
	// 	fmt.Printf("key: %d, left: %v, right: %v, up: %v, bottom: %v\n", k, store[k].left, store[k].right,
	// 		store[k].up, store[k].down)
	// }
}

func buildSea(tiles []*Tile) (int, [][]bool) {
	lookup := make(map[Border][]TileBorder, 4*len(tiles))
	directions := [4]Dir{Top, Right, Bottom, Left}

	for _, t := range tiles {
		for _, d := range directions {
			b := t.BoarderInt(d)
			bR := b.Reverse()
			lookup[b] = append(lookup[b], TileBorder{T: t, D: d})
			lookup[bR] = append(lookup[bR], TileBorder{T: t, D: d})
		}
	}

	res := 1

	corners := make([]*Tile, 0, 4)
	var borders int

	for _, t := range tiles {
		var c int
		for _, d := range directions {
			if len(lookup[t.BoarderInt(d)]) == 1 {
				c++
			}
		}
		switch c {
		case 2:
			corners = append(corners, t)
			t.Type = CornerT
			res *= t.ID
		case 1:
			t.Type = BorderT
			borders++
		case 0:
			t.Type = CenterT
		}
	}

	seaTileSize := borders/4 + 2
	seaTiles := make([][]*Tile, seaTileSize)

	var x, y int

	var currentTile *Tile

lineLoop:
	for {
		y = 0
		var startWithCorner bool
		seaTiles[x] = make([]*Tile, 0, seaTileSize)
		for {
			if x == 0 && y == 0 {
				currentTile = corners[0]
				if len(lookup[currentTile.BoarderInt(Top)]) != 1 {
					currentTile.Flip(Top)
				}
				if len(lookup[currentTile.BoarderInt(Left)]) != 1 {
					currentTile.Flip(Left)
				}
			} else if y == 0 {
				above, targetBorder := seaTiles[x-1][0].TileOn(Bottom, lookup)
				currentTile = above.T
				currentTile.Transform(above, Top, targetBorder)
			} else {
				left, targetBorder := seaTiles[x][y-1].TileOn(Right, lookup)
				currentTile = left.T
				currentTile.Transform(left, Left, targetBorder)
			}
			seaTiles[x] = append(seaTiles[x], currentTile)

			if currentTile.Type == CornerT {
				startWithCorner = true
			}

			if x != 0 && y != 0 && currentTile.Type == CornerT {
				break lineLoop
			}

			// when we're at a corner or border on the right, we move to another line
			if y != 0 && (currentTile.Type == CornerT || (!startWithCorner && currentTile.Type == BorderT)) {
				break
			}
			y++
		}
		x++
	}

	seaSize := len(seaTiles[0][0].D) - 2
	sea := make([][]bool, len(seaTiles)*seaSize)

	for i, l := range seaTiles {
		for _, t := range l {
			for ii, ll := range t.D {
				if ii == 0 || ii == len(ll)-1 {
					continue
				}
				for jj, v := range ll {
					if jj == 0 || jj == len(ll)-1 {
						continue
					}
					sea[i*seaSize+ii-1] = append(sea[i*seaSize+ii-1], v)
				}
			}
		}
	}

	return res, sea
}

func moveAndFindMonsters(sea [][]bool) int {
	v := findMonsters(sea)
	if v != -1 {
		return v
	}
	newSea := flip(sea, Left)
	v = findMonsters(newSea)
	if v != -1 {
		return v
	}
	newSea = flip(newSea, Top)
	v = findMonsters(newSea)
	if v != -1 {
		return v
	}
	newSea = flip(sea, Top)
	v = findMonsters(newSea)
	if v != -1 {
		return v
	}

	rotatedSea := rotate(sea)
	v = findMonsters(rotatedSea)
	if v != -1 {
		return v
	}
	newSea = flip(rotatedSea, Left)
	v = findMonsters(newSea)
	if v != -1 {
		return v
	}
	newSea = flip(newSea, Top)
	v = findMonsters(newSea)
	if v != -1 {
		return v
	}
	newSea = flip(rotatedSea, Top)
	v = findMonsters(newSea)
	if v != -1 {
		return v
	}
	return -1
}

func findMonsters(sea [][]bool) int {
	L := len(sea[0])
	H := len(sea)

	monsterSea := make([][]bool, H)
	for i := range monsterSea {
		monsterSea[i] = make([]bool, L)
	}

	var monsters int

	for i := 0; i < H-len(monster)+1; i++ {
	jLoop:
		for j := 0; j < L-len(monster[0])+1; j++ {
			// not optimal, but good enough
			for ii, ll := range monster {
				for jj, v := range ll {
					if v && !sea[i+ii][j+jj] {
						continue jLoop
					}
				}
			}
			monsters++
			for ii, ll := range monster {
				for jj, v := range ll {
					monsterSea[i+ii][j+jj] = v
				}
			}

		}
	}

	if monsters == 0 {
		return -1
	}
	// fmt.Println("Sea")
	// print(sea)
	// fmt.Println("Monsters")
	// print(monsterSea)

	var count int

	for i, l := range sea {
		for j, v := range l {
			if v && !monsterSea[i][j] {
				count++
			}
		}
	}

	return count

}

func flip(sea [][]bool, d Dir) [][]bool {
	newSea := make([][]bool, len(sea))
	for i := range newSea {
		newSea[i] = make([]bool, len(sea[0]))
	}
	for i, l := range sea {
		for j, v := range l {
			switch d {
			case Left, Right:
				newSea[i][len(newSea[0])-1-j] = v
			case Top, Bottom:
				newSea[len(newSea)-1-i][j] = v
			}
		}
	}
	return newSea
}

func rotate(sea [][]bool) [][]bool {
	newSea := make([][]bool, len(sea[0]))
	for i := range newSea {
		newSea[i] = make([]bool, len(sea))
	}
	for i, l := range sea {
		for j, v := range l {
			newSea[j][len(newSea[0])-1-i] = v
		}
	}
	return newSea
}

func posMod(a, m byte) byte {
	for a < 0 {
		a += m
	}
	return a % m
}

func print(sea [][]bool) {
	for _, l := range sea {
		for _, v := range l {
			c := '.'
			if v {
				c = '#'
			}
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}
