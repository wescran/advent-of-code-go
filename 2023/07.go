package year2023

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	aoc "github.com/wescran/advent-of-code-go"
)

type Solution202307 struct{}

type Hand struct {
	cards    string
	bid      int
	strength int
}

var cardRanks = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}
var cardRanksJoker = map[string]int{
	"J": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"Q": 12,
	"K": 13,
	"A": 14,
}

func sortRanks1(a, b Hand) int {
	// a < b -1
	// a == b 0
	// a > b 1
	if a.strength < b.strength {
		return -1
	} else if a.strength > b.strength {
		return 1
	}
	// same strength
	for i := 0; i < 5; i++ {
		aSize := cardRanks[string(a.cards[i])]
		bSize := cardRanks[string(b.cards[i])]
		if aSize > bSize {
			return 1
		} else if aSize < bSize {
			return -1
		}
	}
	return 0

}
func sortRanks2(a, b Hand) int {
	// a < b -1
	// a == b 0
	// a > b 1
	if a.strength < b.strength {
		return -1
	} else if a.strength > b.strength {
		return 1
	}
	// same strength
	for i := 0; i < 5; i++ {
		aSize := cardRanksJoker[string(a.cards[i])]
		bSize := cardRanksJoker[string(b.cards[i])]
		if aSize > bSize {
			return 1
		} else if aSize < bSize {
			return -1
		}
	}
	return 0

}

func findStrength(hand string) int {
	chars := make(map[rune]int)
	for _, c := range hand {
		chars[c] += 1
	}
	counts := make([]int, 0)
	for _, count := range chars {
		counts = append(counts, count)
	}
	slices.Sort(counts)
	maxC := counts[len(counts)-1]
	switch maxC {
	case 5:
		return 7
	case 4:
		return 6
	case 3:
		//full house or 3 kind
		if counts[len(counts)-2] == 2 {
			return 5
		}
		return 4
	case 2:
		// 2 pair or 2 kind
		if counts[len(counts)-2] == 2 {
			return 3
		}
		return 2
	}
	return 1
}

func findStrengthJoker(hand string) int {
	chars := make(map[rune]int)
	splitHand := strings.Join(strings.Split(hand, "J"), "")
	jCount := 5 - len(splitHand)
	for _, c := range splitHand {
		chars[c] += 1
	}

	counts := make([]int, 0)
	for _, count := range chars {
		counts = append(counts, count)
	}
	slices.Sort(counts)
	maxC := 0
	if len(counts) > 0 {
		maxC += counts[len(counts)-1] + jCount
	} else {
		maxC += jCount
	}
	switch maxC {
	case 5:
		return 7
	case 4:
		return 6
	case 3:
		//full house or 3 kind
		if counts[len(counts)-2] == 2 {
			return 5
		}
		return 4
	case 2:
		// 2 pair or 2 kind
		if counts[len(counts)-2] == 2 {
			return 3
		}
		return 2
	}
	return 1
}

func part1(lines []string) error {
	hands := make([]Hand, 0)
	for _, line := range lines {
		hand, bid, found := strings.Cut(line, " ")
		if found {
			bidNum, _ := strconv.Atoi(bid)
			strength := findStrength(hand)
			hands = append(hands, Hand{hand, bidNum, strength})
		}
	}
	slices.SortFunc(hands, sortRanks1)
	winnings := 0
	for i, hand := range hands {
		winnings += (i + 1) * hand.bid
	}
	fmt.Println(winnings)
	return nil
}

func part2(lines []string) error {
	hands := make([]Hand, 0)
	for _, line := range lines {
		hand, bid, found := strings.Cut(line, " ")
		if found {
			bidNum, _ := strconv.Atoi(bid)
			strength := findStrengthJoker(hand)
			hands = append(hands, Hand{hand, bidNum, strength})
		}
	}
	slices.SortFunc(hands, sortRanks2)
	winnings := 0
	for i, hand := range hands {
		winnings += (i + 1) * hand.bid
	}
	fmt.Println(winnings)
	return nil
}

func (s *Solution202307) Solve() error {
	data, err := aoc.LoadInput(fmt.Sprintf("%s/%d/%02d.txt", aoc.YearsDir, 2023, 7))
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	part1(lines)
	part2(lines)

	return nil
}
