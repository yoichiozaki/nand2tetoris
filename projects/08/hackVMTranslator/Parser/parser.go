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

type CommandType int
const (
	CArithmetic CommandType = iota
	CPush
	CPop
	CLabel
	CGoto
	CIfgoto
	CFunction
	CReturn
	CCall
	CError
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

func (p *Parser) CommandType() CommandType {
	separatedCodes := strings.Split(p.vmCodes[p.processing], " ")
	switch len(separatedCodes) {
	case 1:
		switch separatedCodes[0] {
		// スタックにおける算術コマンド
		case "add": return CArithmetic
		case "sub": return CArithmetic
		case "neg": return CArithmetic
		case "eq": return CArithmetic
		case "gt": return CArithmetic
		case "lt": return CArithmetic

		// スタックにおける論理コマンド
		case "and": return CArithmetic
		case "or": return CArithmetic
		case "not": return CArithmetic

		case "return": return CReturn
		}
	case 2:
		switch separatedCodes[0] {
		case "goto": return CGoto
		case "label": return CLabel
		case "if-goto": return CIfgoto
		}
	case 3:
		switch separatedCodes[0] {
		case "pop": return CPop
		case "push": return CPush
		case "function": return CFunction
		case "call": return CCall
		}
	}
	// エラー
	log.Printf("ERROR: wrong VM command length: %d\n", len(separatedCodes))
	return CError
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
	case CArithmetic: return separatedCodes[0]

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