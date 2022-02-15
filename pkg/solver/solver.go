package solver

import "fmt"

// Инициализация рекурсивного решения
func Solve(sudoku [9][9]uint) ([9][9]uint, bool) {
	var ok bool

	if isSolved(sudoku) {
		return sudoku, true
	}
	i, j := getNextUndefinedPlace(sudoku)
	if i == -1 && j == -1 {
		return sudoku, false
	}
	values := getAlowedValues(uint(i), uint(j), sudoku)
	for _, value := range values {
		sudoku[i][j] = uint(value)
		if sudoku, ok = Solve(sudoku); ok {
			return sudoku, true
		}
	}
	sudoku[i][j] = 0
	return sudoku, false
}

// Получения доступных значений для конкретной клеточки судоку
// Функции вызываются по цепочке т.к. нужно проверить по строчке,
// в которой находится клеточка, потом по столбцу и наконец по квадрату
func getAlowedValues(i, j uint, sudoku [9][9]uint) []int {
	var counter [9]int

	for j := 0; j < 9; j++ {
		if sudoku[i][j] == 0 {
			continue
		}
		counter[sudoku[i][j]-1]++
	}

	return getAlowedValuesByColumns(i, j, sudoku, counter)
}

// Проверка по доступных значений по столбцам
func getAlowedValuesByColumns(i, j uint, sudoku [9][9]uint, counter [9]int) []int {
	for i := 0; i < 9; i++ {
		if sudoku[i][j] == 0 {
			continue
		}
		counter[sudoku[i][j]-1]++
	}
	return getAlowedValuesBySquare(i, j, sudoku, counter)
}

// Проверка по доступных значений по квадратикам
func getAlowedValuesBySquare(i, j uint, sudoku [9][9]uint, counter [9]int) []int {
	result := make([]int, 0, 9)

	squareNum := getSquareNumber(i, j)
	for j := squareNum % 3 * 3; j < squareNum%3*3+3; j++ {
		for k := squareNum / 3 * 3; k < squareNum/3*3+3; k++ {
			if sudoku[j][k] == 0 {
				continue
			}
			counter[sudoku[j][k]-1]++
		}
	}

	for key, value := range counter {
		if value != 0 {
			continue
		}
		result = append(result, key+1)
	}
	return result
}

// Вычисление квадратика, которому принадлежит клеточка
// Каждая клеточка принадлежит одному из 9 квадратиков
func getSquareNumber(i, j uint) int {
	var n uint

	for n = 0; n < 9; n++ {
		if i < n%3*3 || i >= n%3*3+3 {
			continue
		}
		if j < n/3*3 || j >= n/3*3+3 {
			continue
		}
		return int(n)
	}
	return -1
}

// Получение первой найденной клеточки с неопределенным
// значением, в нашем случае это - 0
func getNextUndefinedPlace(sudoku [9][9]uint) (int, int) {
	for i, line := range sudoku {
		for j, value := range line {
			if value == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

// Проверка, что судоку решено по строчкам
// (если встречен неуникальный элемент, то судоку решено неправильно)
func solvedByLines(sudoku [9][9]uint) bool {
	var counter [9]uint

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku[i][j] == 0 {
				return false
			}
			counter[sudoku[i][j]-1]++
		}
		if !eachElementEqual(counter, uint(i+1)) {
			return false
		}
	}
	return true
}

// Проверка, что судоку решено по столбцам
// (если встречен неуникальный элемент, то судоку решено неправильно)
func solvedByColumns(sudoku [9][9]uint) bool {
	var counter [9]uint

	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			counter[sudoku[i][j]-1]++
		}
		if !eachElementEqual(counter, uint(j+1)) {
			return false
		}
	}
	return true
}

// Проверка, что судоку решено по квадратикам
// (если встречен неуникальный элемент, то судоку решено неправильно)
func solvedBySquares(sudoku [9][9]uint) bool {
	var counter [9]uint

	for i := 0; i < 9; i++ {
		for j := i % 3 * 3; j < i%3*3+3; j++ {
			for k := i / 3 * 3; k < i/3*3+3; k++ {
				counter[sudoku[j][k]-1]++
			}
		}
		if !eachElementEqual(counter, uint(i+1)) {
			return false
		}
	}
	return true
}

// Комбинация всех проверок решенности судоку
func isSolved(sudoku [9][9]uint) bool {
	return solvedByLines(sudoku) && solvedByColumns(sudoku) && solvedBySquares(sudoku)
}

// Вспомогательная функции
// Проверка, что все элементы массива имеют данное значение
func eachElementEqual(array [9]uint, value uint) bool {
	for _, el := range array {
		if el != value {
			return false
		}
	}
	return true
}

// Вывод судоку на экран
func PrintMap(sudoku [9][9]uint) {
	for _, line := range sudoku {
		fmt.Println(line)
	}
}
