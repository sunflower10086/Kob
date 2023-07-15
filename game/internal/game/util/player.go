package util

import "strings"

type Player struct {
	Id       int
	BotId    int
	BotCode  string
	Sx       int
	Sy       int
	Steps    []int32
	Username string
	Photo    string
}

func checkTailIncreasing(step int) bool {
	if step <= 10 {
		return true
	} else {
		return step%3 == 1
	}
}

func (p *Player) GetStepsString() string {
	resp := strings.Builder{}

	for i := 0; i < len(p.Steps); i++ {
		resp.WriteByte(byte(p.Steps[i]))
	}

	return resp.String()
}

func (p *Player) GetCells() []Cell {
	res := make([]Cell, 0, 20)

	dx, dy := [4]int{-1, 0, 1, 0}, [4]int{0, 1, 0, -1}

	x, y := p.Sx, p.Sy

	res = append(res, Cell{x, y})

	for i, step := range p.Steps {
		x, y = x+dx[step], y+dy[step]
		res = append(res, Cell{x, y})
		if !checkTailIncreasing(i) {
			res = res[1:]
		}
	}

	return res
}
