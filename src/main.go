package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type coord struct {
	X int
	Y int
}
var game_board = [3][3]string{}
var while_count = 1
var player_choice int
var point string
var selected_coord coord

func player_step() {
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
}
func ai_step() {
	can := 1
	//AI time
	for can != 0 {
		selected_coord.X = rand.Intn(3-0) + 0
		selected_coord.Y = rand.Intn(3-0) + 0
		if can_place_point() == true {
			game_board[selected_coord.X][selected_coord.Y] = point
			can = 0
		}
	}
	show_board()
	win_cheking()
}
func create_board() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			game_board[i][j] = "| |"
		}
	}
}
func point_swicher() {
	if point == "|X|" {
		point = "|O|"
	} else {
		point = "|X|"
	}
}
func show_board() {
	for i := 0; i < 3; i++ {
		fmt.Println()
		for j := 0; j < 3; j++ {
			fmt.Print(game_board[i][j] + " ")
		}
	}
	fmt.Println()
}
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
	for while_count == 1 {
		fmt.Println("Write 1, to play by X, or write 2, to play by O")
		fmt.Fscan(os.Stdin, &player_choice)
		if player_choice != 1 && player_choice != 2 {
			while_count = 1
		} else {
			while_count = 0
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
	create_board()
	show_board()
	if point_select() == 1 {
		point = "|X|"
	} else {
		point = "|O|"
	}
	fmt.Println("U picked " + point)
	win := 0
	for win != 1 {
		player_step()
		show_board()
		if win_cheking() == true {
			win = 1
			fmt.Println()
			fmt.Println(point + " Win!")
		}
		fmt.Println()
		//changing point
		point_swicher()
		ai_step()
		if win_cheking() == true {
			win = 1
			fmt.Println()
			fmt.Println(point + " Win!")
		}
		//Changing point again
		point_swicher()
	}
}
