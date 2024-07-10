package main

import (
	"fmt"
	"math/rand"
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

	//--var wg sync.WaitGroup //Variável de sincronização

	for i := 0; i < n; i++ { //i = linha

		for j := 0; j < p; j++ { //j = coluna

			//--wg.Add(1) //Fornece o número de goroutines

			//--go func(i, j int) {

			//--defer wg.Done() //Indica o fim de uma goroutine

			for k := 0; k < m; k++ {

				fmt.Printf("Calculando C[%d][%d]: A[%d][%d] * B[%d][%d] = %d * %d\n", i, j, i, k, k, j, A[i][k], B[k][j])
				C[i][j] += A[i][k] * B[k][j]

			}
			fmt.Printf("C[%d][%d] = %d\n", i, j, C[i][j])
			//--}(i, j)

		}

	}

	//--wg.Wait()
	return C

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
	C := multiplicaMatriz(A, B)
	fim := time.Since(comeco)

	fmt.Printf("\nMatriz resultante C:\n\n")

	for _, linha := range C {

		fmt.Println(linha) //Itera cada linha de C[n][p]

	}
	fmt.Printf("\nO código gastou %s de tempo para executar.\n", fim)

}
