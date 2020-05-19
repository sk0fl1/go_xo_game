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

var board_size int
var game_board [7][7]string
var while_count = true
var player_choice int
var point string
var selected_coord coord
var game_mode int

func board_size_choice() {
	fmt.Println("Write size of board (supported 3, 4, 6): ")
	fmt.Fscan(os.Stdin, &board_size)
}
func mode_choice() int {
	k := false
	for k != true {
		fmt.Println("Choice game mode(1 - solo with AI, 2 - 1 vs 1 with ur friend):")
		fmt.Fscan(os.Stdin, &game_mode)
		if game_mode == 1 || game_mode == 2 {
			k = true
		}
	}
	return game_mode
}
func base_win_check(k int, h int, p int) bool{
	counter:=0
	answer:=false
	for i := 0; i < 3; i++ {
		if game_board[i+k+h][i+k+h] == point {
			counter++
		}
	}
	if counter == 3 {
		answer = true
	}
	counter = 0
	for i := 0; i < 3; i++ {
		if game_board[p-i-1+k][i+h] == point {
			counter++
		}
	}
	if counter == 3 {
		answer = true
	}
	counter = 0
	for i := 0; i < 3; i++ {
		if game_board[i+k+h][0+k+h] == point {
			counter++
		}
	}
	if counter == 3 {
		answer = true
	}
	counter = 0
	for i := 0; i < 3; i++ {
		if game_board[i+k+h][1+k+h] == point {
			counter++
		}
	}
	if counter == 3 {
		answer = true
	}
	counter = 0
	for i := 0; i < 3; i++ {
		if game_board[i+k+h][2+k+h] == point {
			counter++
		}
	}
	if counter == 3 {
		answer = true
	}
	counter = 0
	for i := 0; i < 3; i++ {
		if game_board[0+k+h][i+k+h] == point {
			counter++
		}
	}
	if counter == 3 {
		answer = true
	}
	counter = 0
	for i := 0; i < 3; i++ {
		if game_board[1+k+h][i+k+h] == point {
			counter++
		}
	}
	if counter == 3 {
		answer = true
	}
	counter = 0
	for i := 0; i < 3; i++ {
		if game_board[2+k+h][i+k+h] == point {
			counter++
		}
	}
	if counter == 3 {
		answer = true
	}
	counter = 0
	return answer
}
func win_check() bool {
	answer:=false
	if board_size == 3 {
		if base_win_check(0, 0,3) == true {
			answer = true
		}
	}
	if board_size == 4 {
		for k := 0; k < 2; k++ {
			for h := 0; h < 2; h++ {
				if base_win_check(k,h,3) == true {
					answer = true
				}
			}
		}
			}
	if board_size == 6 {
		for k := 0; k < 2; k++ {
			for h := 0; h < 2; h++ {
				if base_win_check(k,h,4) == true {
					answer = true
				}
			}
		}
	}
	return answer
}
func player_step() {
	can := false
	//Players time
	for can != true {
		fmt.Println()
		coord_picking()
		if can_place_point() == true {
			game_board[selected_coord.X][selected_coord.Y] = point
			fmt.Println(point)
			can = true
		}
	}
}
func ai_step() {
	can := false
	//AI time
	for can != true {
		selected_coord.X = rand.Intn(board_size-0) + 0
		selected_coord.Y = rand.Intn(board_size-0) + 0
		if can_place_point() == true {
			game_board[selected_coord.X][selected_coord.Y] = point
			can = true
		}
	}
	show_board()
	win_check()
}
func create_board() {
	for i := 0; i < board_size; i++ {
		for j := 0; j < board_size; j++ {
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
	for i := 0; i < board_size; i++ {
		fmt.Println()
		for j := 0; j < board_size; j++ {
			fmt.Print(game_board[i][j] + " ")
		}
	}
	fmt.Println()
}
func coord_picking() coord {
	z := board_size
	for z == board_size {
		fmt.Println("Write coordinates of point: ")
		fmt.Fscan(os.Stdin, &selected_coord.X)
		fmt.Fscan(os.Stdin, &selected_coord.Y)
		if selected_coord.X < board_size && selected_coord.Y < board_size {
			z = 1
		}
	}
	return selected_coord

}
func point_select() int {
	for while_count == true {
		fmt.Println("Write 1, to play by X, or write 2, to play by O")
		fmt.Fscan(os.Stdin, &player_choice)
		if player_choice != 1 && player_choice != 2 {
			while_count = true
		} else {
			while_count = false
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
	mode_choice()
	board_size_choice()
	create_board()
	show_board()
	if point_select() == 1 {
		point = "|X|"
	} else {
		point = "|O|"
	}
	fmt.Println("U picked " + point)
	win := false
	for win != true {
		player_step()
		show_board()
		if win_check() == true {
			win = true
			fmt.Println()
			fmt.Println(point + " Win!")
			os.Exit(0)
		}
		fmt.Println()
		//changing point
		point_swicher()
		if game_mode == 1 {
			ai_step()
		} else {
			player_step()
			show_board()
		}
		if win_check() == true {
			win = true
			fmt.Println()
			fmt.Println(point + " Win!")
			os.Exit(0)
		}
		//Changing point again
		point_swicher()
	}
}
