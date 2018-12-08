// constant
@10
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
@LCL
D=M
@0
D=D+A
// POP
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// constant
@21
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// constant
@22
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
@ARG
D=M
@2
D=D+A
// POP
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
@ARG
D=M
@1
D=D+A
// POP
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// constant
@36
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
@THIS
D=M
@6
D=D+A
// POP
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// constant
@42
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// constant
@45
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
@THAT
D=M
@5
D=D+A
// POP
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
@THAT
D=M
@2
D=D+A
// POP
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// constant
@510
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
@R5
D=A
@6
D=D+A
// POP
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// local
@LCL
D=M
@0
A=D+A
D=M
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// that
@THAT
D=M
@5
A=D+A
D=M
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
// argument
@ARG
D=M
@1
A=D+A
D=M
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
// this
@THIS
D=M
@6
A=D+A
D=M
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// this
@THIS
D=M
@6
A=D+A
D=M
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
// SUB
@SP
AM=M-1
D=M
A=A-1
M=M-D
// temp
@R5
D=A
@6
A=D+A
D=M
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
