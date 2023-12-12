package year2023

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	aoc "github.com/wescran/advent-of-code-go"
)

type Solution202311 struct{}

type Point struct {
	x int
	y int
}
type Pair struct {
	a int
	b int
}

func (s *Solution202311) part1(galMap []Point) {
	sum := 0
	pairMap := make(map[Pair]float64)
	check := 0
	for a, galA := range galMap {
		for b, galB := range galMap {
			if galA == galB {
				continue
			}
			_, foundA := pairMap[Pair{a,b}]
			_, foundB := pairMap[Pair{b,a}]
			if foundA || foundB {
				continue
			}
			dist := math.Abs(float64(galA.x)-float64(galB.x)) + math.Abs(float64(galA.y)-float64(galB.y))
			pairMap[Pair{a,b}] = dist
			pairMap[Pair{b,a}] = dist
			check += 2
			sum += int(dist)
		}
	}
	fmt.Println(len(pairMap) / 2)
	fmt.Println(sum)
	return
}

func (s *Solution202311) Solve() error {
	data, err := aoc.LoadInput(fmt.Sprintf("%s/%d/%02d.txt", aoc.YearsDir, 2023, 11))
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	reGal := regexp.MustCompile(`#`)
	galMap := make([]Point, 0)
	emptyMap := make(map[int]bool)
	expansions := 0
	expanY := make(map[int]int)
	for i := range lines[0] {
		emptyMap[i] = true
	}
	for y, line := range lines {
		gals := reGal.FindAllStringIndex(line, -1)
		if len(gals) > 0 {
			for _, indices := range gals {
				if _, found := emptyMap[indices[0]]; found {
					delete(emptyMap, indices[0])
				}
			}
		} else {
			expansions += 999999
		}
		expanY[y] = expansions
	}
	for y, line := range lines {
		expansion := expanY[y]
		x := 0
		for i, char := range line {
			if _, found := emptyMap[i]; found {
				x += 999999
			}
			if char == rune('#') {
				galMap = append(galMap, Point{x + i, y + expansion})
			}
		}
	}
	fmt.Println(len(galMap))
	s.part1(galMap)
	return nil
}

