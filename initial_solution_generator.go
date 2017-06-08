package main

import (
	"bufio"
	"time"
    "fmt"
    "os"
    "math/rand"
)

func initialSolution(n int) []int{

    board := make([]int, n)
    for i:= 0; i<n; i++ {
        board[i] = i
    }
    return board
}

func createMatrix(n int) [][]int{
    matrix := make([][]int, n)
    for i:=0; i<n; i++ {
        matrix[i] = make([]int, n)
    }
    return matrix
}

func printBoard(board []int) {
    board_len := len(board)
    boardMatrix := createMatrix(board_len)
    for i:=0; i < board_len; i++ {
        boardMatrix[i][board[i]] = 1
        fmt.Println(boardMatrix[i])
    }
}

func checkError(e error) {
	if(e != nil) {
		panic(e)
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	f, err := os.Create("initial_solution.txt")
	defer f.Close()

	checkError(err)
	w := bufio.NewWriter(f)
    defer w.Flush()

    fmt.Println("Informe o tamanho do tabuleiro\n")
    var N int
    fmt.Scan(&N)

    initialSolution := initialSolution(N)

    for i:=0; i < N; i++ {
    	fmt.Fprint(w, initialSolution[i], " ")
    }
    fmt.Println("Solucao inicial gerada!")
    printBoard(initialSolution)
}