package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Tic Tac Toe")

	// Available moves
	board := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	printBoard(&board)

	fmt.Print("\nYou start! Enter your position: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strPos := scanner.Text()
		humanPos, _ := strconv.Atoi(strPos)

		// Human move (1,2,3...9)
		// Positions on board (0,1,2...8)
		board[humanPos-1] = "x"

		if isEndGame(&board) {
			printWinner(&board)
			break
		}

		computerPos := calculateComputerMove(&board)

		if computerPos > -1 {
			fmt.Printf("Computer move: %d\n", computerPos+1)
			board[computerPos] = "o"
		}

		printBoard(&board)

		if isEndGame(&board) {
			printWinner(&board)
			break
		}

		fmt.Print("\nEnter your next move: ")
	}
}

func calculateComputerMove(board *[]string) int {
	bestMaxResult := -100
	var bestMove int
	for i, val := range *board {
		if isPositionAvailable(val) {
			(*board)[i] = "o"

			// posible results for function minimax: -1, 0, 1
			bestResult := minimax(board, 0, false)

			(*board)[i] = strconv.Itoa(i + 1)

			if bestResult > bestMaxResult {
				bestMaxResult = bestResult
				bestMove = i
			}
		}
	}
	return bestMove
}

func minimax(board *[]string, depth int, isMaximizerPlayer bool) int {
	if isWinState(board, "o") {
		return 1
	}
	if isWinState(board, "x") {
		return -1
	}
	if isEndGame(board) {
		return 0
	}
	if isMaximizerPlayer {
		bestMaxResult := -100
		for i, val := range *board {
			if isPositionAvailable(val) {
				(*board)[i] = "o"

				bestResult := minimax(board, 0, false)

				(*board)[i] = strconv.Itoa(i + 1)

				if bestResult > bestMaxResult {
					bestMaxResult = bestResult
				}
			}
		}
		return bestMaxResult
	}

	bestMinResult := 100
	for i, val := range *board {
		if isPositionAvailable(val) {
			(*board)[i] = "x"

			bestResult := minimax(board, depth+1, true)

			if bestResult < bestMinResult {
				bestMinResult = bestResult
			}

			(*board)[i] = strconv.Itoa(i + 1)
		}
	}
	return bestMinResult
}

func isWinState(board *[]string, player string) bool {
	return ((*board)[0] == player && (*board)[1] == player && (*board)[2] == player) || // horizontal line
		((*board)[3] == player && (*board)[4] == player && (*board)[5] == player) || // horizontal line
		((*board)[6] == player && (*board)[7] == player && (*board)[8] == player) || // horizontal line
		((*board)[0] == player && (*board)[3] == player && (*board)[6] == player) || // vertical line
		((*board)[1] == player && (*board)[4] == player && (*board)[7] == player) || // vertical line
		((*board)[2] == player && (*board)[5] == player && (*board)[8] == player) || // vertical line
		((*board)[0] == player && (*board)[4] == player && (*board)[8] == player) || // diagonal line
		((*board)[2] == player && (*board)[4] == player && (*board)[6] == player) // diagonal line
}

func isEndGame(board *[]string) bool {
	pos := -1
	for i, val := range *board {
		if isPositionAvailable(val) {
			pos = i
			break
		}
	}
	return pos == -1
}

func printBoard(board *[]string) {
	fmt.Println("\nBoard and available positions:")
	fmt.Println((*board)[:3])
	fmt.Println((*board)[3:6])
	fmt.Println((*board)[6:])
}

func printWinner(board *[]string) {
	if isWinState(board, "x") {
		fmt.Println("\nYOU WIN!")
		return
	}

	if isWinState(board, "o") {
		fmt.Println("\nCOMPUTER WINS!")
		return
	}

	fmt.Println("\nNOBODY WINS!")
}

func isPositionAvailable(val string) bool {
	return val != "x" && val != "o"
}
