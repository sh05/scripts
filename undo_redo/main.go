package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var cmd string
	history := make([]int, 0)
	tmp := make([]int, 0)
	sum := 0

	for cmd != "exit" {
		fmt.Scan(&cmd)
		switch cmd {
		case "roll":
			dice := get_rand()
			sum += dice
			history = append(history, dice)
			tmp = make([]int, 0)
			fmt.Println("got: " + fmt.Sprint(dice) + "!")
		case "undo":
			if len(history) > 0 {
				sum -= history[len(history)-1]
				move_last_elm(&tmp, &history)
			} else {
				fmt.Println("cant undo any more!")
			}
		case "redo":
			if len(tmp) > 0 {
				sum += tmp[len(tmp)-1]
				move_last_elm(&history, &tmp)
			} else {
				fmt.Println("cant redo any more!")
			}
		default:
			cmd = "exit"
		}

		fmt.Println(sum)
		fmt.Println(history)
		fmt.Println(tmp)
	}
}

func get_rand() int {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	return rand.Intn(6) + 1
}

func move_last_elm(a *[]int, b *[]int) {
	*a = append(*a, (*b)[len(*b)-1])
	*b = (*b)[:len(*b)-1]
}