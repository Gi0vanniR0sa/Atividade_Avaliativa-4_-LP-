package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Função que calcula o produto de uma linha de A por uma coluna de B
func calculateElement(A [][]int, B [][]int, result [][]int, row int, col int, wg *sync.WaitGroup) {
	defer wg.Done() // Sinaliza que a goroutine terminou
	sum := 0
	for k := 0; k < len(B); k++ {
		sum += A[row][k] * B[k][col]
	}
	result[row][col] = sum
}

// Função que multiplica as matrizes A e B paralelamente
func multiplyMatricesParallel(A [][]int, B [][]int) [][]int {
	m := len(A)
	//n := len(A[0])
	p := len(B[0])

	// Inicializa a matriz resultante com zeros
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, p)
	}

	var wg sync.WaitGroup

	// Lança goroutines para calcular cada elemento da matriz resultante
	for i := 0; i < m; i++ {
		for j := 0; j < p; j++ {
			wg.Add(1)
			go calculateElement(A, B, result, i, j, &wg)
		}
	}

	// Aguarda todas as goroutines terminarem
	wg.Wait()

	return result
}

func criaMatriz(size int) [][]int {

	matriz := make([][]int, size)
	for i := range matriz {

		matriz[i] = make([]int, size)

		for j := range matriz[i] {

			matriz[i][j] = rand.Intn(10)

		}
	}
	return matriz
}

func main() {

	size := 50
	A := criaMatriz(size)
	B := criaMatriz(size)

	comeco := time.Now()
	result := multiplyMatricesParallel(A, B)
	fim := time.Since(comeco)

	fmt.Println("Resultado da multiplicação de matrizes:")
	for _, row := range result {
		fmt.Println(row)
	}

	fmt.Printf("\nO código gastou %s de tempo para executar.\n", fim)
}
