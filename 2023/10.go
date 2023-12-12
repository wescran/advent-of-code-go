package year2023

import (
	"fmt"
	"strings"

	aoc "github.com/wescran/advent-of-code-go"
)

type Solution202310 struct{}

type SketchLoc struct {
	x int
	y int
}

var pipes map[string]map[string]SketchLoc = map[string]map[string]SketchLoc{
	"|": {
		"S": SketchLoc{0, 1},
		"N": SketchLoc{0, -1},
	},
	"-": {
		"E": SketchLoc{1, 0},
		"W": SketchLoc{-1, 0},
	},
	"L": {
		"S": SketchLoc{1, 0},
		"W": SketchLoc{0, -1},
	},
	"J": {
		"S": SketchLoc{-1, 0},
		"E": SketchLoc{0, -1},
	},
	"7": {
		"N": SketchLoc{-1, 0},
		"E": SketchLoc{0, 1},
	},
	"F": {
		"N": SketchLoc{1, 0},
		"W": SketchLoc{0, 1},
	},
}
var dirs map[string]map[string]string = map[string]map[string]string{
	"|": {
		"S": "S",
		"N": "N",
	},
	"-": {
		"E": "E",
		"W": "W",
	},
	"L": {
		"S": "E",
		"W": "N",
	},
	"J": {
		"S": "W",
		"E": "N",
	},
	"7": {
		"N": "W",
		"E": "S",
	},
	"F": {
		"N": "E",
		"W": "S",
	},
}

