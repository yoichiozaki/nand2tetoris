package main

import (
	"flag"
	"fmt"
	"hackVMTranslator/CodeWriter"
	"hackVMTranslator/Parser"
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
		cw.SetParser(p)
		defer cw.Close()
		p.PrintVMCodes()
		for p.HasMoreCommands() {
			p.Advance()
			switch p.CommandType() {
			case Parser.C_ARITHMETIC:
				cw.WriteArithmetic(p.Arg1())
			case  Parser.C_PUSH:
				cw.WritePushPop(p.GetCommand(), p.Arg1(), p.Arg2())
			case Parser.C_POP:
				cw.WritePushPop(p.GetCommand(), p.Arg1(), p.Arg2())
			}
		}
	}
}