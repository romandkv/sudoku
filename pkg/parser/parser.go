package parser

import (
	"bufio"
	"log"
	"os"
	"unicode"
)

// Построчное чтение файла и конвертация в двумерный массив интов
func GetSudokuMap(filename string) (*[9][9]uint, error) {
	result := new([9][9]uint)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) != 9 {
			return nil, err
		}
		for j, char := range scanner.Text() {
			if !unicode.IsDigit(char) {
				return nil, err
			}
			result[i][j] = uint(char - '0')
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result, nil
}
