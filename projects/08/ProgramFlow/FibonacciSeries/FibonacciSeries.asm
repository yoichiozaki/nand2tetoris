@256
D=A
@SP
M=D
// call Sys.init 0
// function call
@SP
D=M
@R13
M=D
@RET.1
D=A
@SP
A=M
M=D
@SP
M=M+1
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
@R13
D=M
@1
D=D-A
@ARG
M=D
@SP
D=M
@LCL
M=D
@Sys.init
0;JMP
(RET.1)
0;JMP
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
// pointer
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
@0
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// that
@THAT
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
@1
D=A
// PUSH
@SP
A=M
M=D
@SP
M=M+1
// that
@THAT
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
@2
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
(FibonacciSeries$MAIN_LOOP_START)
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
@FibonacciSeries$COMPUTE_ELEMENT
D;JNE
@FibonacciSeries$END_PROGRAM
0;JMP
(FibonacciSeries$COMPUTE_ELEMENT)
// that
@THAT
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
@1
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
// that
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
// pointer
@THAT
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
// AND
@SP
AM=M-1
D=M
A=A-1
M=D+M
// pointer
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
@FibonacciSeries$MAIN_LOOP_START
0;JMP
(FibonacciSeries$END_PROGRAM)
