package main

import "testing"

func TestCodeDest (t *testing.T) {
    c := Code{}

    got := c.Dest("AM")
    if got != 0b_0000_0101 {
        t.Fatal("Wrong dest code. Got:", got)
    }
}

func TestCodeComp (t *testing.T) {
    c := Code{}

    got := c.Comp("D-1")
    if got != 0b_0000_1110 {
        t.Fatal("Wrong comp code. Got:", got)
    }
}

func TestCodeJump (t *testing.T) {
    c := Code{}

    got := c.Jump("JMP")
    if got != 0b_0000_0111 {
        t.Fatal("Wrong jump code. Got:", got)
    }
}

