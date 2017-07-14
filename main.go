package main

import (
    "fmt"
    "time"
    "math"
    "math/rand"
    "io"
    "os"
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
        board[i] = i
    }
    return board
}

func readInitialSolutionFromFile() []int{
    file, err := os.Open("initial_solution.txt")

        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

        var perline int
        var board []int

        for {
            _, err := fmt.Fscanf(file, "%d", &perline)

            if err != nil {

                    if err == io.EOF {
                        break
                    }
                    fmt.Println(err)
                    os.Exit(1)
            }

            board = append(board, perline)
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
        for j:=i+1; j < board_len; j++ {
            diff_row := j - i
            diff_col := abs(board[j] - board[i])
            if diff_col == diff_row ||  diff_col == 0 {
                collisions++
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
        y = rand.Intn(board_len)
    }

    newBoard[x], newBoard[y] = newBoard[y], newBoard[x]

    return newBoard
}

func updateTemperature(t float64) float64{
    new_t := 0.97 * t
    return new_t
}

func simulatedAnnealing(board []int) []int {
	N := len(board)
	optimal_board := board
	optimal_iter := 0


	var temperature float64 = float64(N)*6
	fmt.Println("Initial Board")
	printBoard(board)
	fmt.Println("Initial Energy: ", calculateEnergy(board))
	fmt.Println("Initial Temperature: ", temperature)
	fmt.Println()

	L := (N/2)-1
	iterations := 0

	start := time.Now()
	for temperature > 0.01 {
	    for i:=0; i<L; i++ {
	        newBoard := generateNewSolution(board)
	        delta := calculateEnergy(newBoard) - calculateEnergy(board)
	        probability := math.Exp(-delta/temperature)
	        randomNumber := rand.Float64()
	        if delta < 0 {
	            board = newBoard
	            if calculateEnergy(newBoard) < calculateEnergy(optimal_board) {
	                copy(optimal_board, newBoard)
	                optimal_iter = iterations
	            }
	        } else if(randomNumber < probability) {
	            board = newBoard
	        }
	        iterations++
	    }
	    temperature = updateTemperature(temperature)
	}

	fmt.Println("Optimal Board")
	printBoard(optimal_board)
	fmt.Println("Final Energy: ", calculateEnergy(optimal_board))
	fmt.Println("Final Temperature: ", temperature)
	fmt.Println("Found at ", optimal_iter, "th iteration")
	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s\n", elapsed)

	return optimal_board
}

func main() {
    rand.Seed(time.Now().Unix())
    board := readInitialSolutionFromFile()
    simulatedAnnealing(board)

}
