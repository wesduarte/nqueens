package main

import (
    "fmt"
    "time"
    "math"
    "math/rand"
)

func abs(x int) int {
    if x <0 {
        return -x
    }
    return x
}

func random(max int) int {
    return rand.Intn(max)
}

func initialSolution(n int) []int{

    board := make([]int, n)
    for i:= 0; i<n; i++ {
        board[i] = rand.Intn(n)
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
}

func createMatrix(n int) [][]int{
    matrix := make([][]int, n)
    for i:=0; i<n; i++ {
        matrix[i] = make([]int, n)
    }
    return matrix
}

func calculateEnergy(board []int) float64 {
    board_len := len(board)

    collisions := 0.0
    for i:=0; i < board_len; i++ {
        //fmt.Println(i, " iteration")
        for j:=i+1; j < board_len; j++ {
            diff := j - i
            //if board[j] == abs(diff - board[i]) || board[j] == abs(diff + board[i]) {
            if abs(board[j] - board[i]) == diff ||  abs(board[j] - board[i]) == 0 {
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
    newBoard := make([]int, board_len)
    copy(newBoard, board)
    x := rand.Intn(board_len)
    y := rand.Intn(board_len)

    for x == y {
        x = rand.Intn(board_len)
        y = rand.Intn(board_len)
    }

    fmt.Println("rand", x, y)
    newBoard[x], newBoard[y] = newBoard[y], newBoard[x]

    return newBoard
}

func updateTemperature(t float64) float64{
    new_t := 0.9 * t
    return new_t
}

func main() {
    rand.Seed(time.Now().Unix())
    N := 8
    board := initialSolution(N)
    optimal_board := board

    //var temperature float64 = 100.0/float64(N)
    var temperature float64 = (float64(N)*31250000)
    //var temperature float64 = 10000

    fmt.Println(temperature)
    printBoard(board)

    fmt.Println("energy: ", calculateEnergy(board))

    L := (N/2)-1
    //L := 1
    //L := int(temperature) + 3
    iterations := 0

    for temperature > 0.001 {
        for i:=0; i<L; i++ {
            fmt.Println("Generate New Solution")

            newBoard := generateNewSolution(board)

            delta := calculateEnergy(newBoard) - calculateEnergy(board)
            probability := math.Exp(-delta/temperature)
            randomNumber := rand.Float64()
            //fmt.Println("probability rand ", probability)
            //fmt.Println("float64 rand ", randomNumber)
            if delta < 0 {
                board = newBoard
                //printBoard(board)

                //fmt.Println("collisions: ", calculateEnergy(board))
                if calculateEnergy(newBoard) < calculateEnergy(optimal_board) {
                    copy(optimal_board, newBoard)
                }
            } else if(randomNumber <= probability) {
                board = newBoard
                //printBoard(board)
            }
            iterations++
        }
        temperature = updateTemperature(temperature)
        //L = int(temperature) + 3
    }

    printBoard(optimal_board)
    fmt.Println("energy: ", calculateEnergy(optimal_board))
    fmt.Println("temperature: ", temperature)
    if temperature <= 0.001 {
        fmt.Printf("underflow temperature: %.10f\n", temperature)
    }
    fmt.Println("Total of ", iterations, " iterations")

}
