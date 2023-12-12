package year2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	aoc "github.com/wescran/advent-of-code-go"
)

type Solution202305 struct{}

func (s *Solution202305) Solve() error {
	data, err := aoc.LoadInput(fmt.Sprintf("%s/%d/%02d.txt", aoc.YearsDir, 2023, 5))
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	reNums := regexp.MustCompile(`\d+`)

	mapType := ""
	//part 1
	//seeds := make([]int, 0)

	type SeedMap struct {
		target int
		source int
		length int
	}
	seedSoil := make([]SeedMap, 0)
	soilFert := make([]SeedMap, 0)
	fertWater := make([]SeedMap, 0)
	waterLight := make([]SeedMap, 0)
	lightTemp := make([]SeedMap, 0)
	tempHum := make([]SeedMap, 0)
	humLoc := make([]SeedMap, 0)

	//part 2
	type Seeds struct {
		start  int
		length int
	}
	seedRange := make([]Seeds, 0)

	for _, line := range lines {
		before, after, found := strings.Cut(line, ":")
		// found map type
		if found {
			mapType = before
			//found seeds
			if after != "" {
				/* Part 1
				for _, num := range reNums.FindAllString(after, -1) {
					n, err := strconv.Atoi(num)
					if err != nil {
						return err
					}
					seeds = append(seeds, n)

				}*/
				s := Seeds{}
				for i, num := range reNums.FindAllString(after, -1) {
					n, _ := strconv.Atoi(num)
					if i%2 == 0 {
						s.start = n
						continue
					}
					s.length = n
					seedRange = append(seedRange, s)
					s = Seeds{}
				}
			}
			continue
		}
		//found empty line
		if before == "" {
			continue
		}

		switch mapType {
		case "seed-to-soil map":
			nums := reNums.FindAllString(before, -1)
			t, _ := strconv.Atoi(nums[0])
			s, _ := strconv.Atoi(nums[1])
			l, _ := strconv.Atoi(nums[2])
			seedSoil = append(seedSoil, SeedMap{t, s, l})
		case "soil-to-fertilizer map":
			nums := reNums.FindAllString(before, -1)
			t, _ := strconv.Atoi(nums[0])
			s, _ := strconv.Atoi(nums[1])
			l, _ := strconv.Atoi(nums[2])
			soilFert = append(soilFert, SeedMap{t, s, l})
		case "fertilizer-to-water map":
			nums := reNums.FindAllString(before, -1)
			t, _ := strconv.Atoi(nums[0])
			s, _ := strconv.Atoi(nums[1])
			l, _ := strconv.Atoi(nums[2])
			fertWater = append(fertWater, SeedMap{t, s, l})
		case "water-to-light map":
			nums := reNums.FindAllString(before, -1)
			t, _ := strconv.Atoi(nums[0])
			s, _ := strconv.Atoi(nums[1])
			l, _ := strconv.Atoi(nums[2])
			waterLight = append(waterLight, SeedMap{t, s, l})
		case "light-to-temperature map":
			nums := reNums.FindAllString(before, -1)
			t, _ := strconv.Atoi(nums[0])
			s, _ := strconv.Atoi(nums[1])
			l, _ := strconv.Atoi(nums[2])
			lightTemp = append(lightTemp, SeedMap{t, s, l})
		case "temperature-to-humidity map":
			nums := reNums.FindAllString(before, -1)
			t, _ := strconv.Atoi(nums[0])
			s, _ := strconv.Atoi(nums[1])
			l, _ := strconv.Atoi(nums[2])
			tempHum = append(tempHum, SeedMap{t, s, l})
		case "humidity-to-location map":
			nums := reNums.FindAllString(before, -1)
			t, _ := strconv.Atoi(nums[0])
			s, _ := strconv.Atoi(nums[1])
			l, _ := strconv.Atoi(nums[2])
			humLoc = append(humLoc, SeedMap{t, s, l})

		}
	}
	loc := 0
	//part 1
	//for _, seed := range seeds {
	for _, seedr := range seedRange {
		for seed := seedr.start; seed < seedr.start+seedr.length; seed++ {
			//get soil
			found := seed
			for _, seedMap := range seedSoil {
				//within range
				if found >= seedMap.source && found < seedMap.source+seedMap.length {
					found = found - seedMap.source + seedMap.target
					break
				}
			}
			//get fert
			for _, seedMap := range soilFert {
				//within range
				if found >= seedMap.source && found < seedMap.source+seedMap.length {
					found = found - seedMap.source + seedMap.target
					break
				}
			}
			//get water
			for _, seedMap := range fertWater {
				//within range
				if found >= seedMap.source && found < seedMap.source+seedMap.length {
					found = found - seedMap.source + seedMap.target
					break
				}
			}
			//get light
			for _, seedMap := range waterLight {
				//within range
				if found >= seedMap.source && found < seedMap.source+seedMap.length {
					found = found - seedMap.source + seedMap.target
					break
				}
			}
			//get temp
			for _, seedMap := range lightTemp {
				//within range
				if found >= seedMap.source && found < seedMap.source+seedMap.length {
					found = found - seedMap.source + seedMap.target
					break
				}
			}
			//get hum
			for _, seedMap := range tempHum {
				//within range
				if found >= seedMap.source && found < seedMap.source+seedMap.length {
					found = found - seedMap.source + seedMap.target
					break
				}
			}
			//get loc
			for _, seedMap := range humLoc {
				//within range
				if found >= seedMap.source && found < seedMap.source+seedMap.length {
					found = found - seedMap.source + seedMap.target
					break
				}
			}
			if loc == 0 || found < loc {
				loc = found
			}
		}
	}
	fmt.Println(loc)
	return nil

}
