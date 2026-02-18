package translator

import (
	"fmt"
)

type Scanner struct {
	// Фиксированные таблицы
	wTable map[string]string
	oTable map[string]string
	rTable map[string]string

	// Динамические таблицы
	iTable map[string]string
	nTable map[string]string
	cTable map[string]string
}

func NewScanner() *Scanner {
	scanner := new(Scanner)

	wTableInit := make(map[string]string)
	for i, lexeme := range Keywords {
		wTableInit[fmt.Sprintf("W%d", i+1)] = lexeme
	}

	oTableInit := make(map[string]string)
	for i, lexeme := range Operations {
		oTableInit[fmt.Sprintf("O%d", i+1)] = lexeme

	}

	rTableInit := make(map[string]string)
	for i, lexeme := range Separators {
		rTableInit[fmt.Sprintf("R%d", i+1)] = lexeme

	}

	scanner.wTable = wTableInit
	scanner.oTable = oTableInit
	scanner.rTable = rTableInit

	scanner.iTable = make(map[string]string)
	scanner.nTable = make(map[string]string)
	scanner.cTable = make(map[string]string)

	return scanner
}

func (s *Scanner) AddIdentifier(lexeme string) int {
	if code, exists := s.iTable[lexeme]; exists {
		var num int
		fmt.Sscanf(code, "I%d", &num)
		return num
	}

	var nextCode int = len(s.iTable) + 1
	s.iTable[lexeme] = fmt.Sprintf("I%d", nextCode)
	return nextCode
}

func (s *Scanner) AddNumConst(lexeme string) int {
	if code, exists := s.nTable[lexeme]; exists {
		var num int
		fmt.Sscanf(code, "N%d", &num)
		return num
	}

	var nextCode int = len(s.nTable) + 1
	s.nTable[lexeme] = fmt.Sprintf("N%d", nextCode)
	return nextCode
}

func (s *Scanner) AddCharConst(lexeme string) int {
	if code, exists := s.cTable[lexeme]; exists {
		var num int
		fmt.Sscanf(code, "C%d", &num)
		return num
	}

	var nextCode int = len(s.cTable) + 1
	s.cTable[lexeme] = fmt.Sprintf("C%d", nextCode)
	return nextCode
}