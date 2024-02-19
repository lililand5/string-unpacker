package main

import (
	"errors"
	"fmt"
	"strings"
)

// разбиваем строку на слайс рун и делаем проверки
func SplitStringIntoSlice(input string) ([]rune, error) {
	var result []rune
	var escapeNext bool

	for _, char := range input {
		// Если текущий символ цифра, то делаем проверки:
		if char >= '0' && char <= '9' {
			if len(result) == 0 {
				return nil, errors.New("первый символ не может быть цифрой")
			}
			if escapeNext {
				return nil, errors.New("цифра не может стоять после слэша")
			}
			// если текущая цифра не первая и если предыдущий символ не слэш,
			// то проверяем что предыдущий символ не цифра
			if len(result) > 0 && !escapeNext {
				prevChar := result[len(result)-1]

				if prevChar >= '0' && prevChar <= '9' {
					return nil, errors.New("строка не может содержать последовательность цифр")
				}
			}
		}

		if char == '\\' {
			if escapeNext {
				return nil, errors.New("нельзя использовать два подряд идущих символа экранирования")
			}
			escapeNext = !escapeNext
			continue
		}

		if escapeNext {
			result = append(result, '\\', char) // Добавляем слэш вместе с символом
			escapeNext = false
		} else {
			result = append(result, char)
		}
	}

	return result, nil
}

// Преобразование слайса rune в строку с учетом повторений символов
func ConvertSlice(slice []rune) (string, error) {
	// fmt.Printf("Slice content: %q\n", slice)
	var stringBuilder strings.Builder

	i := 0
	for i < len(slice) {
		char := slice[i]
		var count int = 1

		// Проверяем, что текущий символ не последний и что он - слэш
		escapeChar := i+1 < len(slice) && char == '\\'

		if escapeChar {
			// Если да, то помечаем след символ, и пропускаем эти два символа
			nextChar := slice[i+1]
			i += 2

			// Проверяем, является ли следующий после экранированной последовательности цифрой
			// если является, то count меняем на новую цифру
			// если нет, то count останется 1
			if i < len(slice) && slice[i] >= '0' && slice[i] <= '9' {
				count = int(slice[i] - '0') // Прямое вычитание для получения числового значения
				i++
			}

			// Добавляем экранированную последовательность в соответствии с её количеством
			stringBuilder.WriteString(strings.Repeat(string(char)+string(nextChar), count))
		} else { // либо если текущий символ не слэш, то просто добавляем этот символ
			i++
			// Проверяем, является ли следующий символ цифрой
			if i < len(slice) && slice[i] >= '0' && slice[i] <= '9' {
				count = int(slice[i] - '0')
				i++
			}

			// Добавляем символ в соответствии с его количеством
			stringBuilder.WriteString(strings.Repeat(string(char), count))
		}
	}

	return stringBuilder.String(), nil
}

// Результирующая функция
func ProcessInputString(input string) (string, error) {
	slice, err := SplitStringIntoSlice(input)
	if err != nil {
		return "", err
	}

	transformedString, err := ConvertSlice(slice)
	if err != nil {
		return "", err
	}

	return transformedString, nil
}

func main() {
	// Тестовые примеры
	testCases := []string{
		`a4bc2d5e`, // => "aaaabccddddde"
		`abcd`,     // => "abcd"
		`3abc`,     // Ошибка
		`45`,       // Ошибка
		`aaa10b`,   // Ошибка
		`aaa0b`,    // => "aab"
		``,         //  => ""
		`d\n5abc`,  // => "d\n\n\n\n\nabc"
	}

	for _, testCase := range testCases {
		result, err := ProcessInputString(testCase)
		if err != nil {
			fmt.Printf("Ошибка обработки строки '%s': %v\n", testCase, err)
		} else {
			fmt.Printf("Вход: '%s' => Выход: '%s'\n", testCase, result)
		}
	}
}
