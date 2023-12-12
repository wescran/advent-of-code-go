package year2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	aoc "github.com/wescran/advent-of-code-go"
)

type Solution202301 struct{}

func (s *Solution202301) Solve() error {
	data, err := aoc.LoadInput(fmt.Sprintf("%s/%d/%02d.txt", aoc.YearsDir, 2023, 1))
	if err != nil {
		return err
	}
	rNum := regexp.MustCompile(`\d{1}`)
	lines := strings.Split(string(data), "\n")
	sum := 0
	for _, v := range lines {
		found := rNum.FindAllString(v, -1)
		if len(found) == 0 {
			continue
		}
		first, last := found[0], found[len(found)-1]
		n, _ := strconv.Atoi(first + last)
		sum += n
	}
	fmt.Println(sum)

	rWordNum := regexp.MustCompile(`eightwo|twone|sevenine|nineight|eighthree|threeight|fiveight|oneight|one|two|three|four|five|six|seven|eight|nine|\d{1}`)
	numMap := map[string]string{
		"one":       "1",
		"two":       "2",
		"three":     "3",
		"four":      "4",
		"five":      "5",
		"six":       "6",
		"seven":     "7",
		"eight":     "8",
		"nine":      "9",
		"twone":     "21",
		"eightwo":   "82",
		"sevenine":  "79",
		"nineight":  "98",
		"eighthree": "83",
		"threeight": "38",
		"fiveight":  "58",
		"oneight":   "18",
	}
	sum = 0
	for _, v := range lines {
		found := rWordNum.FindAllString(v, -1)
		if len(found) == 0 {
			continue
		}
		first, last := found[0], found[len(found)-1]

		if len(first) > 1 {
			first = numMap[first][:1]
		}
		if len(last) > 1 {
			last = numMap[last]
			last = last[len(last)-1:]
		}
		i, _ := strconv.Atoi(first + last)
		sum += i
	}
	fmt.Println(sum)
	return nil
}
