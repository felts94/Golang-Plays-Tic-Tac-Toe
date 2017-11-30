package main

import(
	"fmt"
	"math/rand"
	"bufio"
	"os"
)

type GameBoard struct{
	Outcome string
	//all past moves in the game for calculating win % when game over
	Moves []Move
	Board [3][3]string
}

type MoveList struct {
	moves []possibleMoves
}

type Move struct{
	Tried bool
	WinProbability int
	Board [3][3]string
}

func main() {

	key := "2"
	//create move list
	//load past win results into move list
	//play games
	//save win results to a file

	file, _ := os.Open("WinRates.txt")
	scanner := bufio.NewScanner(file)

	var Moves []Move


	currentBoard := GameBoard{
		Outcome: "Continue",
		Board: toBoard("000000000"),
	}

	Xwin :=0
	Owin :=0
	draw :=0

	for i := 0; i <50000; i++{
		for  currentBoard.Outcome == "Continue"{
			currentBoard = randomMove(currentBoard,"1")
			currentBoard.Outcome = checkGameOver(currentBoard)
			if currentBoard.Outcome == "Continue"{
				currentBoard = makeMove(currentBoard,key)
				currentBoard.Outcome = checkGameOver(currentBoard)
			}
		}

		toStdOut(currentBoard)

		if currentBoard.Outcome == "1"{
			Xwin++
		}
		if currentBoard.Outcome == "2"{
			Owin++
		}
		if currentBoard.Outcome == "Draw"{
			draw++
		}
		currentBoard = GameBoard{
			Outcome: "Continue",
			Board: toBoard("000000000"),
		}

	}

	fmt.Println("Xwin%",100.0*Xwin/(Xwin+Owin+draw))
	fmt.Println("Owin%",100.0*Owin/(Xwin+Owin+draw))
	fmt.Println("draw%",100.0*draw/(Xwin+Owin+draw))

	//currentBoard = randomMove(currentBoard, "1")
	//
	//currentBoard.Outcome = checkGameOver(currentBoard)
	//
	//toStdOut(currentBoard)
}

func toString(brd [3][3]string) string{
	var b string
	for i :=0; i < 3; i++{
		for j := 0; j < 3; j++{
			b +=brd[i][j]
		}
	}
	return b
}

func toStdOut(brd GameBoard){
	var s string

	for i := 0; i < 3; i++{
		for j:= 0; j < 3; j++ {
			if brd.Board[i][j] == "0"{
				s+=" - "
			} else if brd.Board[i][j] == "1"{
				s+=" X "
			} else{
				s+=" O "
			}
		}
		s+="\n"
	}
	fmt.Println(s)
	fmt.Println(brd.Outcome)
}

func toBoard(brd string) [3][3]string{
	var b [3][3]string

	for i := 0; i < 3; i++{
		for j := 0; j < 3; j++{
			b[i][j] = string(brd[i*3 + j])
		}
	}

	return b
}

func getUserMove(brd GameBoard) GameBoard{
	var b GameBoard

	//display current board to the user

	//get the user input

	//check if valid

	//return the updated board
	return b

}

func randomMove(brd GameBoard, key string) GameBoard{
	var newBrd GameBoard

	newBrd = brd

	brds := toString(brd.Board)
	//find open spaces
	var openSpaces []int

	for i:= 0; i < 9; i++{
		if string(brds[i]) == "0"{
			openSpaces = append(openSpaces,i)
		}
	}

	//pick random open space
	a := openSpaces[rand.Intn(len(openSpaces))]

	newBrd.Board = toBoard(brds[:a] + key + brds[1+a:])
	return newBrd
}

func checkGameOver(brd GameBoard) string{

	//check for 3 1s
	R := checkWinRows(brd.Board,"1")
	C := checkWinCols(brd.Board, "1")
	D := checkDiagonalWins(brd.Board, "1")

	if R || C || D {
		return "1"
	}

	//check for 3 2s
	R = checkWinRows(brd.Board,"2")
	C = checkWinCols(brd.Board, "2")
	D = checkDiagonalWins(brd.Board, "2")

	if R || C || D {
		return "2"
	}

	//check for Draw
	if noRemainingMoves(brd.Board) {
		return "Draw"
	}
	return "Continue"
}
func noRemainingMoves(brd [3][3]string) bool{
	for i := 0; i < 3; i++{
		for j := 0; j < 3; j++{
			if brd[i][j] == "0"{
				return false
			}
		}
	}
	return true
}
func checkWinRows(brd [3][3]string, key string) bool{

	for i := 0; i < 3; i++{
		if brd[i][0] == key && brd[i][1] == key && brd[i][2] == key{
			return true
		}
	}
	return false
}

func checkWinCols(brd [3][3]string, key string) bool {

	for j := 0; j < 3; j++{
		if brd[0][j] == key && brd[1][j] == key && brd[2][j] == key{
			return true
		}
	}
	return false
}

func checkDiagonalWins(brd [3][3]string, key string) bool{

	if brd[0][0] == key && brd[1][1] == key && brd[2][2] == key{
		return true
	}

	if brd[0][2] == key && brd[1][1] == key && brd[2][0]== key{
		return true
	}
	return false
}

type possibleMoves struct{

	winProbablity []float32
	moves []string
}

func makeMove(brd GameBoard, key string) GameBoard{
	var b GameBoard
	//get board state

	//check for the best move or a move that hasn't been tried yet

	//store the board and the chosen move for later analysis

	//return the updated game boad

	return b
}