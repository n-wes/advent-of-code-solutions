package main
import (
    "fmt"
    "strings"
    "unicode"
    "strconv"
  )


/*
n: number lines

Time Complexity: O(n^2)
	Iterating over each line O(n)
	Then forward and reverse iteration per line O(6n) = O(n)
  		Check if digit O(1)
		Check if hashmap keys exist O(3) (worst case)

Space Complexity: O(n) 
  	Increasing substring size by 1 each iteration O(n)
*/
func main() {
  input := 
    `two1nine
    eightwothree
    abcone2threexyz
    xtwone3four
    4nineeightseven2
    zoneight234
    7pqrstsixteen`
  
  var lines = parseLines(input)
  
  var res = 0
  for _, line := range lines {
		i, _ := strconv.Atoi(findFirstDigit(line) + findLastDigit(line))
		res = i + res
  }
  fmt.Println(res)
}

func findFirstDigit (line string) (firstDigit string){
	dmap := map[string]string {
		"eno": "1",
		"owt": "2",
		"eerht": "3",
		"ruof": "4",
		"evif": "5",
		"xis": "6",
		"neves": "7",
		"thgie": "8",
		"enin": "9",
	}
	
	firstDigit = ""
	sublineReversed := ""
	
	for i, r := range line {
		// Number will always end search
		if unicode.IsDigit(r) {
			firstDigit = string(r)
			return
		}
		
		// Add new char to chars in reverse
		sublineReversed = string(r) + sublineReversed
		
		// We can use the fact that the range of string length for digit values are 3-5
		// "one" has a length of 3-5
		// "eight" has a length of 5
		if i >= 4 {
			if val, exists := dmap[sublineReversed[:5]]; exists {
				return val
			}
		}
		if i >= 3 {
			if val, exists := dmap[sublineReversed[:4]]; exists {
				return val
			}
		}
		if i >= 2 {
			if val, exists := dmap[sublineReversed[:3]]; exists {
				return val
			}
		}
	}
	return
}

func findLastDigit (line string) (lastDigit string){
	dmap := map[string]string {
		"one": "1",
		"two": "2",
		"three": "3",
		"four": "4",
		"five": "5",
		"six": "6",
		"seven": "7",
		"eight": "8",
		"nine": "9",
	}
	
	lastDigit = ""
	subline := ""
	
	for i := len(line) - 1; i >= 0; i-- {
		var r = rune(line[i])
		// Number will always end search
		if unicode.IsDigit(r) {
			lastDigit = string(r)
			return 
		}
		
		// Add new char to chars in reverse
		subline = string(r) + subline
		
		y := len(line) - i - 1
		
		// We can use the fact that the range of string length for digit values are 3-5
		// "one" has a length of 3-5
		// "eight" has a length of 5
		if y >= 4 {
			if val, exists := dmap[subline[:5]]; exists {
				return val
			}
		}
		if y >= 3 {
			if val, exists := dmap[subline[:4]]; exists {
				return val
			}
		}
		if y >= 2 {
			if val, exists := dmap[subline[:3]]; exists {
				return val
			}
		}
	}
	return
}

// Overhead to trim inputs
func parseLines(input string) (lines []string) {
    lines = strings.Split(input, "\n")
    for i, line := range lines {
        line = strings.TrimSpace(line)
        lines[i] = line
    }
    return lines
}