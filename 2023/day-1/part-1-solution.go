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
	Iterating over each line O(n), then forward and reverse iteration per line O(2n)
Space Complexity: O(1)
*/
func main() {
	var input = 
		`1abc2
		pqr3stu8vwx
		a1b2c3d4e5f
		treb7uchet`
		
	var lines = parseLines(input)
	
	var res = 0
	for _, line := range lines {
		i, _ := strconv.Atoi(findFirstDigit(line) + findLastDigit(line))
		res = i + res
	}

	fmt.Println(res)
}

func findFirstDigit (line string) (firstDigit string){
	firstDigit = ""
	for _, r := range line {
		if unicode.IsDigit(r) {
			firstDigit = string(r)
			return
		}
	}
  return
}

func findLastDigit (line string) (lastDigit string){
	lastDigit = ""
	for i := len(line) - 1; i >= 0; i-- {
		var r = rune(line[i])
		if unicode.IsDigit(r) {
			lastDigit = string(r)
			return 
		}
	}
	return ""
}

// Overhead to trim inputs using bing chat
func parseLines(input string) (lines []string) {
    // Split the input into a slice of lines
    lines = strings.Split(input, "\n")
    // Loop over the lines and trim each one
    for i, line := range lines {
        // Remove leading and trailing whitespace
        line = strings.TrimSpace(line)
        // Update the slice with the trimmed line
        lines[i] = line
    }
    // Return the slice of trimmed lines
    return lines
}