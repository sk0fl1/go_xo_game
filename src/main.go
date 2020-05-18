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
var game_board [10][10]string
var while_count = 1
var player_choice int
var point string
var selected_coord coord
var game_mode int

func board_size_choice() {
	fmt.Println("Write size of board (supported 3, 4, 6): ")
	fmt.Fscan(os.Stdin, &board_size)
}
func mode_choice () int{
	fmt.Println("Choice game mode(1 - solo with AI, 2 - 1 vs 1 with ur friend):")
	fmt.Fscan(os.Stdin, &game_mode)
	return game_mode
}
func win_check() bool {
	answer := false
	counter := 0
	if board_size == 3 {
		for i := 0; i < 3; i++ {
			if game_board[i][i] == point {
				counter++
			}
		}
		if counter == 3 {
			answer = true
		}
		counter = 0
		for i := 0; i < 3; i++ {
			if game_board[3-i-1][i] == point {
				counter++
			}
		}
		if counter == 3 {
			answer = true
		}
		counter = 0
		for i := 0; i < 3; i++ {
			if game_board[i][0] == point {
				counter++
			}
		}
		if counter == 3 {
			answer = true
		}
		counter = 0
		for i := 0; i < 3; i++ {
			if game_board[i][1] == point {
				counter++
			}
		}
		if counter == 3 {
			answer = true
		}
		counter = 0
		for i := 0; i < 3; i++ {
			if game_board[i][2] == point {
				counter++
			}
		}
		if counter == 3 {
			answer = true
		}
		counter = 0
		for i := 0; i < 3; i++ {
			if game_board[0][i] == point {
				counter++
			}
		}
		if counter == 3 {
			answer = true
		}
		counter = 0
		for i := 0; i < 3; i++ {
			if game_board[1][i] == point {
				counter++
			}
		}
		if counter == 3 {
			answer = true
		}
		counter = 0
		for i := 0; i < 3; i++ {
			if game_board[2][i] == point {
				counter++
			}
		}
		if counter == 3 {
			answer = true
		}
		counter = 0
	}
	if board_size == 4 {
		for k := 0; k < 2; k++ {
			for h := 0; h < 2; h++ {
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
					if game_board[3-i-1+k][i+h] == point {
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
			}
		}
	}
	if board_size == 6 {
		for k := 0; k < 3; k++ {
			for h := 0; h < 3; h++ {
				for i := 0; i < 4; i++ {
					if game_board[i+k+h][i+k+h] == point {
						counter++
					}
				}
				if counter == 4 {
					answer = true
				}
				counter = 0
				for i := 0; i < 4; i++ {
					if game_board[4-i-1+k][i+h] == point {
						counter++
					}
				}
				if counter == 4 {
					answer = true
				}
				counter = 0
				for i := 0; i < 4; i++ {
					if game_board[i+k+h][0+k+h] == point {
						counter++
					}
				}
				if counter == 4 {
					answer = true
				}
				counter = 0
				for i := 0; i < 4; i++ {
					if game_board[i+k+h][1+k+h] == point {
						counter++
					}
				}
				if counter == 4 {
					answer = true
				}
				counter = 0
				for i := 0; i < 4; i++ {
					if game_board[i+k+h][2+k+h] == point {
						counter++
					}
				}
				if counter == 4 {
					answer = true
				}
				counter = 0
				for i := 0; i < 4; i++ {
					if game_board[0+k+h][i+k+h] == point {
						counter++
					}
				}
				if counter == 4 {
					answer = true
				}
				counter = 0
				for i := 0; i < 4; i++ {
					if game_board[1+k+h][i+k+h] == point {
						counter++
					}
				}
				if counter == 4 {
					answer = true
				}
				counter = 0
				for i := 0; i < 4; i++ {
					if game_board[2+k+h][i+k+h] == point {
						counter++
					}
				}
				if counter == 4 {
					answer = true
				}
				counter = 0
			}
		}
	}
	//if board_size == 4 {
	//	for h := 0; h < 2; h++ {
	//		for k := 0; k < 0; k++ {
	//			for i := 0; i < 3; i++ {
	//				if game_board[i+k][i+h] == point {
	//					counter++
	//				}
	//			}
	//			if counter == 3 {
	//				answer = true
	//			}
	//			counter = 0
	//			for i := 2; i >= 0; i-- {
	//				if game_board[i+k][i+h] == point {
	//					counter++
	//				}
	//			}
	//			if counter == 3 {
	//				answer = true
	//			}
	//			counter = 0
	//			for i := 0; i < 3; i++ {
	//				if game_board[i+k][0+h] == point {
	//					counter++
	//				}
	//			}
	//			if counter == 3 {
	//				answer = true
	//			}
	//			counter = 0
	//			for i := 0; i < 3; i++ {
	//				if game_board[i+k][1+h] == point {
	//					counter++
	//				}
	//			}
	//			if counter == 3 {
	//				answer = true
	//			}
	//			counter = 0
	//			for i := 0; i < 3; i++ {
	//				if game_board[i+k][2+h] == point {
	//					counter++
	//				}
	//			}
	//			if counter == 3 {
	//				answer = true
	//			}
	//			counter = 0
	//			for i := 0; i < 3; i++ {
	//				if game_board[0+k][i+h] == point {
	//					counter++
	//				}
	//			}
	//			if counter == 3 {
	//				answer = true
	//			}
	//			counter = 0
	//			for i := 0; i < 3; i++ {
	//				if game_board[1+k][i+h] == point {
	//					counter++
	//				}
	//			}
	//			if counter == 3 {
	//				answer = true
	//			}
	//			counter = 0
	//			for i := 0; i < 3; i++ {
	//				if game_board[2+k][i+h] == point {
	//					counter++
	//				}
	//			}
	//			if counter == 3 {
	//				answer = true
	//			}
	//			counter = 0
	//		}
	//	}
	//}
	return answer
}
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
		selected_coord.X = rand.Intn(board_size-0) + 0
		selected_coord.Y = rand.Intn(board_size-0) + 0
		if can_place_point() == true {
			game_board[selected_coord.X][selected_coord.Y] = point
			can = 0
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
	win := 0
	for win != 1 {
		player_step()
		show_board()
		if win_check() == true {
			win = 1
			fmt.Println()
			fmt.Println(point + " Win!")
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
			win = 1
			fmt.Println()
			fmt.Println(point + " Win!")
		}
		//Changing point again
		point_swicher()
	}
}
