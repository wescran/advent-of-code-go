package year2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	aoc "github.com/wescran/advent-of-code-go"
)

type Solution202306 struct{}

func (s *Solution202306) Solve() error {
	data, err := aoc.LoadInput(fmt.Sprintf("%s/%d/%02d.txt", aoc.YearsDir, 2023, 6))
	if err != nil {
		panic(err)
	}

	reNums := regexp.MustCompile(`\d+`)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	type Race struct {
		time     int
		distance int
	}
	races := make([]Race, 0)
	for _, line := range lines {
		before, after, _ := strings.Cut(line, ":")
		if before == "Time" {
			fullNum := ""
			for _, num := range reNums.FindAllString(after, -1) {
				//part 1
				//n, _ := strconv.Atoi(num)
				//races = append(races, Race{time: n})

				//part 2
				fullNum += num

			}
			n, _ := strconv.Atoi(fullNum)
			races = append(races, Race{time: n})
		}
		if before == "Distance" {
			fullNum := ""
			for _, num := range reNums.FindAllString(after, -1) {
				//part 1
				//n, _ := strconv.Atoi(num)
				//races[i].distance = n

				//part 2
				fullNum += num
			}
			n, _ := strconv.Atoi(fullNum)
			races[0].distance = n
		}
	}

	res := 1
	for _, race := range races {
		wins := 0
		//start half-way
		start := race.time / 2
		end := race.time - start
		dist := start * end
		for dist > race.distance {
			if start != end {
				wins += 2
			} else {
				wins += 1
			}
			start -= 1
			end += 1
			dist = start * end
		}
		res *= wins
	}
	fmt.Println(res)
	return nil
}

