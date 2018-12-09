package CodeWriter

import (
	"hackVMTranslator/Parser"
	"os"
	"strconv"
	"strings"
)

const (
	ADD = "// AND\n" +
		"@SP\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"A=A-1\n" +
		"M=D+M\n"
	SUB = "// SUB\n" +
		"@SP\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"A=A-1\n" +
		"M=M-D\n"
	NEG = "// NEG\n" +
		"@SP\n" +
		"A=M-1\n" +
		"M=-M\n"
	AND = "// AND\n" +
		"@SP\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"A=A-1\n" +
		"M=D&M\n"
	OR = "// OR\n" +
		"@SP\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"A=A-1\n" +
		"M=D|M\n"
	NOT = "// NOT\n" +
		"@SP\n" +
		"A=M-1\n" +
		"M=!M\n"
	PUSH = "// PUSH\n" +
		"@SP\n" +
		"A=M\n" +
		"M=D\n" +
		"@SP\n" +
		"M=M+1\n"
	POP = "// POP\n" +
		"@R13\n" +
		"M=D\n" +
		"@SP\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"@R13\n" +
		"A=M\n" +
		"M=D\n"
	RETURN = "// RETURN\n" +
		// *(LCL - 5) -> R13
		"@LCL\n" +
		"D=M\n" +
		"@5\n" +
		"A=D-A\n" +
		"D=M\n" +
		"@R13\n" +
		"M=D\n" +
		// *(SP - 1) -> *ARG
		"@SP\n" +
		"A=M-1\n" +
		"D=M\n" +
		"@ARG\n" +
		"A=M\n" +
		"M=D \n" +
		// ARG + 1 -> SP
		"D=A+1\n" +
		"@SP\n" +
		"M=D\n" +
		// *(LCL - 1) -> THAT; LCL--
		"@LCL\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"@THAT\n" +
		"M=D\n" +
		// *(LCL - 1) -> THIS; LCL--
		"@LCL\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"@THIS\n" +
		"M=D\n" +
		// *(LCL - 1) -> ARG; LCL--
		"@LCL\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"@ARG\n" +
		"M=D\n" +
		// *(LCL - 1) -> LCL
		"@LCL\n" +
		"A=M-1\n" +
		"D=M\n" +
		"@LCL\n" +
		"M=D\n" +
		// R13 -> A
		"@R13\n" +
		"A=M\n" +
		"0;JMP\n"
)

type codeWriter struct {
	count   int
	parser  *Parser.Parser
	asmFile *os.File
}

func New() *codeWriter {
	return &codeWriter{
		count:   0,
		parser:  nil,
		asmFile: nil,
	}
}

func (cw *codeWriter) SetParser(parser *Parser.Parser) {
	cw.parser = parser
	asmFile, err := os.Create(strings.Trim(parser.GetVMFileName(), ".vm") + ".asm")
	if err != nil {
		panic(err)
	}
	cw.asmFile = asmFile
}

func (cw *codeWriter) WriteBootstrap() {
	_, err := cw.asmFile.Write([]byte(cw.BOOTSTRAP()))
	if err != nil {
		panic(err)
	}
}

// 与えられた算術コマンドをアセンブリコードに変換して出力する
func (cw *codeWriter) WriteArithmetic(command string) {
	switch command {
	case "add":
		_, err := cw.asmFile.Write([]byte(ADD))
		if err != nil {
			panic(err)
		}
	case "sub":
		_, err := cw.asmFile.Write([]byte(SUB))
		if err != nil {
			panic(err)
		}
	case "neg":
		_, err := cw.asmFile.Write([]byte(NEG))
		if err != nil {
			panic(err)
		}
	case "and":
		_, err := cw.asmFile.Write([]byte(AND))
		if err != nil {
			panic(err)
		}
	case "or":
		_, err := cw.asmFile.Write([]byte(OR))
		if err != nil {
			panic(err)
		}
	case "not":
		_, err := cw.asmFile.Write([]byte(NOT))
		if err != nil {
			panic(err)
		}
	case "gt":
		_, err := cw.asmFile.Write([]byte(cw.GT()))
		if err != nil {
			panic(err)
		}
	case "lt":
		_, err := cw.asmFile.Write([]byte(cw.LT()))
		if err != nil {
			panic(err)
		}
	case "eq":
		_, err := cw.asmFile.Write([]byte(cw.EQ()))
		if err != nil {
			panic(err)
		}
	}
}

func (cw *codeWriter) WritePushPop(command string, segment string, index int) {
	switch command {
	case "push":
		_, err := cw.asmFile.Write([]byte(cw.parsePush(segment, index)))
		if err != nil {
			panic(err)
		}
	case "pop":
		_, err := cw.asmFile.Write([]byte(cw.parsePop(segment, index)))
		if err != nil {
			panic(err)
		}
	}
}

