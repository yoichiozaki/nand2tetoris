// constant
@17
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// constant
@17
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// eq
@SP
AM=M-1
D=M
A=A-1
D=M-D
@EQ.true.1
D;JEQ
@SP
A=M-1
M=0
@EQ.after.1
0;JMP
(EQ.true.1)
@SP
A=M-1
M=-1
(EQ.after.1)
// constant
@17
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// constant
@16
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// eq
@SP
AM=M-1
D=M
A=A-1
D=M-D
@EQ.true.2
D;JEQ
@SP
A=M-1
M=0
@EQ.after.2
0;JMP
(EQ.true.2)
@SP
A=M-1
M=-1
(EQ.after.2)
// constant
@16
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// constant
@17
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// eq
@SP
AM=M-1
D=M
A=A-1
D=M-D
@EQ.true.3
D;JEQ
@SP
A=M-1
M=0
@EQ.after.3
0;JMP
(EQ.true.3)
@SP
A=M-1
M=-1
(EQ.after.3)
// constant
@892
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// constant
@891
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// lt
@SP
AM=M-1
D=M
A=A-1
D=M-D
@LT.true.4
D;JLT
@SP
A=M-1
M=0
@LT.after.4
0;JMP
(LT.true.4)
@SP
A=M-1
M=-1
(LT.after.4)
// constant
@891
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// constant
@892
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// lt
@SP
AM=M-1
D=M
A=A-1
D=M-D
@LT.true.5
D;JLT
@SP
A=M-1
M=0
@LT.after.5
0;JMP
(LT.true.5)
@SP
A=M-1
M=-1
(LT.after.5)
// constant
@891
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// constant
@891
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// lt
@SP
AM=M-1
D=M
A=A-1
D=M-D
@LT.true.6
D;JLT
@SP
A=M-1
M=0
@LT.after.6
0;JMP
(LT.true.6)
@SP
A=M-1
M=-1
(LT.after.6)
// constant
@32767
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// constant
@32766
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// gt
@SP
AM=M-1
D=M
A=A-1
D=M-D
@GT.true.7

D;JGT
@SP
A=M-1
M=0
@GT.after.7
0;JMP
(GT.true.7)
@SP
A=M-1
M=-1
(GT.after.7)
// constant
@32766
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// constant
@32767
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// gt
@SP
AM=M-1
D=M
A=A-1
D=M-D
@GT.true.8

D;JGT
@SP
A=M-1
M=0
@GT.after.8
0;JMP
(GT.true.8)
@SP
A=M-1
M=-1
(GT.after.8)
// constant
@32766
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// constant
@32766
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// gt
@SP
AM=M-1
D=M
A=A-1
D=M-D
@GT.true.9

D;JGT
@SP
A=M-1
M=0
@GT.after.9
0;JMP
(GT.true.9)
@SP
A=M-1
M=-1
(GT.after.9)
// constant
@57
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// constant
@31
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// constant
@53
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// AND
@SP
AM=M-1
D=M
A=A-1
M=D+M
// constant
@112
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// SUB
@SP
AM=M-1
D=M
A=A-1
M=M-D
// NEG
@SP
A=M-1
M=-M
// AND
@SP
AM=M-1
D=M
A=A-1
M=D&M
// constant
@82
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// OR
@SP
AM=M-1
D=M
A=A-1
M=D|M
// NOT
@SP
A=M-1
M=!M
