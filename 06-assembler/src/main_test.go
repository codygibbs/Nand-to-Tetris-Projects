package main

import (
	"bufio"
	"strings"
	"testing"
)

type AssemberTest struct {
	Source string
	Expect []string
}

func TestAssembler(t *testing.T) {
	tests := []AssemberTest{
		{
			Source: `M=1`,
			Expect: []string{"1110111111001000"},
		},
		{
			Source: `@i`,
			Expect: []string{"0000000000010000"},
		},
		{
			Source: `
// Add 1 + ... + 100
    @i
    M=1  // i=1
    @sum
    M=0  // sum=0
(LOOP)
    @i
    D=M
    @100
    D=D-A
    @END
    D;JGT
    @i
    D=M
    @sum
    M=D+M
    @i
    M=M+1
    @LOOP
    0;JMP
    (END)
@END
    0;JMP
            `,
			Expect: []string{
				"0000000000010000",
				"1110111111001000",
				"0000000000010001",
				"1110101010001000",
				"0000000000010000",
				"1111110000010000",
				"0000000001100100",

				"1110001100000001",
				"0000000000010010",
				"1110001100000001",
				"0000000000010000",
				"1111110000010000",
				"0000000000010001",
				"1111000010001000",
				"0000000000010000",
				"1111110111001000",

				"0000000000000100",
				"1110101010000111",
				"0000000000010010",
				"1110101010000111",
			},
		},
	}

	for _, test := range tests {
		sBuilder := strings.Builder{}
		parser := NewParserFromString(test.Source)
		output := bufio.NewWriter(&sBuilder)

		assemble(*parser, output)

		got := strings.Split(strings.Trim(sBuilder.String(), "\n"), "\n")
		for i, line := range test.Expect {
			if got[i] != line {
				t.Errorf("Line %d mismatch. Expected '%s', got '%s'.", i+1, line, got[i])
			}
		}
	}
}
