package main

import (
	"fmt"
	"sync"
	"time"
)

// Realiza a multiplicação de duas matrizes e retorna a resultante
func multiplicaMatriz(A, B [][]int) [][]int {

	n := len(A)    //Armazena o número de linhas de A
	m := len(A[0]) //Armazena o número de colunas de A
	p := len(B[0]) //Armazena o número de colunas de B

	//Define a dimensão da matriz como C[n][p] para uma matriz A[n][m] e B[n][p]
	C := make([][]int, n)

	for i := range C {

		C[i] = make([]int, p) //Cria p colunas em vista de i linhas

	}

	var wg sync.WaitGroup //Variável de sincronização

	for i := 0; i < n; i++ { //i = linha

		for j := 0; j < p; j++ { //j = coluna

			wg.Add(1) //Fornece o número de goroutines

			go func(i, j int) {

				defer wg.Done() //Indica o fim de uma goroutine

				for k := 0; k < m; k++ {

					fmt.Printf("Calculando C[%d][%d]: A[%d][%d] * B[%d][%d] = %d * %d\n", i, j, i, k, k, j, A[i][k], B[k][j])
					C[i][j] += A[i][k] * B[k][j]

				}
				fmt.Printf("C[%d][%d] = %d\n", i, j, C[i][j])
			}(i, j)

		}

	}

	wg.Wait()
	return C

}

func main() {

	A := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	B := [][]int{
		{26, 27, 28, 29, 30},
		{31, 32, 33, 34, 35},
		{36, 37, 38, 39, 40},
		{41, 42, 43, 44, 45},
		{46, 47, 48, 49, 50},
	}

	comeco := time.Now()
	C := multiplicaMatriz(A, B)
	fim := time.Since(comeco)

	fmt.Printf("\nMatriz resultante C:\n\n")

	for _, linha := range C {

		fmt.Println(linha) //Itera cada linha de C[n][p]

	}
	fmt.Printf("\nO código gastou %s de tempo para executar.\n", fim)

}
