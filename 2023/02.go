package year2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	aoc "github.com/wescran/advent-of-code-go"
)

type Solution202302 struct{}

func (s *Solution202302) Solve() error {
	data, err := aoc.LoadInput(fmt.Sprintf("%s/%d/%02d.txt", aoc.YearsDir, 2023, 2))
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	colourMatch := map[string]int{
		"r": 12,
		"g": 13,
		"b": 14,
	}
	sum := 0
	reNum := regexp.MustCompile(`(\d+ [grb])`)

Loop:
	for _, game := range lines {
		gSplit := strings.Split(game, ":")
		gNum, err := strconv.Atoi(strings.Split(gSplit[0], " ")[1])
		if err != nil {
			return err
		}
		parts := strings.Split(gSplit[1], ";")

		for _, p := range parts {
			matches := reNum.FindAllString(p, -1)
			for _, m := range matches {
				numClr := strings.Split(m, " ")
				if i, _ := strconv.Atoi(numClr[0]); colourMatch[numClr[1]] < i {
					continue Loop
				}
			}
		}
		sum += gNum

	}

	fmt.Println(sum)

	// Part 2
	sum = 0
	for _, game := range lines {
		gSplit := strings.Split(game, ":")
		parts := strings.Split(gSplit[1], ";")

		maxNum := map[string]int{
			"r": 0,
			"g": 0,
			"b": 0,
		}
		for _, p := range parts {
			matches := reNum.FindAllString(p, -1)
			for _, m := range matches {
				numClr := strings.Split(m, " ")
				if i, _ := strconv.Atoi(numClr[0]); i > maxNum[numClr[1]] {
					maxNum[numClr[1]] = i
				}
			}
		}
		product := 1
		for _, v := range maxNum {
			product *= v
		}
		sum += product

	}
	fmt.Println(sum)

	return nil

}
