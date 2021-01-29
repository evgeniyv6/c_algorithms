package main

import (
	"fmt"
	"sync"
)

type hossLegalMoves struct {
	x,y int
}

var (
	hossMovePossibilities = []hossLegalMoves{
		{1, -2},
		{1, 2},
		{2, -1},
		{2, 1},
		{-1, -2},
		{-1, 2},
		{-2, -1},
		{-2, 1},
	}
)

func passBoard(width, height int) [][]int {
	var wg sync.WaitGroup
	res, quit := make(chan [][]int), make(chan int)
	defer close(quit)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			wg.Add(1)
			xStartPos := x
			yStartPos := y
			go func() {
				if ok, boardRes := fillBoard(quit, width, height, xStartPos, yStartPos); ok {
					res <- boardRes
				}
				wg.Done()
			}()
		}
	}

	go func() {
		wg.Wait()
		close(res)
	}()
	return <-res
}

func fillBoard(quit chan int, width, height, col, row int) (bool, [][]int) {
	board := make([][]int, height)
	for i := range board {
		board[i] = make([]int, width) // default by zeros
	}

	if hossJump(quit, width, height, board, 1, col, row) {
		return true, board
	}
	return false, nil
}

func hossJump(quit chan int, width, height int, board [][]int, pos int, col int, row int) bool {
	select {
	case <-quit:
		return false
	default:
	}

	if col < 0 || row < 0 || col >= width || row >= height {
		return false
	}
	if board[row][col] > 0 {
		return false
	}
	board[row][col] = pos
	if pos == width * height {
		return true
	}
	for _, step := range hossMovePossibilities {
		if hossJump(quit, width, height, board, pos+1, col+step.x, row+step.y) {
			return true
		}
	}
	board[row][col] = 0
	return false
}

func main() {
	var N,M = 8,7
	if okBoard:=passBoard(N, M); okBoard != nil {
		for i:=0;i<len(okBoard);i++ {
			fmt.Println(okBoard[i])
		}
	} else {
		fmt.Println("Cannot eval.")
	}
}