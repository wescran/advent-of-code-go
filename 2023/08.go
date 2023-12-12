package year2023

import (
	"fmt"
	"regexp"
	"strings"

	aoc "github.com/wescran/advent-of-code-go"
)

type Solution202308 struct{}

type node struct {
	loc   string
	left  string
	right string
}

func (s *Solution202308) part1(moves string, nodes map[string]node) {
	steps := 0
	start := "AAA"
	end := "ZZZ"
	cur := start

	for cur != end {
		move := string(moves[steps%len(moves)])
		if move == "L" {
			cur = nodes[cur].left
		} else if move == "R" {
			cur = nodes[cur].right
		}
		steps++
	}
	fmt.Println(steps)
	return

}
func (s *Solution202308) doneMoves(nodeLocs map[string]string) bool {
	for _, loc := range nodeLocs {
		if !strings.HasSuffix(loc, "Z") {
			return false
		}
	}
	return true
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func (s *Solution202308) part2(moves string, nodes map[string]node, startNodes map[string]string) {
	steps := 0

	found := make(map[string]int)
	loops := make(map[string]int)

	for !s.doneMoves(startNodes) {
		for start, cur := range startNodes {
			move := string(moves[steps%len(moves)])
			if move == "L" {
				next := nodes[cur].left
				startNodes[start] = next
				if strings.HasSuffix(next, "Z") {
					if prev, ok := found[start]; ok {
						loops[start] = steps - prev
					}
					found[start] = steps
				}
			} else if move == "R" {
				next := nodes[cur].right
				startNodes[start] = next
				if strings.HasSuffix(next, "Z") {
					if prev, ok := found[start]; ok {
						loops[start] = steps - prev
					}
					found[start] = steps
				}
			}

		}
		if len(loops) == len(startNodes) {
			break
		}
		steps++
	}
	fmt.Println(loops)
	multiples := make([]int, 0)
	for _, count := range loops {
		multiples = append(multiples, count)
	}
	// with the loop steps, we find the LCM of step counts
	fmt.Println(LCM(multiples[0], multiples[1], multiples[2:]...))
	return
}

func (s *Solution202308) Solve() error {
	data, err := aoc.LoadInput(fmt.Sprintf("%s/%d/%02d.txt", aoc.YearsDir, 2023, 8))
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	movements := lines[0]
	nodes := make(map[string]node)
	startNodes := make(map[string]string)
	reLocs := regexp.MustCompile(`\w+`)
	for _, maps := range lines[1:] {
		locs := reLocs.FindAllString(maps, -1)
		if len(locs) > 0 {
			nodes[locs[0]] = node{
				loc:   locs[0],
				left:  locs[1],
				right: locs[2],
			}
			if strings.HasSuffix(locs[0], "A") {
				startNodes[locs[0]] = locs[0]
			}
		}
	}
	s.part1(movements, nodes)
	s.part2(movements, nodes, startNodes)
	return nil
}

