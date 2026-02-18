package translator

import "fmt"

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

	// Initialize dynamic tables
	scanner.iTable = make(map[string]string)
	scanner.nTable = make(map[string]string)
	scanner.cTable = make(map[string]string)

	return scanner
}

func (s *Scanner) AddIdentifier(lexeme string) {
	if _, exists := s.iTable[lexeme]; exists {
		return
	}
	var nextCode int = len(s.iTable) + 1
	s.iTable[fmt.Sprintf("I%d", nextCode)] = lexeme
}

func (s *Scanner) AddNumConst(lexeme string) {
	if _, exists := s.nTable[lexeme]; exists {
		return
	}
	var nextCode int = len(s.nTable) + 1
	s.nTable[fmt.Sprintf("N%d", nextCode)] = lexeme
}

func (s *Scanner) AddCharConst(lexeme string) {
	if _, exists := s.cTable[lexeme]; exists {
		return
	}
	var nextCode int = len(s.cTable) + 1
	s.cTable[fmt.Sprintf("C%d", nextCode)] = lexeme
}