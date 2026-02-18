package translator

// Keywords представляет список ключевых слов C# (категория W)
var Keywords = []string{
	"using",
	"namespace",
	"class",
	"static",
	"void",
	"int",
	"string",
	"float",
	"double",
	"bool",
	"if",
	"else",
	"for",
	"while",
	"return",
	"new",
	"true",
	"false",
}

// Operations представляет список операций C# (категория O)
var Operations = []string{
	// Арифметические операции
	"+", "-", "*", "/", "%",
	// Операции сравнения
	"==", "!=", "<", ">", "<=", ">=",
	// Логические операции
	"&&", "||", "!",
	// Операции присваивания
	"=", "+=", "-=", "*=", "/=", "%=",

}

// Separators представляет список разделителей C# (категория R)
var Separators = []string {
	";", ",", ".", "(", ")", "{", "}", "[", "]",
}

