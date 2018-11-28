// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)
// R0, R1, R2 is addresses.
// Put your code here.

// @value means "load the value to the A register."
// and a value on the memory referenced by M is the one which address is the value A register holds.

//
// int RAM[3] = 0
// while (i - RAM[1] < 0) {
// 		RAM[2] += RAM[0];
// 		RAM[3]++;
// }
//

// initialize R2
@R2 			// A = R2
M = 0 			// RAM[2] = 0

@R3 			// A = R3
M = 0 			// RAM[3] = 0

(LOOP)
	@R3 		// A = R3
	D = M 		// D = RAM[3]
	@R1 		// A = R1
	D = D - M 	// RAM[3] = RAM[3] - RAM[1]
	@END 		// A = &END
	D; JGE 		// if D (= RAM[3] - RAM[1]) >= 0 then jump to END
	@R0 		// A = R0
	D = M 		// D = RAM[0]
	@R2 		// A = R2
	M = M + D 	// RAM[2] = RAM[2] + RAM[0]
	@R3 		// A = R3
	M = M + 1 	// RAM[R3] = RAM[R3]  + 1
	@LOOP 		// A = &LOOP
	0; JMP 		// jump to LOOP
(END)
	@END 		// A = &END
	0; JMP 		// jump to END