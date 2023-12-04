package day1

import (
	"strconv"
	"strings"
)

func ProcessLine(line string, opt *ProcessingOptions) int {
	newLine := line
	if opt.ProcessNumWords {
		newLine = ReplaceNumWords(newLine)
	}

	var (
		first int
		last  int
	)
	for i := 0; i < len(newLine); i++ {
		if val, err := strconv.Atoi(string(newLine[i])); err == nil {
			if first <= 0 {
				first = val
				last = val
			} else {
				last = val
			}
		}
	}
	calibration := first*10 + last
	return calibration
}

func ReplaceNumWords(line string) string {
	compositeNums := map[string]string{
		"eightwo":   "82",
		"nineight":  "98",
		"oneight":   "18",
		"threeight": "38",
		"twone":     "21",
	}
	for k, v := range compositeNums {
		line = strings.ReplaceAll(line, k, v)
	}
	nums := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	for k, v := range nums {
		line = strings.ReplaceAll(line, k, v)
	}
	return line
}
