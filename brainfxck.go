package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type (
	Token int
)

const (
	Pinc Token = iota
	Pdec
	Inc
	Dec
	Write
	Read
	Start
	End
)

var (
	tokens = map[Token]string{
		Pinc:  ">",
		Pdec:  "<",
		Inc:   "+",
		Dec:   "-",
		Write: ".",
		Read:  ",",
		Start: "[",
		End:   "]",
	}
)

func main() {
	src := " +++++++++[->++++++++>+++++++++++>+++++<<<]>.>++.+++++++..+++.>-.--------- ---.<++++++++.--------.+++.------.--------.>+.  "
	parsed := parse(splitSpace(addSpace(src)))
	run(parsed)
}

func addSpace(str string) string {
	spaced := str
	for _, val := range tokens {
		spaced = strings.Replace(spaced, val, " "+val+" ", -1)
	}
	return spaced
}

func splitSpace(str string) []string {
	return strings.Split(str, " ")
}

func parse(str []string) []Token {
	parsed := make([]Token, 0)
	for _, s := range str {
		switch s {
		case tokens[Pinc]:
			parsed = append(parsed, Pinc)
		case tokens[Pdec]:
			parsed = append(parsed, Pdec)
		case tokens[Inc]:
			parsed = append(parsed, Inc)
		case tokens[Dec]:
			parsed = append(parsed, Dec)
		case tokens[Write]:
			parsed = append(parsed, Write)
		case tokens[Read]:
			parsed = append(parsed, Read)
		case tokens[Start]:
			parsed = append(parsed, Start)
		case tokens[End]:
			parsed = append(parsed, End)
		}
	}
	return parsed
}

func run(tokens []Token) {
	memory := make([]int, 100)
	pointer := 0
	jumpPoint := 0
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case Pinc:
			pointer++
		case Pdec:
			pointer--
		case Inc:
			memory[pointer]++
		case Dec:
			memory[pointer]--
		case Write:
			fmt.Print(string(memory[pointer]))
		case Read:
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			memory[pointer] = int([]byte(input)[0])
		case Start:
			if memory[pointer] == 0 {
				for ; tokens[i] != End; i++ {
				}
			} else {
				jumpPoint = i
			}
		case End:
			i = jumpPoint - 1
		}
	}
}
