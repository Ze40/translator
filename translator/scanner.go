package translator

import (
	"fmt"
	"unicode"
)

type Scanner struct {
	//Текст для анализа
	text string

	// Фиксированные таблицы
	wTable map[string]string
	oTable map[string]string
	rTable map[string]string

	// Динамические таблицы
	iTable map[string]string
	nTable map[string]string
	cTable map[string]string
}

func NewScanner(text string) *Scanner {
	scanner := new(Scanner)
	scanner.text = text

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

// ITable возвращает таблицю идентификаторов
func (s *Scanner) ITable() map[string]string {
	return s.iTable
}

// NTable возвращает таблицу числовых констант
func (s *Scanner) NTable() map[string]string {
	return s.nTable
}

// CTable возвращает таблицу символьных/строковых констант
func (s *Scanner) CTable() map[string]string {
	return s.cTable
}

func (s *Scanner) Scan() ([]Token, error) {
	tokens := make([]Token, 0)

	i := 0
	line := 1
	col := 1

	peek := func(offset int) rune {
		j := i + offset
		if j >= 0 && j < len(s.text) {
			return rune(s.text[j])
		}
		return '\000'
	}

	advance := func(n int) {
		for k := 0; k < n; k++ {
			if i >= len(s.text) {
				return
			}
			ch := s.text[i]
			i++
			if ch == '\n' {
				line++
				col = 1
			} else {
				col++
			}
		}
	}

	isIdentStart := func(ch rune) bool {
		return unicode.IsLetter(ch) || ch == '_'
	}

	isIdentPart := func(ch rune) bool {
		return unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_'
	}

	for i < len(s.text) {
		ch := peek(0)

		// пропуск пробелов/табов/переводов строк
		if ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n' {
			advance(1)
			continue
		}

		// комментарии // ... \n
		if ch == '/' && peek(1) == '/' {
			advance(2)
			for i < len(s.text) && peek(0) != '\n' {
				advance(1)
			}
			continue
		}

		// комментарии /* ... */
		if ch == '/' && peek(1) == '*' {
			startLine, startCol := line, col
			advance(2)
			for i < len(s.text) && !(peek(0) == '*' && peek(1) == '/') {
				advance(1)
			}
			if i >= len(s.text) {
				return nil, fmt.Errorf("незакрытый комментарий /* */ в строке %d, колонке %d", startLine, startCol)
			}
			advance(2)
			continue
		}

		// строковые литералы "..."
		if ch == '"' {
			startLine, startCol := line, col
			advance(1)
			buf := make([]rune, 0)
			for i < len(s.text) {
				c := peek(0)
				if c == '\n' {
					return nil, fmt.Errorf("перевод строки внутри строкового литерала в строке %d, колонке %d", startLine, startCol)
				}
				if c == '"' {
					advance(1)
					break
				}
				if c == '\\' {
					advance(1)
					esc := peek(0)
					if esc == '\000' {
						return nil, fmt.Errorf("незакрытый строковый литерал в строке %d, колонке %d", startLine, startCol)
					}
					buf = append(buf, '\\', esc)
					advance(1)
					continue
				}
				buf = append(buf, c)
				advance(1)
			}
			if i >= len(s.text) {
				return nil, fmt.Errorf("незакрытый строковый литерал в строке %d, колонке %d", startLine, startCol)
			}

			lexeme := "\"" + string(buf) + "\""
			code := s.AddCharConst(lexeme)
			tokens = append(tokens, Token{Type: C, Code: code, Lexeme: lexeme, Line: startLine, Col: startCol})
			continue
		}

		// символьные литералы '...'
		if ch == '\'' {
			startLine, startCol := line, col
			advance(1)
			buf := make([]rune, 0)
			for i < len(s.text) {
				c := peek(0)
				if c == '\n' {
					return nil, fmt.Errorf("перевод строки внутри символьного литерала в строке %d, колонке %d", startLine, startCol)
				}
				if c == '\'' {
					advance(1)
					break
				}
				if c == '\\' {
					advance(1)
					esc := peek(0)
					if esc == '\000' {
						return nil, fmt.Errorf("незакрытый символьный литерал в строке %d, колонке %d", startLine, startCol)
					}
					buf = append(buf, '\\', esc)
					advance(1)
					continue
				}
				buf = append(buf, c)
				advance(1)
			}
			if i >= len(s.text) {
				return nil, fmt.Errorf("незакрытый символьный литерал в строке %d, колонке %d", startLine, startCol)
			}

			lexeme := "'" + string(buf) + "'"
			code := s.AddCharConst(lexeme)
			tokens = append(tokens, Token{Type: C, Code: code, Lexeme: lexeme, Line: startLine, Col: startCol})
			continue
		}

		// идентификаторы / ключевые слова / метки
		if isIdentStart(ch) {
			startLine, startCol := line, col
			buf := []rune{ch}
			advance(1)
			for isIdentPart(peek(0)) {
				buf = append(buf, peek(0))
				advance(1)
			}
			lexeme := string(buf)

			// проверка на метку (label:)
			if peek(0) == ':' {
				advance(1)
				lexeme = lexeme + ":"
				code := s.AddIdentifier(lexeme)
				tokens = append(tokens, Token{Type: I, Code: code, Lexeme: lexeme, Line: startLine, Col: startCol})
				continue
			}

			// ищем в таблице ключевых слов
			found := false
			for code, wLexeme := range s.wTable {
				if wLexeme == lexeme {
					var num int
					fmt.Sscanf(code, "W%d", &num)
					tokens = append(tokens, Token{Type: W, Code: num, Lexeme: lexeme, Line: startLine, Col: startCol})
					found = true
					break
				}
			}
			if !found {
				code := s.AddIdentifier(lexeme)
				tokens = append(tokens, Token{Type: I, Code: code, Lexeme: lexeme, Line: startLine, Col: startCol})
			}
			continue
		}

		// числа (int/float/exp)
		if unicode.IsDigit(ch) || (ch == '.' && unicode.IsDigit(peek(1))) {
			startLine, startCol := line, col
			buf := make([]rune, 0)
			hasDot := false
			hasExp := false

			// целая часть
			if ch != '.' {
				for unicode.IsDigit(peek(0)) {
					buf = append(buf, peek(0))
					advance(1)
				}
			}
			// дробная часть
			if peek(0) == '.' && unicode.IsDigit(peek(1)) {
				hasDot = true
				buf = append(buf, '.')
				advance(1)
				for unicode.IsDigit(peek(0)) {
					buf = append(buf, peek(0))
					advance(1)
				}
			}

			// экспонента
			if peek(0) == 'e' || peek(0) == 'E' {
				hasExp = true
				buf = append(buf, peek(0))
				advance(1)
				if peek(0) == '+' || peek(0) == '-' {
					buf = append(buf, peek(0))
					advance(1)
				}
				if !unicode.IsDigit(peek(0)) {
					return nil, fmt.Errorf("некорректная экспонента в числе в строке %d, колонке %d", startLine, startCol)
				}
				for unicode.IsDigit(peek(0)) {
					buf = append(buf, peek(0))
					advance(1)
				}
			}

			_ = hasDot
			_ = hasExp

			lexeme := string(buf)
			code := s.AddNumConst(lexeme)
			tokens = append(tokens, Token{Type: N, Code: code, Lexeme: lexeme, Line: startLine, Col: startCol})
			continue
		}

		// операции и разделители: пробуем самые длинные
		startLine, startCol := line, col

		// 3-символьные
		tri := string(ch) + string(peek(1)) + string(peek(2))
		found := false
		for code, op := range s.oTable {
			if op == tri {
				var num int
				fmt.Sscanf(code, "O%d", &num)
				tokens = append(tokens, Token{Type: O, Code: num, Lexeme: tri, Line: startLine, Col: startCol})
				advance(3)
				found = true
				break
			}
		}
		if found {
			continue
		}

		// 2-символьные
		duo := string(ch) + string(peek(1))
		for code, op := range s.oTable {
			if op == duo {
				var num int
				fmt.Sscanf(code, "O%d", &num)
				tokens = append(tokens, Token{Type: O, Code: num, Lexeme: duo, Line: startLine, Col: startCol})
				advance(2)
				found = true
				break
			}
		}
		if found {
			continue
		}

		// 1-символьные операции
		for code, op := range s.oTable {
			if op == string(ch) {
				var num int
				fmt.Sscanf(code, "O%d", &num)
				tokens = append(tokens, Token{Type: O, Code: num, Lexeme: string(ch), Line: startLine, Col: startCol})
				advance(1)
				found = true
				break
			}
		}
		if found {
			continue
		}

		// разделители
		for code, sep := range s.rTable {
			if sep == string(ch) {
				var num int
				fmt.Sscanf(code, "R%d", &num)
				tokens = append(tokens, Token{Type: R, Code: num, Lexeme: string(ch), Line: startLine, Col: startCol})
				advance(1)
				found = true
				break
			}
		}
		if found {
			continue
		}

		return nil, fmt.Errorf("недопустимый символ: %q в строке %d, колонке %d", ch, startLine, startCol)
	}

	return tokens, nil
}

func Scan(scanner *Scanner) ([]Token, error) {
	return scanner.Scan()
}