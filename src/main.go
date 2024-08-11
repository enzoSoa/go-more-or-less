package main

import "fmt"

func getUserGuess() int {
  var guess int
  fmt.Scanf("%v", &guess)
  
  return guess
}

func compareGuess(number, guess int) bool {
  switch {
    case guess < number:
      fmt.Println("More")
    case guess > number:
      fmt.Println("Less")
    default:
      return true
  }
  return false
} 

func main() {
  const NUMBER_TO_GUESS int = 23
  fmt.Println("Enter a number between 0 and 100:")

  gameOver := false
  var guesses []int

  for !gameOver {
    guess := getUserGuess()
    gameOver = compareGuess(NUMBER_TO_GUESS, guess)
    guesses = append(guesses, guess)
  }

  fmt.Println("GG!!", "You won in", len(guesses), "guesses")
}
