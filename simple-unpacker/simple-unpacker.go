package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// разбиваем строку на слайс символов
// Возвращает ошибку, если после экранирования стоит цифра,
// Возвращает ошибку, если встречается число,
// Возвращает ошибку, если первый символ это цифра
func SplitStringIntoSlice(input string) ([]string, error) {
	var result []string // слайс в наполняем символами из строки input
	var escapeNext bool // нужно ли добавлять слэш след. символу

	for i := 0; i < len(input); i++ {
		if i == 0 && input[i] >= '0' && input[i] <= '9' {
			return nil, errors.New("первый символ в строке не может быть цифрой")
		}

		if escapeNext {
			if input[i] >= '0' && input[i] <= '9' {
				return nil, errors.New("нельзя использовать цифру после экранирования")
			}
			result = append(result, "\\"+string(input[i]))
			escapeNext = false
			continue
		}

		// Если escapeNext - true, значит в предыдущей итерации(в предыдущем символе)
		// обнаружен символ слэш
		// тогда нельзя чтобы текущий символ тоже был слэш
		if input[i] == '\\' {
			if escapeNext {
				return nil, errors.New("нельзя использовать два подряд идущих символа экранирования")
			}
			// Если escapeNext - false, но следующий символ является слэшем, то тоже ошибка
			if i+1 < len(input) && input[i+1] == '\\' {
				return nil, errors.New("нельзя использовать два подряд идущих символа экранирования")
			}
			// Если текущий символ сэш, и следующий символ не слэш и не цифра,
			// значит обнуляем escapeNext
			escapeNext = true
			continue
		}

		if input[i] >= '0' && input[i] <= '9' {
			if i > 0 && input[i-1] >= '0' && input[i-1] <= '9' {
				return nil, errors.New("строка не может содержать последовательность цифр")
			}
		}

		result = append(result, string(input[i]))
	}

	return result, nil
}

// преобразование слайса в двумерный слайс,
// где каждый подслайс группирует символ и количество его повторений
// считаем количество повторений для каждого символа
func ConvertSlice(slice []string) ([][]string, error) {
	var result [][]string

	for i := 0; i < len(slice); i++ {
		char := slice[i]
		countChar := "1" // По умолчанию каждый символ встречается один раз

		// Если след.элемент - количество повторений текущего символа,
		// то countChar присваивается количество,
		// и пропускаем следующий элемент, так как он уже обработан
		if i+1 < len(slice) && slice[i+1][0] >= '0' && slice[i+1][0] <= '9' {
			countChar = slice[i+1]
			i++
		}

		// Преобразуем строку с количеством в число
		count, err := strconv.Atoi(countChar)
		if err != nil {
			fmt.Println("Ошибка при преобразовании строки в число:", err)
			continue
		}

		if count == 0 {
			continue // Пропускаем символы с количеством 0
		}

		// Повторяем символ count раз
		transformedChar := strings.Repeat(char, count)
		result = append(result, []string{transformedChar, countChar})
	}

	return result, nil
}

// Преобразуем двумерный слайс в строку
func SliceToString(slice [][]string) string {
	var stringBuilder strings.Builder

	for _, item := range slice {
		stringBuilder.WriteString(item[0])
	}

	return stringBuilder.String()
}

// Результирующая функция
func ProcessInputString(input string) (string, error) {
	slice, err := SplitStringIntoSlice(input)
	if err != nil {
		return "", err // Ошибка на этапе разбивки строки на символы
	}

	convertedSlice, err := ConvertSlice(slice)
	if err != nil {
		return "", err // Ошибка на этапе подсчета символов
	}

	transformedString := SliceToString(convertedSlice)

	return transformedString, nil
}

func main() {
	// Тестовые примеры
	testCases := []string{
		`a4bc2d5e`,
		`abcd`,
		`3abc`,
		`45`,
		`aaa10b`,
		`aaa0b`,
		``,
		`d\n5abc`,
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
