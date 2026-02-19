package main

import (
	"fmt"
	"os"

	"translator/translator"
)

func main() {
	// Чтение тестового файла
	content, err := os.ReadFile("tests/test1_all.cs")
	if err != nil {
		fmt.Printf("Ошибка чтения файла: %v\n", err)
		return
	}

	// Создание сканера и анализ
	scanner := translator.NewScanner(string(content))
	tokens, err := scanner.Scan()
	if err != nil {
		fmt.Printf("Лексическая ошибка: %v\n", err)
		return
	}

	// Вывод результатов
	fmt.Println("Токены:")
	fmt.Println("-------")
	for _, token := range tokens {
		fmt.Printf("%s: %-15s (строка %d, колонка %d)\n",
			token.String(), token.Lexeme, token.Line, token.Col)
	}

	fmt.Println("\n-------")
	fmt.Println("Таблицы:")
	fmt.Println("-------")

	fmt.Println("\nИдентификаторы (I):")
	for lexeme, code := range scanner.ITable() {
		fmt.Printf("  %s -> %s\n", code, lexeme)
	}

	fmt.Println("\nЧисловые константы (N):")
	for lexeme, code := range scanner.NTable() {
		fmt.Printf("  %s -> %s\n", code, lexeme)
	}

	fmt.Println("\nСимвольные/строковые константы (C):")
	for lexeme, code := range scanner.CTable() {
		fmt.Printf("  %s -> %s\n", code, lexeme)
	}
}