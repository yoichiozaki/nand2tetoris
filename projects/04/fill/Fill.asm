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

// Put your code here.
	@status 	// A = status
	M = -1		// RAM[status] = -1
	D = 0		// D = 0
	@SETSCREEN	// A = SETSCREEN
	0; JMP		// jump to SETSCREEN

(LOOP)
	@KBD		// A = KBD
	D = M 		// D = RAM[KBD]
	@SETSCREEN	// A = SETSCREEN
	D; JEQ 		// if D == 0 then jump to SETSCREEN
	D = -1		// D = -1

(SETSCREEN)
	@ARG 		// A = R2
	M = D 		// RAM[2] = D
	@status 	// A = status
	D = D - M 	// D = D - RAM[status]
	@LOOP 		// A = LOOP
	D; JEQ 		// if D == 0 then jump tp LOOP

	@ARG 		// A = R2
	D = M 		// D = RAM[2]
	@status 	// A = status
	M = D 		// RAM[status] = D

	@SCREEN 	// A = SCREEN
	D = A 		// D = SCREEN
	@8192 		// A = 8192
	D = D + A   // D = D + 8192
	@i 			// A = i
	M = D 		// RAM[i] = D

(SETLOOP)
	@i 			// A = i
	D = M - 1 	// D = RAM[i] - 1
	M = D 		// RAM[i]
	@LOOP 		// A = LOOP
	D; JLT 		// if D < 0 then jump LOOP

	@status 	// A = status
	D = M 		// D = RAM[status]
	@i 			// A = i
	A = M 		// A = RAM[i]
	M = D 		// RAM[i] = D
	@SETLOOP 	// A = SETLOOP
	0; JMP      // jump to SETLOOP