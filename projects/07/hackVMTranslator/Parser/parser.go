package Parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type COMMAND_TYPE int
const (
	C_ARITHMETIC COMMAND_TYPE = iota
	C_PUSH
	C_POP
	C_LABEL
	C_GOTO
	C_IFGOTO
	C_FUNCTION
	C_RETURN
	C_CALL
	C_ERROR
)

type Parser struct {
	vmFileName string
	vmCodes []string
	processing int
	length int
}

// 指定された.vmファイルを開いて内容をパーサーに取り込み、あらかたの初期設定を行う
func New(vmFileName string) *Parser {
	// .vmファイルを開く
	vmFile, err := os.Open(vmFileName)
	if err != nil {
		panic(err)
	}
	defer vmFile.Close()

	vmCodes := make([]string, 0, 100)
	scanner := bufio.NewScanner(vmFile)

	// .vmファイルの内容をスライスにぶち込む
	for scanner.Scan() {
		// コメント行と空行をすっ飛ばし
		if strings.HasPrefix(scanner.Text(), "//") || scanner.Text() == "" {
			continue
		}
		// 同一行内の余計な空白とコメントを削除
		vmCode := stripComment(scanner.Text())
		vmCodes = append(vmCodes, vmCode)
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	p := &Parser{
		vmFileName: vmFileName,
		vmCodes: vmCodes,
		processing: -1,
		length: len(vmCodes),
	}
	return p
}

// 次のvmコマンドが存在するか
func (p *Parser) HasMoreCommands() bool {
	if p.processing + 1 < p.length {
		return true
	}
	return false
}

// 次のvmコマンドを読み込む
func (p *Parser) Advance() {
	if p.HasMoreCommands() {
		p.processing++
		return
	}
}

func (p *Parser) CommandType() COMMAND_TYPE {
	separatedCodes := strings.Split(p.vmCodes[p.processing], " ")
	switch len(separatedCodes) {
	case 1:
		switch separatedCodes[0] {
		// スタックにおける算術コマンド
		case "add": return C_ARITHMETIC
		case "sub": return C_ARITHMETIC
		case "neg": return C_ARITHMETIC
		case "eq": return C_ARITHMETIC
		case "gt": return C_ARITHMETIC
		case "lt": return C_ARITHMETIC

		// スタックにおける論理コマンド
		case "and": return C_ARITHMETIC
		case "or": return C_ARITHMETIC
		case "not": return C_ARITHMETIC

		case "return": return C_RETURN
		}
	case 2:
		switch separatedCodes[0] {
		case "goto": return C_GOTO
		case "label": return C_LABEL
		case "if-goto": return C_IFGOTO
		}
	case 3:
		switch separatedCodes[0] {
		case "pop": return C_POP
		case "push": return C_PUSH
		case "function": return C_FUNCTION
		case "call": return C_CALL
		}
	}
	// エラー
	log.Printf("ERROR: wrong VM command length: %d\n", len(separatedCodes))
	return C_ERROR
}

// 現在見ているvmコマンドを返す
func (p *Parser) GetCommand() string {
	separatedCodes := strings.Split(p.vmCodes[p.processing], " ")
	return separatedCodes[0]
}

// 現在見ているvmコマンドの第一引数を返す
func (p *Parser) Arg1() string {
	separatedCodes := strings.Split(p.vmCodes[p.processing], " ")
	switch p.CommandType() {

	// 算術コマンドはそれ自体を返す
	// e.g. "add" -> "add"
	case C_ARITHMETIC: return separatedCodes[0]

	// e.g. "goto LOOP" -> "LOOP"
	default: return separatedCodes[1]
	}
}

// 現在見ているvmコマンドの第二引数を返す
// この関数を呼ぶときはp.CommandType()の結果が
// C_PUSHかC_POP、C_FUNCTION、C_CALLのいずれかであることを確認しておかなければならない
func (p *Parser) Arg2() int {
	separatedCodes := strings.Split(p.vmCodes[p.processing], " ")
	arg2, _ := strconv.Atoi(separatedCodes[2])
	return arg2
}

// このパーサーが読み込んでいるvmのコードを表示する
func (p *Parser) PrintVMCodes() {
	fmt.Println("VM CODES:")
	for _, code := range p.vmCodes {
		fmt.Printf("\t%s\n", code)
	}
	fmt.Println()
}

func (p *Parser) GetVMFileName() string {
	return p.vmFileName
}

// 同一行内の余計な空文字とコメントを削除してくれるヘルパー関数
// "	hello, world		// this is comment"
// -> "hello, world"
func stripComment(source string) string {
	spaceTrimmed := strings.TrimSpace(source)
	if cut := strings.IndexAny(spaceTrimmed, "//"); cut >= 0 {
		return strings.TrimRightFunc(spaceTrimmed[:cut], unicode.IsSpace)
	}
	return spaceTrimmed
}