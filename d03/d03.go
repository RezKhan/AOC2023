package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"unicode"
)

type matrixNumber struct {
	number   int
	row      int
	start    int
	end      int
	adjacent bool
}

type matrixGear struct {
	row int
	col int
}

func matrixAdjacency(matrix [][]rune, mnum matrixNumber) matrixNumber {
	rowmin := mnum.row - 1
	colmin := mnum.start - 1
	rowmax := mnum.row + 2
	colmax := mnum.end + 2
	if rowmin < 0 {
		rowmin = 0
	}
	if colmin < 0 {
		colmin = 0
	}
	if rowmax > len(matrix) {
		rowmax = len(matrix)
	}
	if colmax > len(matrix[0]) {
		colmax = len(matrix[0])
	}
	// fmt.Printf("range [%d,%d] to [%d,%d]\n", rowmin, colmin, rowmax, colmax)
	for i := rowmin; i < rowmax; i++ {
		for j := colmin; j < colmax; j++ {
			// fmt.Printf("range [%d,%d] to [%d,%d] - [%d,%d]: %s", rowmin, colmin, rowmax, colmax, i, j, string(matrix[i][j]))
			if !unicode.IsDigit(matrix[i][j]) && string(matrix[i][j]) != "." {
				// fmt.Printf(" | is punct")
				mnum.adjacent = true
			}
			// fmt.Printf("\n")
		}
	}
	return mnum
}

func getMatrixNumber(line []rune, row int, start int, matrix [][]rune) (matrixNumber, int) {
	var mnum matrixNumber
	var tempnums []int
	for i := start; i < len(line); i++ {
		if unicode.IsDigit(line[i]) {
			tempnums = append(tempnums, int(line[i]-'0'))
		} else {
			break
		}
	}
	ll := len(tempnums)
	num := 0
	for j := 0; j < ll; j++ {
		num += int(math.Pow10(ll-j-1)) * tempnums[j]
	}
	mnum.number = num
	mnum.row = row
	mnum.start = start
	mnum.end = start + ll - 1
	mnum = matrixAdjacency(matrix, mnum)

	return mnum, mnum.end + 1
}

func ReadMatrix(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	fScanner := bufio.NewScanner(f)

	var matrix [][]rune
	for fScanner.Scan() {
		line := fScanner.Text()
		matrix = append(matrix, []rune(line))
	}

	var matrixnums []matrixNumber
	var mnum matrixNumber
	var matrixgears []matrixGear
	var mgear matrixGear
	for i, row := range matrix {
		for j := 0; j < len(row); j++ {
			if unicode.IsDigit(row[j]) {
				mnum, j = getMatrixNumber(row, i, j, matrix)
				matrixnums = append(matrixnums, mnum)
			} else if string(row[j]) == "*" {
				mgear.row = i
				mgear.col = j
				matrixgears = append(matrixgears, mgear)
			}
		}
	}
	total := 0
	for i := 0; i < len(matrixnums); i++ {
		// fmt.Println(matrixnums[i])
		if matrixnums[i].adjacent {
			total += matrixnums[i].number
		}
	}
	log.Println("Total of adjacent numbers is: ", total)
	for i := 0; i < len(matrixgears); i++ {
		fmt.Printf("gear found at %d:%d. ", matrixgears[i].row, matrixgears[i].col)
		if i%5 == 0 {
			fmt.Printf("\n")
		}
	}
}

func main() {
	filePath := "./d03input.txt"
	ReadMatrix(filePath)
}
