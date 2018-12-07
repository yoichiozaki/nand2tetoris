package main

import (
	"fmt"
	"hackAssembler/Code"
	"hackAssembler/Parser"
	"hackAssembler/SymbolTable"
	"os"
	"strconv"
	"strings"
)

var table SymbolTable.SymbolTable

func main() {
	table := SymbolTable.New()
	// TODO: 入力ファイル名を固定してしまっている
	result, err := os.Create("./sample.hack")
	if err != nil {
		panic(err)
	}
	defer result.Close()
	p := Parser.New()
	// TODO: 出力ファイル名を固定してしまっている
	p.ReadFile("./sample.asm")
	binary := make([]string, p.GetLength())
	p.PrintASM()

	// 2passしてテーブルの整備をする。
	for p.HasMoreCommands() {
		switch p.CommandType() {
		case Parser.L_COMMAND:
			if !table.Contains(p.Symbol()) {
				table.AddEntry(p.Symbol(), table.Address)
			}
			p.Advance()
		case Parser.A_COMMAND:
			table.Address++
			p.Advance()
		case Parser.C_COMMAND:
			table.Address++
			p.Advance()
		}

	}

	fmt.Println(table.Table)
	p.Reset()
	// 本チャンの機械語コードの生成部分
	for p.HasMoreCommands() {
		fmt.Println()
		p.PrintProcessingCommand()
		// p.CommandType() // set command type
		if p.CommandType() == Parser.A_COMMAND {
			binary = append(binary, "0"+fmt.Sprintf("%015b", str2int(p.Symbol(), table))+"\n")
			fmt.Println("0" + fmt.Sprintf("%015b", str2int(p.Symbol(), table)))
			p.Advance()
			continue
		} else if p.CommandType() == Parser.L_COMMAND {
			p.Advance()
			continue
		}

		if p.Symbol() != "" {
			fmt.Println("symbol:\t" + fmt.Sprintf("%015b", str2int(p.Symbol(), table)))
		} else {
			fmt.Println("symbol:\t" + p.Symbol())
		}
		fmt.Println("symbol:\t" + p.Symbol())
		fmt.Println("dest:\t" + Code.Dest(p.Dest()))
		fmt.Println("comp:\t" + Code.Comp(p.Comp()))
		fmt.Println("jump:\t" + Code.Jump(p.Jump()))

		binary = append(binary, "111"+Code.Comp(p.Comp())+Code.Dest(p.Dest())+Code.Jump(p.Jump())+"\n")
		fmt.Println("111" + Code.Comp(p.Comp()) + Code.Dest(p.Dest()) + Code.Jump(p.Jump()))
		p.Advance() // get next token
	}
	output := strings.Join(binary, "")
	result.Write(([]byte)(output))
}

func str2int(str string, table *SymbolTable.SymbolTable) int {
	// 数字だったらそのまま返す
	// if ret, err := strconv.Atoi(str); err != nil {
	// 	return ret
	// }
	// 文字列だったらテーブルを引く
	if ret, ok := table.Table[str]; ok {
		return ret
	} else {
		if ret, err := strconv.Atoi(str); err == nil {
			return ret
		}
		table.AddEntry(str, table.VariableAddress)
		table.VariableAddress++
		return table.Table[str]
	}
}
