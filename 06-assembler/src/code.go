package main

type Code struct {
}

func (c *Code) Dest(mnemonic string) uint8 {
    var result uint8

    for _, l := range mnemonic {
        if l == 'A' {
            result |= 0b_100
        } else if l == 'D' {
            result |= 0b_010
        } else if l == 'M' {
            result |= 0b_001
        }
    }

    return result
}

func (c *Code) Comp(mnemonic string) uint8 {
    if mnemonic == "0" {
        return 0b_0010_1010
    } else if mnemonic == "1" {
        return 0b_0011_1111
    } else if mnemonic == "-1" {
        return 0b_0011_1010
    } else if mnemonic == "D" {
        return 0b_0000_1100
    } else if mnemonic == "A" {
        return 0b_0011_0000
    } else if mnemonic == "M" {
        return 0b_0111_0000
    } else if mnemonic == "!D" {
        return 0b_0000_1101
    } else if mnemonic == "!A" {
        return 0b_0011_0001
    } else if mnemonic == "!M" {
        return 0b_0111_0001
    } else if mnemonic == "-D" {
        return 0b_0000_1111
    } else if mnemonic == "-A" {
        return 0b_0011_0011
    } else if mnemonic == "-M" {
        return 0b_0111_0011
    } else if mnemonic == "D+1" {
        return 0b_0001_1111
    } else if mnemonic == "A+1" {
        return 0b_0011_0111
    } else if mnemonic == "M+1" {
        return 0b_0111_0111
    } else if mnemonic == "D-1" {
        return 0b_0000_1110
    } else if mnemonic == "A-1" {
        return 0b_0011_0010
    } else if mnemonic == "M-1" {
        return 0b_0111_0010
    } else if mnemonic == "D+A" {
        return 0b_0000_0010
    } else if mnemonic == "D+M" {
        return 0b_0100_0010
    } else if mnemonic == "D-A" {
        return 0b_0001_0011
    } else if mnemonic == "D-M" {
        return 0b_0101_0011
    } else if mnemonic == "A-D" {
        return 0b_0000_0111
    } else if mnemonic == "M-D" {
        return 0b_0100_0111
    } else if mnemonic == "D&A" {
        return 0b_0000_0000
    } else if mnemonic == "D&M" {
        return 0b_0100_0000
    } else if mnemonic == "D|A" {
        return 0b_0001_0101
    } else if mnemonic == "D|M" {
        return 0b_0101_0101
    }

    return 0
}

func (c *Code) Jump(mnemonic string) uint8 {
    if mnemonic == "JGT" {
        return 0b_0001
    } else if mnemonic == "JEQ" {
        return 0b_0010
    } else if mnemonic == "JGE" {
        return 0b_0011
    } else if mnemonic == "JLT" {
        return 0b_0100
    } else if mnemonic == "JNE" {
        return 0b_0101
    } else if mnemonic == "JLE" {
        return 0b_0110
    } else if mnemonic == "JMP" {
        return 0b_0111
    }

    return 0
}

