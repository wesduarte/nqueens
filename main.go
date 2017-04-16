package main

import (
    "fmt"
    "math/rand"
)

var board []int

func abs(x int) int {
    if x <0 {
        return -x
    }
    return x
}

func initialSolution(n int) []int{

    var board = make([]int, n)
    for i:= 0; i<n; i++ {
        board[i] = i
    }
    return board
}

func printBoard(board []int) {
    board_len := len(board)
    boardMatrix := createMatrix(board_len)
    for i:=0; i < board_len; i++ {
        boardMatrix[i][board[i]] = 1
        fmt.Println(boardMatrix[i])
    }
    //fmt.Println(boardMatrix)
}

func createMatrix(n int) [][]int{
    matrix := make([][]int, n)
    for i:=0; i<n; i++ {
        matrix[i] = make([]int, n)
    }
    return matrix
}

func calculateTemperature(board []int) int {
    board_len := len(board)

    collisions := 0
    for i:=0; i < board_len; i++ {
        fmt.Println(i, " iteration")
        for j:=i+1; j < board_len; j++ {
            diff := j - i
            //if board[j] == abs(diff - board[i]) || board[j] == abs(diff + board[i]) {
            if abs(board[j] - board[i]) == diff {
                //fmt.Printf(" j=%d - i=%d\n", j, i)
                //fmt.Printf("board[i]=%d , board[j]=%d, diff=%d\n", board[i], board[j], diff)
                collisions++
            } else {
                //fmt.Printf("not match i=%d , j=%d, diff=%d\n", board[i], board[j], diff)
            }
        }
    }
    return collisions
}

func generateNewSolution(board []int) []int {
    board_len := len(board)
    x := rand.Intn(board_len)
    y := rand.Intn(board_len)

    for x == y {
        x = rand.Intn(board_len)
        y = rand.Intn(board_len)
    }


    fmt.Println("rand", x, y)
    board[x], board[y] = board[y], board[x]

    return board
}

func main() {

    board = initialSolution(4)

    printBoard(board)

    fmt.Println("collisions: ", calculateTemperature(board))

    for calculateTemperature(board) > 0 {
        fmt.Println("Generate New Solution")

        board = generateNewSolution(board)

        printBoard(board)

        fmt.Println("collisions: ", calculateTemperature(board))

    }


}