func (cw *codeWriter) WriteLabel(fn, label string) {
	_, err := cw.asmFile.Write([]byte(cw.parseLabel(fn, label)))
	if err != nil {
		panic(err)
	}
}

func (cw *codeWriter) WriteGoTo(fn, destination string) {
	_, err := cw.asmFile.Write([]byte(cw.parseGoTo(fn, destination)))
	if err != nil {
		panic(err)
	}
}

func (cw *codeWriter) WriteIfGoTo(fn, destination string) {
	_, err := cw.asmFile.Write([]byte(cw.parseIfGoTo(fn, destination)))
	if err != nil {
		panic(err)
	}
}

func (cw *codeWriter) WriteFunction(fn string, argc int) {
	_, err := cw.asmFile.Write([]byte(cw.parseFunction(fn, argc)))
	if err != nil {
		panic(err)
	}
}

func (cw *codeWriter) WriteCall(f string, n int) {
	_, err := cw.asmFile.Write([]byte(cw.parseCall(f, n)))
	if err != nil {
		panic(err)
	}
}

func (cw *codeWriter) WriteReturn() {
	_, err := cw.asmFile.Write([]byte(RETURN))
	if err != nil {
		panic(err)
	}
}

func (cw *codeWriter) Close() {
	err := cw.asmFile.Close()
	if err != nil {
		panic(err)
	}
}

func (cw *codeWriter) parsePush(segment string, idx int) string {
	switch segment {
	case "local":
		return "// local\n" +
			"@LCL\n" +
			"D=M\n" +
			"@" + itoa(idx) + "\n" +
			"A=D+A\n" +
			"D=M\n" +
			PUSH
	case "argument":
		return "// argument\n" +
			"@ARG\n" +
			"D=M\n" +
			"@" + itoa(idx) + "\n" +
			"A=D+A\n" +
			"D=M\n" +
			PUSH
	case "this":
		return "// this\n" +
			"@THIS\n" +
			"D=M\n" +
			"@" + itoa(idx) + "\n" +
			"A=D+A\n" +
			"D=M\n" +
			PUSH
	case "that":
		return "// that\n" +
			"@THAT\n" +
			"D=M\n" +
			"@" + itoa(idx) + "\n" +
			"A=D+A\n" +
			"D=M\n" +
			PUSH
	case "pointer":
		if idx == 0 {
			return "// pointer\n" +
				"@THIS\n" +
				"D=M\n" +
				PUSH
		} else {
			return "// pointer\n" +
				"@THAT\n" +
				"D=M\n" +
				PUSH
		}
	case "constant":
		return "// constant\n" +
			"@" + itoa(idx) + "\n" +
			"D=A\n" +
			PUSH
	case "static":
		return "// static\n" +
			"@" + strings.Trim(cw.parser.GetVMFileName(), ".vm") + "." + itoa(idx) + "\n" +
			"D=M\n" +
			PUSH
	case "temp":
		return "// temp\n" +
			"@R5\n" +
			"D=A\n" +
			"@" + itoa(idx) + "\n" +
			"A=D+A\n" +
			"D=M\n" +
			PUSH
	default:
		return ""
	}
}

func (cw *codeWriter) parsePop(segment string, idx int) string {
	switch segment {
	case "local":
		return "// local\n" +
			"@LCL\n" +
			"D=M\n" +
			"@" + itoa(idx) + "\n" +
			"D=D+A\n" +
			POP
	case "argument":
		return "// argument\n" +
			"@ARG\n" +
			"D=M\n" +
			"@" + itoa(idx) + "\n" +
			"D=D+A\n" +
			POP
	case "this":
		return "// this\n" +
			"@THIS\n" +
			"D=M\n" +
			"@" + itoa(idx) + "\n" +
			"D=D+A\n" +
			POP
	case "that":
		return "// that\n" +
			"@THAT\n" +
			"D=M\n" +
			"@" + itoa(idx) + "\n" +
			"D=D+A\n" +
			POP
	case "pointer":
		if idx == 0 {
			return "// pointer\n" +
				"@THIS\n" +
				"D=A\n" +
				POP
		} else {
			return "// pointer\n" +
				"@THAT\n" +
				"D=A\n" +
				POP
		}
	case "static":
		return "// static\n" +
			"@" + strings.Trim(cw.parser.GetVMFileName(), ".vm") + "." + itoa(idx) + "\n" +
			"D=A\n" +
			POP
	case "temp":
		return "// temp\n" +
			"@R5\n" +
			"D=A\n" +
			"@" + itoa(idx) + "\n" +
			"D=D+A\n" +
			POP
	default:
		return ""
	}
}

