package tasks

/*
Задание №4: Подсчёт количества слов в строке
Напиши функцию, которая принимает строку и возвращает количество слов в ней.

Уточнения:
- Слова разделяются пробелами, табуляциями и переводами строки.
- Можно считать, что знаки препинания являются частью слова.
- Пустая строка — 0 слов.

Пример:
input := "Hello, world!\nThis is Go."
output := 5
*/

func GetCountWords(input string) int {
	countWord := 0
	words := MultiSplit(input, " ", "\n", "\t")
	for _, word := range words {
		if len(word) > 0 {
			countWord++
		}
	}
	return countWord
}

func MultiSplit(input string, seps ...string) []string {
	var result []string
	lastIndex := 0
	for i := 0; i < len(input); i++ {
		for _, sep := range seps {
			if string(input[i]) == sep {
				result = append(result, input[lastIndex:i])
				lastIndex = i + 1
			}
		}
	}
	result = append(result, input[lastIndex:])
	return result
}

// MultiSplitUpdated input any string, seps - 1 character
func MultiSplitUpdated(input string, seps ...string) []string {
	var result []string

	sepMap := make(map[rune]bool)
	for _, sep := range seps {
		if len(sep) != 1 {
			panic("Separator must be single character '" + sep + "'")
		}
		sepMap[rune(sep[0])] = true
	}

	lastIndex := 0
	for i, r := range input {
		if _, ok := sepMap[r]; ok {
			result = append(result, input[lastIndex:i])
			lastIndex = i + 1

		}
	}

	result = append(result, input[lastIndex:])
	return result
}
