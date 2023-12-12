package year2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	aoc "github.com/wescran/advent-of-code-go"
)

type Solution202303 struct{}

func (s *Solution202303) Solve() error {
	data, err := aoc.LoadInput(fmt.Sprintf("%s/%d/%02d.txt", aoc.YearsDir, 2023, 3))
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	//lines = []string{"467..114..", "...*......", "..35..633.", "......#...", "617*......", ".....+.58.", "..592.....", "......755.", "...$.*....", ".664.598.."}

	sum, ratio := 0, 0

	reNum := regexp.MustCompile(`\d+`)
	reSym := regexp.MustCompile(`[^\.\d]`)
	reGear := regexp.MustCompile(`\*`)

	type gKey struct {
		start int
		stop  int
	}
	type pKey struct {
		line  int
		start int
		stop  int
	}
	var gearLoc map[int][]int
	gears := make(map[pKey][]int)

	for y, line := range lines {
		numMatches := reNum.FindAllStringIndex(line, -1)
		for _, index := range numMatches {
			gearLoc = make(map[int][]int)
			start, stop := index[0], index[1]
			xMin, xMax := start-1, stop+1
			if xMin < 0 {
				xMin = start
			}
			if xMax > len(line) {
				xMax = stop
			}
			//above num
			var adjSlice string
			if y > 0 {
				yAbove := y - 1
				adj := lines[yAbove][xMin:xMax]
				gearMatch := reGear.FindAllStringIndex(adj, -1)
				for _, m := range gearMatch {
					gearLoc[m[0]] = []int{yAbove, xMin + m[0], xMin + m[1]}
				}
				adjSlice += adj
			}
			//beside num
			preLength := len(adjSlice)
			left := line[xMin:start]
			right := line[stop:xMax]
			if left == "*" {
				gearLoc[preLength] = []int{y, xMin, start}
			}
			if right == "*" {
				gearLoc[preLength+1] = []int{y, stop, xMax}
			}
			adjSlice += left + right

			if y < len(lines)-1 {
				yBelow := y + 1
				preLength = len(adjSlice)
				adj := lines[yBelow][xMin:xMax]
				gearMatch := reGear.FindAllStringIndex(adj, -1)
				for _, m := range gearMatch {
					gearLoc[preLength+m[0]] = []int{yBelow, xMin + m[0], xMin + m[1]}
				}
				adjSlice += adj
			}
			symMatch := reSym.FindAllStringIndex(adjSlice, -1)
			if symMatch != nil {
				n, err := strconv.Atoi(line[start:stop])
				if err != nil {
					return err
				}
				sum += n
				gearMatch := reGear.FindAllStringIndex(adjSlice, -1)
				fmt.Println(y, gearLoc, gearMatch)
				for _, i := range gearMatch {
					yGear := gearLoc[i[0]][0]
					startGear := gearLoc[i[0]][1]
					stopGear := gearLoc[i[0]][2]
					g := pKey{yGear, startGear, stopGear}
					gears[g] = append(gears[g], n)
				}

			}
		}

	}
	fmt.Println(sum)
	fmt.Println(ratio)
	finalSum := 0
	for _, v := range gears {
		if len(v) == 2 {
			finalSum += v[0] * v[1]
		}
	}
	fmt.Println(finalSum)
	return nil
}