func (cw *codeWriter) GT() string {
	cw.count++
	return "// gt\n" +
		"@SP\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"A=A-1\n" +
		"D=M-D\n" +
		"@GT.true." + itoa(cw.count) + "\n" +
		"\nD;JGT\n" +
		"@SP\n" +
		"A=M-1\n" +
		"M=0\n" +
		"@GT.after." + itoa(cw.count) + "\n" +
		"0;JMP\n" +
		"(GT.true." + itoa(cw.count) + ")\n" +
		"@SP\n" +
		"A=M-1\n" +
		"M=-1\n" +
		"(GT.after." + itoa(cw.count) + ")\n"
}

func (cw *codeWriter) LT() string {
	cw.count++
	return "// lt\n" +
		"@SP\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"A=A-1\n" +
		"D=M-D\n" +
		"@LT.true." + itoa(cw.count) + "\n" +
		"D;JLT\n" +
		"@SP\n" +
		"A=M-1\n" +
		"M=0\n" +
		"@LT.after." + itoa(cw.count) + "\n" +
		"0;JMP\n" +
		"(LT.true." + itoa(cw.count) + ")\n" +
		"@SP\n" +
		"A=M-1\n" +
		"M=-1\n" +
		"(LT.after." + itoa(cw.count) + ")\n"
}

func (cw *codeWriter) EQ() string {
	cw.count++
	return "// eq\n" +
		"@SP\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"A=A-1\n" +
		"D=M-D\n" +
		"@EQ.true." + itoa(cw.count) + "\n" +
		"D;JEQ\n" +
		"@SP\n" +
		"A=M-1\n" +
		"M=0\n" +
		"@EQ.after." + itoa(cw.count) + "\n" +
		"0;JMP\n" +
		"(EQ.true." + itoa(cw.count) + ")\n" +
		"@SP\n" +
		"A=M-1\n" +
		"M=-1\n" +
		"(EQ.after." + itoa(cw.count) + ")\n"
}

func (cw *codeWriter) parseLabel(fn, label string) string {
	return "(" + fn + "$" + label + ")\n"
}

func (cw *codeWriter) parseGoTo(fn, destination string) string {
	return "@" + fn + "$" + destination + "\n" +
		"0;JMP\n"
}

func (cw *codeWriter) parseIfGoTo(fn, destination string) string {
	return "@SP\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"@" + fn + "$" + destination + "\n" +
		"D;JNE\n"
}

func (cw *codeWriter) parseFunction(fn string, argc int) string {
	ret := "(" + fn + ")\n" +
		"@SP\n" +
		"A=M\n"
	for i := 0; i < argc; i++ {
		ret += "M=0\n" +
			"A=A+1\n"
	}
	return ret +
		"D=A\n" +
		"@SP\n" +
		"M=D\n"
}

func (cw *codeWriter) parseCall(f string, n int) string {
	cw.count++
	return "// function call\n" +
		// SP -> R13
		"@SP\n" +
		"D=M\n" +
		"@R13\n" +
		"M=D\n" +
		// @RET -> *SP
		"@RET." + itoa(cw.count) + "\n" +
		"D=A\n" +
		"@SP\n" +
		"A=M\n" +
		"M=D\n" +
		// SP++
		"@SP\n" +
		"M=M+1\n" +
		// LCL -> *SP
		"@LCL\n" +
		"D=M\n" +
		"@SP\n" +
		"A=M\n" +
		"M=D\n" +
		// SP++
		"@SP\n" +
		"M=M+1\n" +
		// ARG -> *SP
		"@ARG\n" +
		"D=M\n" +
		"@SP\n" +
		"A=M\n" +
		"M=D\n" +
		// SP++
		"@SP\n" +
		"M=M+1\n" +
		// THIS -> *SP
		"@THIS\n" +
		"D=M\n" +
		"@SP\n" +
		"A=M\n" +
		"M=D\n" +
		// SP++
		"@SP\n" +
		"M=M+1\n" +
		// THAT -> *SP
		"@THAT\n" +
		"D=M\n" +
		"@SP\n" +
		"A=M\n" +
		"M=D\n" +
		// SP++
		"@SP\n" +
		"M=M+1\n" +
		// R13 - n -> ARG
		"@R13\n" +
		"D=M\n" +
		"@" + itoa(cw.count) + "\n" +
		"D=D-A\n" +
		"@ARG\n" +
		"M=D\n" +
		// SP -> LCL
		"@SP\n" +
		"D=M\n" +
		"@LCL\n" +
		"M=D\n" +
		"@" + f + "\n" +
		"0;JMP\n" +
		"(RET." + itoa(cw.count) + ")\n"
}

func (cw *codeWriter) BOOTSTRAP() string {
	return "@256\n" +
		"D=A\n" +
		"@SP\n" +
		"M=D\n" +
		"// call Sys.init 0\n" +
		cw.parseCall("Sys.init", 0) +
		"0;JMP\n"
}

func itoa(integer int) string {
	return strconv.Itoa(integer)
}

// func atoi(str string) int {
// 	ret, _ := strconv.Atoi(str)
// 	return ret
// }