func (s *Solution202310) part1(sketch map[SketchLoc]string, start SketchLoc) map[SketchLoc]SketchLoc {
	//find surrounding pipes
	foundLocs := make([]SketchLoc, 0)
	foundDirs := make([]string, 0)
	if start.y > 0 {
		dir := "N"
		loc := SketchLoc{start.x, start.y - 1}
		p := sketch[loc]
		if foundPipe, ok := pipes[p]; ok {
			if _, ok := foundPipe[dir]; ok {
				foundLocs = append(foundLocs, loc)
				foundDirs = append(foundDirs, dir)
			}
		}
	}
	if start.y < 139 {
		dir := "S"
		loc := SketchLoc{start.x, start.y + 1}
		p := sketch[loc]
		if foundPipe, ok := pipes[p]; ok {
			if _, ok := foundPipe[dir]; ok {
				foundLocs = append(foundLocs, loc)
				foundDirs = append(foundDirs, dir)
			}
		}
	}
	if start.x > 0 {
		dir := "W"
		loc := SketchLoc{start.x - 1, start.y}
		p := sketch[loc]
		if foundPipe, ok := pipes[p]; ok {
			if _, ok := foundPipe[dir]; ok {
				foundLocs = append(foundLocs, loc)
				foundDirs = append(foundDirs, dir)
			}
		}
	}
	if start.x < 139 {
		dir := "E"
		loc := SketchLoc{start.x + 1, start.y}
		p := sketch[loc]
		if foundPipe, ok := pipes[p]; ok {
			if _, ok := foundPipe[dir]; ok {
				foundLocs = append(foundLocs, loc)
				foundDirs = append(foundDirs, dir)
			}
		}
	}
	locA, locB := foundLocs[0], foundLocs[1]
	stepsA := 1
	dirA, dirB := foundDirs[0], foundDirs[1]
	fmt.Println(locA, locB, dirA, dirB)
	route := make(map[SketchLoc]SketchLoc)
	route[start] = locA
	for locA != start {
		pipeA := sketch[locA]
		nextA := pipes[pipeA][dirA]
		route[locA] = SketchLoc{locA.x + nextA.x, locA.y + nextA.y}
		locA.x += nextA.x
		locA.y += nextA.y
		dirA = dirs[pipeA][dirA]
		stepsA += 1
	}
	fmt.Println(stepsA, stepsA/2)
	return route
}
func (s *Solution202310) part2(sketch map[SketchLoc]string, start SketchLoc, route map[SketchLoc]SketchLoc, size int) {
	// go clockwise around loop, taking anything to the LEFT
	// to be inside
	inside := map[string]map[string]SketchLoc{
		"left": {
			"S": SketchLoc{-1, 0},
			"N": SketchLoc{1, 0},
			"E": SketchLoc{0, 1},
			"W": SketchLoc{0, -1},
		},
		"right": {
			"S": SketchLoc{1, 0},
			"N": SketchLoc{-1, 0},
			"E": SketchLoc{0, -1},
			"W": SketchLoc{0, 1},
		},
	}
	nearTurns := map[string]map[string]map[string][]SketchLoc{
		"L": {
			"N": {
				"right": {SketchLoc{-1, 0}, SketchLoc{-1, 1}},
			},
			"E": {
				"left": {SketchLoc{-1, 0}, SketchLoc{-1, 1}},
			},
		},
		"J": {
			"N": {
				"left": {SketchLoc{0, 1}, SketchLoc{1, 1}},
			},
			"W": {
				"right": {SketchLoc{0, 1}, SketchLoc{1, 1}},
			},
		},
		"7": {
			"S": {
				"right": {SketchLoc{0, -1}, SketchLoc{1, -1}},
			},
			"W": {
				"left": {SketchLoc{0, -1}, SketchLoc{1, -1}},
			},
		},
		"F": {
			"S": {
				"left": {SketchLoc{-1, -1}, SketchLoc{0, -1}},
			},
			"E": {
				"right": {SketchLoc{-1, -1}, SketchLoc{0, -1}},
			},
		},
	}
	side := "right"
	curLoc := route[start]
	curPipe := sketch[curLoc]
	curDir := dirs[curPipe]["S"]
	curInsideDiff := inside[side][curDir]
	curInside := SketchLoc{curLoc.x + curInsideDiff.x, curLoc.y + curInsideDiff.y}
	foundInside := make(map[SketchLoc]int)
	for curLoc != start {
		//if not a pipe piece
		if _, found := route[curInside]; !found {
			foundInside[curInside] += 1
		}
		if extraLocs, found := nearTurns[curPipe][curDir][side]; found {
			for _, l := range extraLocs {
				loc := SketchLoc{curLoc.x + l.x, curLoc.y + l.y}
				if _, f := route[loc]; !f {
					foundInside[loc] += 1
				}
			}
		}
		curLoc = route[curLoc]
		curPipe = sketch[curLoc]
		curDir = dirs[curPipe][curDir]
		curInsideDiff = inside[side][curDir]
		curInside = SketchLoc{curLoc.x + curInsideDiff.x, curLoc.y + curInsideDiff.y}

	}
	for y := -1; y < 2; y++ {
		for x := -1; x < 2; x++ {
			l := SketchLoc{start.x + x, start.y + y}
			_, found := route[l]
			if !found {
				foundInside[l] += 1
			}
		}

	}
	fmt.Println(foundInside)
	fmt.Println(len(foundInside))

	for {
		added := 0
		for loc := range foundInside {
			for y := -1; y < 2; y++ {
				for x := -1; x < 2; x++ {
					l := SketchLoc{loc.x + x, loc.y + y}
					_, found := route[l]
					if !found {
						if _, found := foundInside[l]; !found {
							foundInside[l] += 1
							added += 1
						}
					}
				}
			}
		}
		if added == 0 {
			break
		}
	}
	for y := 0; y < 140; y++ {
		row := ""
		for x := 0; x < 140; x++ {
			if _, found := route[SketchLoc{x, y}]; found {
				row += sketch[SketchLoc{x, y}]
			} else if _, found := foundInside[SketchLoc{x, y}]; found {
				row += "#"
			} else {
				row += "."
			}
		}
		fmt.Println(row)
	}

	fmt.Println(len(foundInside))
	return

}
func (s *Solution202310) Solve() error {
	data, err := aoc.LoadInput(fmt.Sprintf("%s/%d/%02d.txt", aoc.YearsDir, 2023, 10))
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	sketch := make(map[SketchLoc]string)
	start := SketchLoc{}
	for y, line := range lines {
		for x, char := range line {
			sketch[SketchLoc{x, y}] = string(char)
			if char == rune('S') {
				start.x = x
				start.y = y
			}
		}
	}
	fmt.Println(len(lines), len(lines[0]))
	route := s.part1(sketch, start)
	s.part2(sketch, start, route, len(lines)*len(lines[0]))
	return nil
}

