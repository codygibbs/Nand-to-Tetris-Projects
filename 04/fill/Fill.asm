// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

(LOOP)
	@KBD
	D=M

	@ON
	D;JGT

	@OFF
	D;JEQ

	// In case of timing issues
	@LOOP
	0;JMP

(ON)
	// Fill screen
	@0
	D=A-1

	@R1
	M=D

	@FILL
	0;JMP

(OFF)
	// Empty screen
	@0
	D=A

	@R1
	M=D

	@FILL
	0;JMP

(FILL)
	// R1 = fill value, R2 = current screen index (it counts backwards)
	@16384
	D=A
	@SCREEN
	D=D+A
	@R2
	M=D

	(FILL_LOOP)
		// Decrement index (R2)
		@R2
		MD=M-1

		// Draw pixel
		@R1
		D=M
		@R2
		M=D
		
		@SCREEN
		D=A
		@R2
		D=M-D
		@FILL_LOOP
		D;JEQ

	@LOOP
	0;JMP
