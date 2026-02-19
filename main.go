package main

import (
	"fmt"
	"os"
	"strings"

	"translator/translator"
)

// WriteTokensToFile записывает токены и таблицы в файл
func WriteTokensToFile(tokens []translator.Token, scanner *translator.Scanner, filename string) error {
	var sb strings.Builder

	// Запись токенов
	sb.WriteString("Токены:\n")
	sb.WriteString("-------\n")
	for _, token := range tokens {
		sb.WriteString(fmt.Sprintf("%s: %-15s (строка %d, колонка %d)\n",
			token.String(), token.Lexeme, token.Line, token.Col))
	}

	sb.WriteString("\n-------\n")
	sb.WriteString("Таблицы:\n")
	sb.WriteString("-------\n")

	sb.WriteString("\nИдентификаторы (I):\n")
	for lexeme, code := range scanner.ITable() {
		sb.WriteString(fmt.Sprintf("  %s -> %s\n", code, lexeme))
	}

	sb.WriteString("\nЧисловые константы (N):\n")
	for lexeme, code := range scanner.NTable() {
		sb.WriteString(fmt.Sprintf("  %s -> %s\n", code, lexeme))
	}

	sb.WriteString("\nСимвольные/строковые константы (C):\n")
	for lexeme, code := range scanner.CTable() {
		sb.WriteString(fmt.Sprintf("  %s -> %s\n", code, lexeme))
	}

	return os.WriteFile(filename, []byte(sb.String()), 0644)
}

func main() {
	content, err := os.ReadFile("tests/test1_all.cs")
	if err != nil {
		fmt.Printf("Ошибка чтения файла: %v\n", err)
		return
	}

	scanner := translator.NewScanner(string(content))
	tokens, err := scanner.Scan()
	if err != nil {
		fmt.Printf("Лексическая ошибка: %v\n", err)
		return
	}

	err = WriteTokensToFile(tokens, scanner, "tokens.txt")
	if err != nil {
		fmt.Printf("Ошибка записи в файл: %v\n", err)
		return
	}
}