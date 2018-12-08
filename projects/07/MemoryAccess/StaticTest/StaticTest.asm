// constant
@111
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// constant
@333
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// constant
@888
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
@StaticTest.8
D=A
// POP
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
@StaticTest.3
D=A
// POP
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
@StaticTest.1
D=A
// POP
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// static
@StaticTest.3
D=M
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// static
@StaticTest.1
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
// static
@StaticTest.8
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
