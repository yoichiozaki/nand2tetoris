// constant
@3030
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
@THIS
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
// constant
@3040
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
@THAT
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
// constant
@32
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
@THIS
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
@46
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
@THAT
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
// pointer
@THIS
D=M
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// pointer
@THAT
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
// this
@THIS
D=M
@2
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
// that
@THAT
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
