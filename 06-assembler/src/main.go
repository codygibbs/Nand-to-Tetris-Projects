package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

//	func main() {
//		if len(os.Args) < 2 {
//			log.Fatal("not enough args")
//		}
//
//		program, err := NewProgramFromFile(os.Args[1])
//		if err != nil {
//			panic(err)
//		}
//
//		NewParser(program)
//	}

var (
	reDigit = regexp.MustCompile(`\d+`)
)

func assemble(p Parser, out *bufio.Writer) {
	encoder := Code{}

	var code uint16

	for p.HasMoreCommands() {
		p.Advance()

		cType := p.CommandType()
		if cType == C_COMMAND {
			comp := encoder.Comp(p.Comp())
			dest := encoder.Dest(p.Dest())
			jump := encoder.Jump(p.Jump())

			code = 0b1110_0000_0000_0000
			code |= uint16(comp) << 6
			code |= uint16(dest) << 3
			code |= uint16(jump)

			out.Write([]byte(fmt.Sprintf("%b\n", code)))
		} else if cType == A_COMMAND {
			var value int

			if reDigit.Match([]byte(p.Symbol())) {
				var err error
				value, err = strconv.Atoi(p.Symbol())
				if err != nil {
					panic(err)
				}
			} else {
				if _, ok := p.symbols[p.Symbol()]; !ok {
					p.symbols[p.Symbol()] = p.symbolPlace
					p.symbolPlace++
				}
				value = 16 + p.symbols[p.Symbol()]
			}

			out.Write([]byte(fmt.Sprintf("%016b\n", value)))
		} else if cType == L_COMMAND {
			if _, ok := p.symbols[p.Symbol()]; !ok {
				p.symbols[p.Symbol()] = p.symbolPlace
				p.symbolPlace++
			}
		}

		err := out.Flush()
		if err != nil {
			panic(err)
		}
	}
}
