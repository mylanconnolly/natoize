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
		fmt.Println(PhoneticFast(arg))
	}
}

// Phonetic is the abstraction of the natoize logic.
func Phonetic(str string) (line string) {
	for _, r := range strings.ToLower(str) {
		nato, found := natoLetters[r]

		if !found {
			line += string(r)
		} else {
			line += nato
		}
		line += " "
	}
	return line
}

// PhoneticFast is the faster version of natoize, using a switch.
func PhoneticFast(str string) string {
	builder := strings.Builder{}

	for _, s := range str {
		switch s {
		case 'a', 'A':
			builder.WriteString("alfa ")
		case 'b', 'B':
			builder.WriteString("bravo ")
		case 'c', 'C':
			builder.WriteString("charlie ")
		case 'd', 'D':
			builder.WriteString("delta ")
		case 'e', 'E':
			builder.WriteString("echo ")
		case 'f', 'F':
			builder.WriteString("foxtrot ")
		case 'g', 'G':
			builder.WriteString("golf ")
		case 'h', 'H':
			builder.WriteString("hotel ")
		case 'i', 'I':
			builder.WriteString("india ")
		case 'j', 'J':
			builder.WriteString("juliet ")
		case 'k', 'K':
			builder.WriteString("kilo ")
		case 'l', 'L':
			builder.WriteString("lima ")
		case 'm', 'M':
			builder.WriteString("mike ")
		case 'n', 'N':
			builder.WriteString("november ")
		case 'o', 'O':
			builder.WriteString("oscar ")
		case 'p', 'P':
			builder.WriteString("papa ")
		case 'q', 'Q':
			builder.WriteString("quebec ")
		case 'r', 'R':
			builder.WriteString("romeo ")
		case 's', 'S':
			builder.WriteString("sierra ")
		case 't', 'T':
			builder.WriteString("tango ")
		case 'u', 'U':
			builder.WriteString("uniform ")
		case 'v', 'V':
			builder.WriteString("victor ")
		case 'w', 'W':
			builder.WriteString("whiskey ")
		case 'x', 'X':
			builder.WriteString("xray ")
		case 'y', 'Y':
			builder.WriteString("yankee ")
		case 'z', 'Z':
			builder.WriteString("zulu ")
		case '1':
			builder.WriteString("one ")
		case '2':
			builder.WriteString("two ")
		case '3':
			builder.WriteString("three ")
		case '4':
			builder.WriteString("four ")
		case '5':
			builder.WriteString("five ")
		case '6':
			builder.WriteString("six ")
		case '7':
			builder.WriteString("seven ")
		case '8':
			builder.WriteString("eight ")
		case '9':
			builder.WriteString("nine ")
		case '0':
			builder.WriteString("zero ")
		default:
			builder.WriteString(string(s) + " ")
		}
	}
	return builder.String()
}
