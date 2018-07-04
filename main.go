package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"time"
	"os"
	"strconv"
	"strings"
)

const (
	MIN             = 1
	MAX             = 100
	GUESS_TOO_HIGH  = 1
	GUESS_CORRECT   = 0
	GUESS_TOO_LOW   = -1
	GUESS_NOT_VALID = "Please enter a valid number (%d to %d, inclusive)\n"
)

var reader = bufio.NewReader(os.Stdin)

type Game struct {
	answer int
}

// Initializes the game with a random number between 1 and 100, inclusive.
func (g *Game) initialize() {
	rand.Seed(time.Now().Unix())
	g.answer = rand.Intn(MAX-MIN) + MIN
}

// Prompts the user for input and verifies the input is valid.
func (g *Game) takeGuess() int {
	fmt.Print("Please take a guess: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	guess, err := strconv.Atoi(text)
	if err != nil || guess < MIN || guess > MAX {
		fmt.Printf(GUESS_NOT_VALID, MIN, MAX)
		g.takeGuess()
	}
	return g.checkAnswer(guess)
}

// Checks whether the given guess is the correct answer.
func (g *Game) checkAnswer(guess int) int {
	if guess == g.answer {
		return GUESS_CORRECT
	} else if guess < g.answer {
		return GUESS_TOO_LOW
	} else {
		return GUESS_TOO_HIGH
	}
}

func main() {
	g := new(Game)
	g.initialize()
	fmt.Printf("I am thinking of a number between %d and %d (inclusive).\n", MIN, MAX)
	playing := true
	var result int
	var guesses = 1
	for playing {
		result = g.takeGuess()
		if result == GUESS_CORRECT {
			break
		} else if result == GUESS_TOO_LOW {
			fmt.Println("Too low.")
		} else if result == GUESS_TOO_HIGH {
			fmt.Println("Too high.")
		}
		guesses++
	}
	fmt.Printf("You guessed correctly!  It took you %d tries.\n", guesses)
}
