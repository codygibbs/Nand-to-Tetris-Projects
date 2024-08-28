package main

import (
	"testing"
)

type CommandTypeTest struct {
	Source string
	Expect Command
}

func TestParser(t *testing.T) {
	tests := []CommandTypeTest{
		{
			Source: "(hello)",
			Expect: L_COMMAND,
		},
		{
			Source: "@foo",
			Expect: A_COMMAND,
		},
		{
			Source: "D;JEQ",
			Expect: C_COMMAND,
		},
		{
			Source: "M=D+M",
			Expect: C_COMMAND,
		},
		{
			Source: "// Add 1 + ... + 100",
			Expect: NON_COMMAND,
		},
	}

	for i, test := range tests {
		p := buildParser(test.Source)
		p.Advance()

		got := p.CommandType()
		if got != test.Expect {
			t.Errorf("Test %d - Expected: (%d), got: (%d)", i, test.Expect, got)
		}
	}
}

type SymbolTest struct {
	Source string
	Expect string
}

func TestSymbol(t *testing.T) {
	tests := []SymbolTest{
		{
			Source: "(FOO)",
			Expect: "FOO",
		},
		{
			Source: "@bar",
			Expect: "bar",
		},
	}

	for _, test := range tests {
		p := buildParser(test.Source)
		p.Advance()

		got := p.Symbol()
		if got != test.Expect {
			t.Fatal("Expected:", test.Expect, "got:", got)
		}
	}
}

func TestDest(t *testing.T) {
	p := buildParser("M=D+M;JMP")
	p.Advance()

	if p.Dest() != "M" {
		t.Fatal("Didn't get dest 'M'. Got:", p.Dest())
	}
}

func TestComp(t *testing.T) {
	p := buildParser("M=D+M;JMP")
	p.Advance()

	if p.Comp() != "D+M" {
		t.Fatal("Didn't get dest 'D+M'. Got:", p.Comp())
	}

	p = buildParser("M=D-M;JMP")
	p.Advance()

	if p.Comp() != "D-M" {
		t.Fatal("Didn't get dest 'D-M'. Got:", p.Comp())
	}
}

func TestJump(t *testing.T) {
	p := buildParser("M=D+M;JMP")
	p.Advance()

	got := p.Jump()
	if got != "JMP" {
		t.Fatal("Didn't get jump 'JMP'. Got:", got)
	}
}

func buildParser(asm string) *Parser {
	p := NewParserFromString(asm)
	return p
}
