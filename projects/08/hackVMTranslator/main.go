package main

import (
	"flag"
	"fmt"
	"hackVMTranslator/CodeWriter"
	"hackVMTranslator/Parser"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("ERROR: no input .vm files")
		return
	}
	for _, vmFile := range args {
		p := Parser.New(vmFile)
		cw := CodeWriter.New()
		// TODO: 複数の.vmファイルごとにparserを設定して、出力されるasmファイルは一つにする
		cw.SetParser(p)
		defer cw.Close()
		p.PrintVMCodes()
		// cw.WriteBootstrap()
		for p.HasMoreCommands() {
			p.Advance()
			switch p.CommandType() {
			case Parser.CArithmetic:
				cw.WriteArithmetic(p.Arg1())
			case  Parser.CPush:
				cw.WritePushPop(p.GetCommand(), p.Arg1(), p.Arg2())
			case Parser.CPop:
				cw.WritePushPop(p.GetCommand(), p.Arg1(), p.Arg2())
			case Parser.CLabel:
				cw.WriteLabel(strings.Trim(p.GetVMFileName(), ".vm"), p.Arg1())
			case Parser.CGoto:
				cw.WriteGoTo(strings.Trim(p.GetVMFileName(), ".vm"), p.Arg1())
			case Parser.CIfgoto:
				cw.WriteIfGoTo(strings.Trim(p.GetVMFileName(), ".vm"), p.Arg1())
			case Parser.CFunction:
				cw.WriteFunction(p.Arg1(), p.Arg2())
			case Parser.CCall:
				cw.WriteCall(p.Arg1(), p.Arg2())
			case Parser.CReturn:
				cw.WriteReturn()
			}
		}
	}
}