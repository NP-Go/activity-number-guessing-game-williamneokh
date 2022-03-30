package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//This code creates random number and place as variable
	rand.Seed(time.Now().UnixNano())
	hiddenNumber := int64(rand.Intn(100))
	var userInput string
	fmt.Println(hiddenNumber)

	//var input int

	for i := 0; i <= 4; i++ {
		fmt.Println("Guess the secret number from 1-100")
		fmt.Printf("You have %v tries\n", 5-i)
		fmt.Println("To end the game, enter 101")
		_, _ = fmt.Scanln(&userInput)
		guessNum, err := strconv.ParseInt(userInput, 10, 0)
		if err != nil {
			log.Fatal("Enter only numbers from 1-101")
		}
		if guessNum == hiddenNumber {
			fmt.Println("Congrats! You guess correctly!")
			return
		} else if guessNum == 101 {
			fmt.Println("You have end the game.")
			return
		} else if guessNum > 101 || guessNum < 1 {
			fmt.Println("You have give up 1 try, please use 1-100 and 101 to exit the game.")
		}

	}
	fmt.Println("Game Over!")
	//Insert code here
}
