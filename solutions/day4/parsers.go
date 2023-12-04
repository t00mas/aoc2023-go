package day4

import (
	"strconv"
	"strings"
)

type Card struct {
	Index       int
	Matches     int
	Name        string
	Nums        []int
	PointValue  int
	WinningNums []int
}

func (c *Card) Parse(line string) {
	c.Nums = make([]int, 0)
	c.WinningNums = make([]int, 0)
	c.Matches = 0

	c.Name, line, _ = strings.Cut(line, ": ")
	idxs := strings.Split(c.Name, " ")
	c.Index, _ = strconv.Atoi(idxs[len(idxs)-1])

	wins, nums, _ := strings.Cut(line, " | ")
	c.WinningNums = extractNums(wins)
	c.Nums = extractNums(nums)
	c.calculateWinnings()
}

func (c *Card) calculateWinnings() {
	for _, wn := range c.WinningNums {
		for _, n := range c.Nums {
			if n == wn {
				c.Matches = c.Matches + 1
				if c.PointValue == 0 {
					c.PointValue = 1
				} else {
					c.PointValue = c.PointValue * 2
				}
			}
		}
	}
}

func extractNums(numStr string) []int {
	nums := []int{}
	for _, v := range strings.Split(numStr, " ") {
		v = strings.Trim(v, " ")
		iv, _ := strconv.Atoi(v)
		if iv == 0 {
			continue
		}
		nums = append(nums, iv)
	}
	return nums
}
