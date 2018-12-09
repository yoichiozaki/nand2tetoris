// constant
@0
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// local
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
(BasicLoop$LOOP_START)
// argument
@ARG
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
// AND
@SP
AM=M-1
D=M
A=A-1
M=D+M
// local
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
// argument
@ARG
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
// constant
@1
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
// argument
@ARG
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
// argument
@ARG
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
@SP
AM=M-1
D=M
@BasicLoop$LOOP_START
D;JNE
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
