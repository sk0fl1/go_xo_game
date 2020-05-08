package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var game_board = [3][3]string{}
var q = 1
var player_choice int
var point string

type coord struct {
	X int
	Y int
}

var selected_coord coord

func coord_picking() coord {
	z := 3
	for z == 3 {
		fmt.Println("Write coordinates of point: ")
		fmt.Fscan(os.Stdin, &selected_coord.X)
		fmt.Fscan(os.Stdin, &selected_coord.Y)
		if selected_coord.X < 3 && selected_coord.Y < 3 {
			z = 1
		}
	}
	return selected_coord

}
func win_cheking() bool {
	answer := false
	if game_board[0][0] == point && game_board[0][1] == point && game_board[0][2] == point {
		answer = true
	}
	if game_board[1][0] == point && game_board[1][1] == point && game_board[1][2] == point {
		answer = true
	}
	if game_board[2][0] == point && game_board[2][1] == point && game_board[2][2] == point {
		answer = true
	}
	if game_board[0][0] == point && game_board[1][1] == point && game_board[2][2] == point {
		answer = true
	}
	if game_board[0][2] == point && game_board[1][1] == point && game_board[0][0] == point {
		answer = true
	}
	if game_board[0][0] == point && game_board[1][0] == point && game_board[2][0] == point {
		answer = true
	}
	if game_board[0][1] == point && game_board[1][1] == point && game_board[2][1] == point {
		answer = true
	}
	if game_board[0][2] == point && game_board[1][2] == point && game_board[2][2] == point {
		answer = true
	}
	return answer
}
func point_select() int {
	for q == 1 {
		fmt.Println("Write 1, to play by X, or write 2, to play by O")
		fmt.Fscan(os.Stdin, &player_choice)
		if player_choice != 1 && player_choice != 2 {
			q = 1
		} else {
			q = 0
		}
	}
	return player_choice
}
func can_place_point() bool {
	if game_board[selected_coord.X][selected_coord.Y] == "| |" {
		return true
	} else {
		return false
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	//Creating game board
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			game_board[i][j] = "| |"
		}
	}
	//Showing game board 1st time
	for i := 0; i < 3; i++ {
		fmt.Println()
		for j := 0; j < 3; j++ {
			fmt.Print(game_board[i][j] + " ")
		}
	}
	fmt.Println()
	point_select()
	if point_select() == 1 {
		point = "|X|"
	} else {
		point = "|O|"
	}
	fmt.Println("U picked " + point)
	win := 0
	for win != 1 {
		can := 1
		//Players time
		for can != 0 {
			fmt.Println()
			coord_picking()
			if can_place_point() == true {
				game_board[selected_coord.X][selected_coord.Y] = point
				fmt.Println(point)
				can = 0
			}
		}
		for i := 0; i < 3; i++ {
			fmt.Println()
			for j := 0; j < 3; j++ {
				fmt.Print(game_board[i][j] + " ")
			}
		}
		win_cheking()
		if win_cheking() == true {
			win = 1
			fmt.Println()
			fmt.Println(point + " Win!")
		}
		fmt.Println()
		//changing point
		if point == "|X|" {
			point = "|O|"
		} else {
			point = "|X|"
		}
		can = 1
		//AI time
		for can != 0 {
			selected_coord.X = rand.Intn(3-0) + 0
			selected_coord.Y = rand.Intn(3-0) + 0
			if can_place_point() == true {
				game_board[selected_coord.X][selected_coord.Y] = point
				can = 0
			}
		}
		for i := 0; i < 3; i++ {
			fmt.Println()
			for j := 0; j < 3; j++ {
				fmt.Print(game_board[i][j] + " ")

			}
		}
		win_cheking()
		if win_cheking() == true {
			win = 1
			fmt.Println()
			fmt.Println(point + " Win!")
		}
		//Changing point again
		if point == "|X|" {
			point = "|O|"
		} else {
			point = "|X|"
		}
	}
}
