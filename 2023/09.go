package year2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	aoc "github.com/wescran/advent-of-code-go"
)

type Solution202309 struct{}

func (s *Solution202309) allZero(nums []int) bool {
	for _, i := range nums {
		if i != 0 {
			return false
		}
	}
	return true
}

func (s *Solution202309) findDiff(nums []int) int {
	if s.allZero(nums) {
		return 0
	}
	next := 1
	newNums := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		newNums = append(newNums, nums[next]-nums[i])
		next++
	}
	return s.findDiff(newNums) + nums[len(nums)-1]
}
func (s *Solution202309) findSub(nums []int) int {
	if s.allZero(nums) {
		return 0
	}
	next := 1
	newNums := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		newNums = append(newNums, nums[next]-nums[i])
		next++
	}
	return nums[0] - s.findSub(newNums)
}

func (s *Solution202309) part1(series [][]int) {
	sum := 0
	for _, nums := range series {
		sum += s.findDiff(nums)
	}
	fmt.Println(sum)
	return
}
func (s *Solution202309) part2(series [][]int) {
	sum := 0
	for _, nums := range series {
		sum += s.findSub(nums)
	}
	fmt.Println(sum)
	return
}

func (s *Solution202309) Solve() error {
	data, err := aoc.LoadInput(fmt.Sprintf("%s/%d/%02d.txt", aoc.YearsDir, 2023, 9))
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	reNum := regexp.MustCompile(`[-]?\d+`)
	series := make([][]int, 0)
	for _, line := range lines {
		nums := make([]int, 0)
		for _, num := range reNum.FindAllString(line, -1) {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
		}
		series = append(series, nums)
	}
	s.part1(series)
	s.part2(series)

	return nil
}

