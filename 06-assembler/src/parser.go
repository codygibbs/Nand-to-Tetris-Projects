package main

import (
	"bytes"
	"os"
	"regexp"
)

type Parser struct {
	program     [][]byte
	currentLine int
	symbols     map[string]int
	symbolPlace int
}

func (p *Parser) HasMoreCommands() bool {
	return len(p.program) > (p.currentLine + 1)
}

func (p *Parser) Advance() {
	p.currentLine++
}

type Command = int

const (
	NON_COMMAND = Command(iota)
	A_COMMAND
	C_COMMAND
	L_COMMAND
)

var (
	RE_A_COMMAND = regexp.MustCompile(`@(\d+|\w+)`)
	RE_C_COMMAND = regexp.MustCompile(`^(?:\s*)(?:([DM])=)?([DM]?(?:[+-]?[ADM01]))(?:;((?:JMP)|(?:JGT)))?`)
	RE_L_COMMAND = regexp.MustCompile(`\((\w+)\)`)
)

func (p *Parser) CommandType() Command {
	if RE_A_COMMAND.Match(p.program[p.currentLine]) {
		return A_COMMAND
	} else if RE_L_COMMAND.Match(p.program[p.currentLine]) {
		return L_COMMAND
	} else if RE_C_COMMAND.Match(p.program[p.currentLine]) {
		return C_COMMAND
	}

	return NON_COMMAND
}

func (p *Parser) Symbol() string {
	cmdType := p.CommandType()

	if cmdType == A_COMMAND {
		matches := RE_A_COMMAND.FindSubmatch(p.program[p.currentLine])
		return string(matches[1])
	} else if cmdType == L_COMMAND {
		matches := RE_L_COMMAND.FindSubmatch(p.program[p.currentLine])
		return string(matches[1])
	}

	return ""
}

func (p *Parser) Dest() string {
	if p.CommandType() != C_COMMAND {
		return ""
	}

	matches := RE_C_COMMAND.FindSubmatch(p.program[p.currentLine])
	return string(matches[1])
}

func (p *Parser) Comp() string {
	if p.CommandType() != C_COMMAND {
		return ""
	}

	matches := RE_C_COMMAND.FindSubmatch(p.program[p.currentLine])
	return string(matches[2])
}

func (p *Parser) Jump() string {
	if p.CommandType() != C_COMMAND {
		return ""
	}

	matches := RE_C_COMMAND.FindSubmatch(p.program[p.currentLine])
	return string(matches[3])
}

func NewParserFromString(input string) *Parser {
	return &Parser{
		program:     bytes.Split([]byte(input), []byte("\n")),
		currentLine: -1,
		symbols:     map[string]int{},
		symbolPlace: 0,
	}
}

func NewParserFromFile(filename string) (*Parser, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Parser{
		currentLine: -1,
		program:     bytes.Split([]byte(content), []byte("\n")),
		symbols:     map[string]int{},
		symbolPlace: 1,
	}, nil
}
