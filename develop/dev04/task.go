package main

import (
	"fmt"
	"sort"
	"strings"
)

// UniqLower принимает входной массив строк и возвращает новый массив,
// содержащий уникальные строки, приведенные к нижнему регистру.
func UniqLower(in []string) []string {
	res := make([]string, 0, len(in))
	u := make(map[string]bool)

	for _, i := range in {
		if !u[i] {
			res = append(res, strings.ToLower(i)) // Приводим текущую строку к нижнему регистру и добавляем в результат
			u[i] = true                           // Запоминаем, что строка уже встречалась
		}
	}
	return res
}

// AnagramDict принимает входной массив строк и возвращает map,
// где ключи - отсортированные по алфавиту буквы строки (ключ для анаграмм),
// значения - массив строк, являющихся анаграммами для данного ключа.
func AnagramDict(in []string) map[string][]string {
	if len(in) < 2 {
		return nil // Если входной массив содержит меньше двух элементов, возвращаем nil
	}

	buffer := make(map[string][]string)

	uniqIn := UniqLower(in) // Получаем уникальные строки, приведенные к нижнему регистру
	for _, i := range uniqIn {
		sorted := []rune(i)
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i] < sorted[j] // Сортируем буквы текущей строки по возрастанию
		})

		word := string(sorted)                 // Преобразуем отсортированный массив букв обратно в строку
		buffer[word] = append(buffer[word], i) // Добавляем текущую строку в соответствующий ключ в буфере
	}

	res := make(map[string][]string)
	for _, words := range buffer {
		if len(words) > 1 { // Если в группе анаграмм больше одного слова
			sort.Strings(words)   // Сортируем группу анаграмм по алфавиту
			res[words[0]] = words // Записываем в результат первое слово как ключ и всю группу как значение
		}
	}
	return res
}

func main() {
	input := []string{"тест", "листок", "пятка", "пятак", "тяпка", "листок", "пятка", "слиток", "столик"}

	fmt.Println(input)              // Выводим исходный массив
	fmt.Println(AnagramDict(input)) // Выводим словарь анаграмм
}
