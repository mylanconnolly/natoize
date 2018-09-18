package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
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

func PhoneticFast(str string) string {
	builder := strings.Builder{}

	for _, s := range str {
		switch unicode.ToLower(s) {
		case 'a':
			builder.WriteString("alfa ")
		case 'b':
			builder.WriteString("bravo ")
		case 'c':
			builder.WriteString("charlie ")
		case 'd':
			builder.WriteString("delta ")
		case 'e':
			builder.WriteString("echo ")
		case 'f':
			builder.WriteString("foxtrot ")
		case 'g':
			builder.WriteString("golf ")
		case 'h':
			builder.WriteString("hotel ")
		case 'i':
			builder.WriteString("india ")
		case 'j':
			builder.WriteString("juliet ")
		case 'k':
			builder.WriteString("kilo ")
		case 'l':
			builder.WriteString("lima ")
		case 'm':
			builder.WriteString("mike ")
		case 'n':
			builder.WriteString("november ")
		case 'o':
			builder.WriteString("oscar ")
		case 'p':
			builder.WriteString("papa ")
		case 'q':
			builder.WriteString("quebec ")
		case 'r':
			builder.WriteString("romeo ")
		case 's':
			builder.WriteString("sierra ")
		case 't':
			builder.WriteString("tango ")
		case 'u':
			builder.WriteString("uniform ")
		case 'v':
			builder.WriteString("victor ")
		case 'w':
			builder.WriteString("whiskey ")
		case 'x':
			builder.WriteString("xray ")
		case 'y':
			builder.WriteString("yankee ")
		case 'z':
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
