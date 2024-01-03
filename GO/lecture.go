package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const taille int = 3

func lectureMat(filename string) ([taille][taille]int, int) {

	var mat [taille][taille]int

	file, err := os.Open(filename)

	if err != nil {
		return mat, 1
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()

	for i := 0; i < taille; i++ {

		for j := 0; j < taille; j++ {

			num, err := strconv.Atoi(scanner.Text())

			if err != nil {
				return mat, 1
			}

			mat[i][j] = num
		}
	}
	return mat, 0
}

func main() {

	matA, err := lectureMat("matriceA.txt")
	if err == 1 {
		fmt.Println("Erreur lors de la lecture")
	}
	fmt.Println(matA)
}
