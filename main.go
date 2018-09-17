package main

import (
	"fmt"
	"os"
	"strings"
)

var natoLetters = map[rune]string{
	'a': "alfa",
	'b': "bravo",
	'c': "charlie",
	'd': "delta",
	'e': "echo",
	'f': "foxtrot",
	'g': "golf",
	'h': "hotel",
	'i': "india",
	'j': "juliet",
	'k': "kilo",
	'l': "lima",
	'm': "mike",
	'n': "november",
	'o': "oscar",
	'p': "papa",
	'q': "quebec",
	'r': "romeo",
	's': "sierra",
	't': "tango",
	'u': "uniform",
	'v': "victor",
	'w': "whiskey",
	'x': "xray",
	'y': "yankee",
	'z': "zulu",
	'1': "one",
	'2': "two",
	'3': "three",
	'4': "four",
	'5': "five",
	'6': "six",
	'7': "seven",
	'8': "eight",
	'9': "nine",
	'0': "zero",
}

func main() {
	for _, arg := range os.Args[1:] {
		var line string

		for _, r := range strings.ToLower(arg) {
			nato, found := natoLetters[r]

			if !found {
				line += string(r)
			} else {
				line += nato
			}
			line += " "
		}
		fmt.Println(line)
	}
}
