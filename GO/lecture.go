package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const taille int = 3

func lectureMat(mat [][]int, filename string) int {

	file, err := os.Open(filename)

	if err != nil {
		return 1
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		values := strings.Fields(scanner.Text())
		var line []int

		for _, value := range values {
			num, err := strconv.Atoi(value)
			if err != nil {
				return 1
			}
			line = append(line, num)
		}

		fmt.Println(line)
		mat = append(mat, line)

	}
	fmt.Println(mat)

	return 0
}

func main() {

	var matA [][]int

	err := lectureMat(matA, "matriceA.txt")
	if err == 1 {
		fmt.Println("Erreur lors de la lecture")
	}
	fmt.Println(matA)
}
