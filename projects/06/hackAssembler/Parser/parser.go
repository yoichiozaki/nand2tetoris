package Parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const (
	A_COMMAND = iota
	C_COMMAND
	L_COMMAND
)

type Parser struct {
	asm []string // 読み込んだアセンブリを行単位でスライスに格納
	ctype int // 現在見ているコマンドのタイプ
	length int // 読み込んだアセンブリファイルの行数
	processing int // 現在見ているコマンドへのポインタ
}

// アセンブリファイルを読み込んでその行数分のスライスに詰め込む関数
func (p *Parser) ReadFile(filename string) {

	// 読み込む対象のアセンブリファイルを開く
	asmFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer asmFile.Close()

	// パーサーの中にアセンブリファイルを配列としてもつ
	p.asm = make([]string, 0, 100)

	// スキャナーの生成
	scanner := bufio.NewScanner(asmFile)

	// 一行づつasmのスライスに追加していく
	for scanner.Scan() {
		// コメントと空行はすっ飛ばし
		if strings.HasPrefix(scanner.Text(), "//") || scanner.Text() == "" {
			continue
		}
		// 空白文字は削除しておく
		x := strings.TrimSpace(scanner.Text())
		p.asm = append(p.asm, stripComment(x))
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}
	p.length = len(p.asm)
}

func New() *Parser {
	return &Parser{length: 0, processing: 0, ctype: -1}
}

// 入力にまだコマンドが存在すればtrueを返す
func (p *Parser) HasMoreCommands() bool {
	if p.processing + 1 <= p.length {
		return true
	}
	return false
}

// 入力から次のコマンドを読み、それを現在のコマンドにする
func (p *Parser) Advance() {
	if p.HasMoreCommands() {
		p.processing++
	}
}

// 現在見ているコマンドのタイプを返す関数
func (p *Parser) CommandType() int {
	switch p.asm[p.processing][0] {
	case '@':
		p.ctype = A_COMMAND
	case '(':
		p.ctype = L_COMMAND
	default:
		p.ctype = C_COMMAND
	}
	return p.ctype
}

// 現在見ているコマンドのシンボルを文字列で返す
func (p *Parser) Symbol() string {
	switch p.CommandType() {
	case A_COMMAND:
		return strings.TrimPrefix(p.asm[p.processing], "@")
	case L_COMMAND:
		return strings.TrimLeft(strings.TrimRight(p.asm[p.processing], ")"), "(")
	default:
		return ""
	}
}
// C命令のdestニーモニックを返す
func (p *Parser) Dest() string {
	tmp := strings.Split(p.asm[p.processing], "=")
	if len(tmp) == 2 {
		return tmp[0]
	}
	return ""
}

// C命令のcompニーモニックを返す
func (p *Parser) Comp() string {
	tmp := strings.Split(p.asm[p.processing], "=")
	if len(tmp) == 2 {
		tmp1 := tmp[1]
		return strings.Split(tmp1, ";")[0]
	}
	return strings.Split(tmp[0], ";")[0]

}

// C命令のjumpニーモニックを返す
func (p *Parser) Jump() string {
	tmp := strings.Split(p.asm[p.processing], ";")
	if len(tmp) == 2 {
		return tmp[1]
	}
	return ""
}

func (p *Parser) PrintASM() {
	fmt.Print("asm: ")
	fmt.Println(p.asm)
}

func (p *Parser) GetLength() int {
	return p.length
}

func (p *Parser) Reset() {
	p.processing = 0
}

func (p *Parser) PrintProcessingCommand() {
	fmt.Println(p.asm[p.processing])
}

func stripComment(source string) string {
	if cut := strings.IndexAny(source, "//"); cut >= 0 {
		return strings.TrimRightFunc(source[:cut], unicode.IsSpace)
	}
	return source
}