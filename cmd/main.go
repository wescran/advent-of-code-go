package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	aoc "github.com/wescran/advent-of-code-go"
	year2023 "github.com/wescran/advent-of-code-go/2023"
)

func main() {
	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	runYearFlag := runCmd.Int("year", time.Now().Year(), "which year for advent of code")
	runDayFlag := runCmd.Int("day", 1, "which day for advent of code year")
	newCmd := flag.NewFlagSet("new", flag.ExitOnError)
	newYearFlag := newCmd.Int("year", 0, "which year for advent of code")
	newDayFlag := newCmd.Int("day", 0, "which day for advent of code")

	if len(os.Args) < 2 {
		fmt.Println("expected subcommand, exiting")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "new":
		newCmd.Parse(os.Args[2:])
		if *newYearFlag == 0 || *newDayFlag == 0 {
			fmt.Println("need to specify year and day for new solution file")
			os.Exit(1)
		}
		if time.Date(*newYearFlag, time.December, *newDayFlag, 0, 0, 0, 0, time.FixedZone("UTC-5", -5*60*60)).After(time.Now()) {
			fmt.Println("This year/day has not happened yet for AoC, please wait for puzzle to unlock")
			return
		}
		err := aoc.CreateSolutionFile(*newYearFlag, *newDayFlag)
		if err != nil {
			panic(err)
		}
		err = aoc.GetInput(*newYearFlag, *newDayFlag)
		if err != nil {
			panic(err)
		}
	case "run":
		runCmd.Parse(os.Args[2:])
		if *runYearFlag == 0 || *runDayFlag == 0 {
			fmt.Println("need to specify year and day to run solution file")
			os.Exit(1)
		}
		if time.Date(*runYearFlag, time.December, *runDayFlag, 0, 0, 0, 0, time.FixedZone("UTC-5", -5*60*60)).After(time.Now()) {
			fmt.Println("This year/day has not happened yet for AoC, please wait for puzzle to unlock")
			return
		}
		switch *runYearFlag {
		case 2023:
			year2023.Register()
		default:
			fmt.Println("no solutions found")
		}
		if solver, ok := aoc.Solutions[*runYearFlag][*runDayFlag]; ok {
			err := solver.Solve()
			if err != nil {
				panic(err)
			}
			return
		}
		fmt.Println("no solution found for year/day")
		return

	default:
		fmt.Println("expected subcommand, exiting")
		os.Exit(1)
	}
}
