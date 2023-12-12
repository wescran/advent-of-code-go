package year2023

import aoc "github.com/wescran/advent-of-code-go"

func Register() {
	aoc.RegisterSolver(2023, map[int]aoc.Solver{
		1:  &Solution202301{},
		2:  &Solution202302{},
		3:  &Solution202303{},
		4:  &Solution202304{},
		5:  &Solution202305{},
		6:  &Solution202306{},
		7:  &Solution202307{},
		8:  &Solution202308{},
		9:  &Solution202309{},
		10: &Solution202310{},
		11: &Solution202311{},
		//12: &Solution202312{},
		//13: &Solution202313{},
		//14: &Solution202314{},
		//15: &Solution202315{},
		//16: &Solution202316{},
		//17: &Solution202317{},
		//18: &Solution202318{},
		//19: &Solution202319{},
		//20: &Solution202320{},
		//21: &Solution202321{},
		//22: &Solution202322{},
		//23: &Solution202323{},
		//24: &Solution202324{},
		//25: &Solution202325{},
	},
	)
}
