package main

import "fmt"

func main() {
  const NUMBER_TO_GUESS int = 23
  fmt.Println("Enter a number between 0 and 100:")

  gameOver := false
  var guesses []int
  for !gameOver {
    var guess int
    fmt.Scanf("%v", &guess)

    guesses = append(guesses, guess)
    switch {
    case guess < NUMBER_TO_GUESS:
      fmt.Println("More")
    case guess > NUMBER_TO_GUESS:
      fmt.Println("Less")
    default:
      gameOver = true
      fmt.Println("GG!!", "You won in", len(guesses), "guesses")
    }
  }
}
