package main

import (
	"flag"
	"fmt"

	"github.com/t00mas/aoc2023-go/solutions/day1"
	"github.com/t00mas/aoc2023-go/solutions/day2"
	"github.com/t00mas/aoc2023-go/solutions/day3"
	"github.com/t00mas/aoc2023-go/solutions/day4"
	"github.com/t00mas/aoc2023-go/solutions/types"
)

const (
	LAST_COMPLETED_DAY = 4
	MAX_STARS          = 2
)

var (
	day  = flag.Int("day", 1, "which day to execute the solution for")
	star = flag.Int("star", 1, "which star to execute, for the given day")
)

func main() {
	flag.Parse()
	if *day > LAST_COMPLETED_DAY || *star > MAX_STARS {
		fmt.Printf("[Day %d]::[Star %d] Not available\n", *day, *star)
		return
	}
	var s types.Solver
	switch *day {
	case 1:
		s = &day1.Solution{}
	case 2:
		s = &day2.Solution{}
	case 3:
		s = &day3.Solution{}
	case 4:
		s = &day4.Solution{}
	}
	s.Init()
	res := s.Solve(*star)
	fmt.Printf("[Day %d]::[Star %d] Solution => %d\n", *day, *star, res)
}
