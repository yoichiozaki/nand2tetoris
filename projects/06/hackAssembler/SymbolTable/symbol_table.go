package SymbolTable

type SymbolTable struct {
	VariableAddress int
	Address int
	Table   map[string]int
}

func New() *SymbolTable {
	table := SymbolTable{VariableAddress: 16, Address: 0, Table: map[string]int{}}
	// 定義済みシンボルの登録
	table.Table["SP"] = 0
	table.Table["LCL"] = 1
	table.Table["ARG"] = 2
	table.Table["THIS"] = 3
	table.Table["THAT"] = 4
	table.Table["R0"] = 0
	table.Table["R1"] = 1
	table.Table["R2"] = 2
	table.Table["R3"] = 3
	table.Table["R4"] = 4
	table.Table["R5"] = 5
	table.Table["R6"] = 6
	table.Table["R7"] = 7
	table.Table["R8"] = 8
	table.Table["R9"] = 9
	table.Table["R10"] = 10
	table.Table["R11"] = 11
	table.Table["R12"] = 12
	table.Table["R13"] = 13
	table.Table["R14"] = 14
	table.Table["R15"] = 15
	table.Table["SCREEN"] = 16384
	table.Table["KBD"] = 24576
	return &table
}

func (s *SymbolTable) AddEntry(symbol string, address int) {
	s.Table[symbol] = address
}

func (s *SymbolTable) Contains(symbol string) bool {
	_, ok := s.Table[symbol]
	return ok
}

func (s *SymbolTable) GetAddress(symbol string) int {
	ret, _ := s.Table[symbol]
	return ret
}
