package day3

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Number struct {
	Value        int
	Line         int
	Start        int
	End          int
	IsPartNumber bool
}

var LEGEND = map[string]int{
	".": math.MinInt,
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func getNumbers(line string, lineCursor int) []*Number {
	numbers := []*Number{}
	foundNumber := false
	currentNumber := &Number{
		Line: lineCursor,
	}
	for i := 1; i < len(line); i++ {
		char := string(line[i])
		val, ok := LEGEND[char]
		if !ok {
			val = math.MaxInt
		}
		if 0 <= val && val <= 9 {
			// a number at this cursor position
			if !foundNumber {
				foundNumber = true
				// first digit for this number
				currentNumber.Start = i
				currentNumber.End = i
				currentNumber.Value = val
			} else {
				// second+ digit for this number
				currentNumber.End = i
				currentNumber.Value, _ = strconv.Atoi(line[currentNumber.Start : currentNumber.End+1])
			}
		} else {
			// not a number at this cursor position
			if foundNumber {
				numbers = append(numbers, currentNumber)
				foundNumber = false
				currentNumber = &Number{
					Line: lineCursor,
				}
			}
		}
	}
	return numbers
}

func flagPartNumbers(number *Number, paddedLines []string) *Number {
	for ln := number.Line - 1; ln <= number.Line+1; ln++ {
		lrPadded := fmt.Sprintf(".%s.", paddedLines[ln])
		l := lrPadded[number.Start-1 : number.End+2]
		// fmt.Printf("debugWindow: %+v\n", l)
		replacer := strings.NewReplacer("0", "", "1", "", "2", "", "3", "", "4", "", "5", "", "6", "", "7", "", "8", "", "9", "", ".", "")
		l = replacer.Replace(l)
		if len(l) > 0 {
			number.IsPartNumber = true
			break
		}
	}
	return number
}

func findNumbersInAround(center int, topMidBottomLines []string) []int {
	numbers := []int{}
	// window of width 9 with a "*" centered
	leftBoundary := max(center-3, 0)
	rightBoundary := min(len(topMidBottomLines[0]), center+6)
	for ln, l := range topMidBottomLines {
		window := l[leftBoundary:rightBoundary]
		if ln == 1 {
			for i := 0; i < 4; i++ {
				maybeNum := window[i:4]
				if val, err := strconv.Atoi(maybeNum); err == nil {
					// num on left found
					numbers = append(numbers, val)
					break
				}
			}
			for j := 9; j > 4; j-- {
				maybeNum := window[5:j]
				if val, err := strconv.Atoi(maybeNum); err == nil {
					// num on right found
					numbers = append(numbers, val)
					break
				}
			}
		} else {
			if string(window[4]) == "." {
				// a number ends at index 3, then a ".", then a number at index 5
				i := 4
				for {
					_, errI := strconv.Atoi(string(window[i-1]))
					if errI == nil {
						i = i - 1
						continue
					}
					numberSlice := window[i:4]
					_, err := strconv.Atoi(string(numberSlice))
					if err != nil {
						break
					}
					val, _ := strconv.Atoi(numberSlice)
					numbers = append(numbers, val)
					break
				}
				j := 4
				for {
					_, errJ := strconv.Atoi(string(window[j+1]))
					if errJ == nil {
						j = j + 1
						continue
					}
					numberSlice := window[5 : j+1]
					_, err := strconv.Atoi(string(numberSlice))
					if err != nil {
						break
					}
					val, _ := strconv.Atoi(numberSlice)
					numbers = append(numbers, val)
					break
				}
			} else {
				// number has a digit at index 4, it's the only one
				i, j := 4, 4
				for {
					_, errI := strconv.Atoi(string(window[i-1]))
					if errI == nil {
						i = i - 1
						continue
					}
					_, errJ := strconv.Atoi(string(window[j+1]))
					if errJ == nil {
						j = j + 1
						continue
					}
					numberSlice := window[i : j+1]
					val, _ := strconv.Atoi(numberSlice)
					numbers = append(numbers, val)
					break
				}
			}
		}
	}

	return numbers
}
