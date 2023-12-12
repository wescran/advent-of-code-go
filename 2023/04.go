package year2023

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	aoc "github.com/wescran/advent-of-code-go"
)

type Solution202304 struct{}

func hasMore(cards []int) bool {
	for _, i := range cards {
		if i > 0 {
			return true
		}
	}
	return false
}

func (s *Solution202304) Solve() error {
	data, err := aoc.LoadInput(fmt.Sprintf("%s/%d/%02d.txt", aoc.YearsDir, 2023, 4))
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	//lines = []string{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"}
	reGame := regexp.MustCompile(`\d+`)

	var won map[string]bool
	var score int
	total := float64(0)

	generates := make(map[int]int)
	cards := make([]int, len(lines))

	//part 1
	for g, line := range lines {
		score = 0
		gameDivide := strings.Index(line, ":")
		numsDivide := strings.Index(line, "|")
		winNums := line[gameDivide:numsDivide]
		won = make(map[string]bool)
		for _, num := range reGame.FindAllString(winNums, -1) {
			won[num] = true
		}
		myNums := line[numsDivide:]
		for _, num := range reGame.FindAllString(myNums, -1) {
			if _, ok := won[num]; ok {
				score += 1
			}
		}
		if score > 0 {
			total += math.Pow(float64(2), float64(score-1))
		}
		generates[g+1] = score
		cards[g] += 1
	}
	fmt.Println(total)

	//part 2
	processed := 0
	for i := 0; i < len(cards); i++ {
		processed += cards[i]
		if n, ok := generates[i+1]; ok {
			for j := 0; j < n; j++ {
				cards[i+j+1] += cards[i]
			}
		}
		cards[i] -= cards[i]
	}
	fmt.Println(processed)

	return nil
}
