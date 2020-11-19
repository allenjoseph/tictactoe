package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Tic Tac Toe")

	// available positions
	board := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	printBoard(&board)

	fmt.Print("\nYou start! Enter your position: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strPos := scanner.Text()
		humanPos, _ := strconv.Atoi(strPos)

		// human move
		board[humanPos-1] = "x"

		computerPos := minimax(&board, "o")
		if computerPos > -1 {
			fmt.Printf("Computer move: %d\n", computerPos)
			board[computerPos] = "o"
		}

		printBoard(&board)

		if winState(&board, "x") {
			fmt.Println("\nYOU WIN!")
			break
		}

		if winState(&board, "o") {
			fmt.Println("\nCOMPUTER WINS!")
			break
		}

		fmt.Print("\nEnter your next move: ")
	}
}

func minimax(board *[]string, player string) int {
	pos := -1

	for i, val := range *board {
		if isEmpty(val) {
			// TODO: perform minimax algorithm
			pos = i
			break
		}
	}
	return pos
}

func winState(board *[]string, player string) bool {
	return ((*board)[0] == player && (*board)[1] == player && (*board)[2] == player) ||
		((*board)[3] == player && (*board)[4] == player && (*board)[5] == player) ||
		((*board)[6] == player && (*board)[7] == player && (*board)[8] == player) ||
		((*board)[0] == player && (*board)[3] == player && (*board)[6] == player) ||
		((*board)[1] == player && (*board)[4] == player && (*board)[7] == player) ||
		((*board)[2] == player && (*board)[5] == player && (*board)[8] == player) ||
		((*board)[0] == player && (*board)[4] == player && (*board)[8] == player) ||
		((*board)[2] == player && (*board)[4] == player && (*board)[6] == player)
}

func printBoard(board *[]string) {
	fmt.Println("\nBoard and available positions:")
	fmt.Println((*board)[:3])
	fmt.Println((*board)[3:6])
	fmt.Println((*board)[6:])
}

func isEmpty(val string) bool {
	return val != "x" && val != "o"
}
