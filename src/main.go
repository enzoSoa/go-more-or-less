package main

import "fmt"

func getUserGuess() int {
  var guess int
  fmt.Scanf("%v", &guess)
  
  return guess
}

func compareGuess(numberToGuess, guess int) bool {
  switch {
    case guess < numberToGuess:
      fmt.Println("More")
    case guess > numberToGuess:
      fmt.Println("Less")
    default:
      return true
  }
  return false
} 

func playNumbersSequence(numbersToGuess ...int) {
  fmt.Println("This sequence is", len(numbersToGuess), "numbers long")

  var guessesByNumber = make(map[int][]int)
  for numberIndex, numberToGuess := range numbersToGuess {
    fmt.Println("Enter a guess for the number", numberIndex + 1, ":")

    guessesByNumber[numberIndex] = []int{}
    numberGuessed := false
    for !numberGuessed  {
      guess := getUserGuess()
      guessesByNumber[numberIndex] = append(guessesByNumber[numberIndex], guess)
      numberGuessed = compareGuess(numberToGuess, guess)
    }
  }

  fmt.Println("GG!! Score :")

  totalGuessesCount := 0
  for guessesIndex, guesses := range guessesByNumber {
    guessesCount := len(guesses)
    totalGuessesCount += guessesCount
    fmt.Println("number", guessesIndex, "-", guessesCount, "guesses")
  }
  
  fmt.Println("total guesses :", totalGuessesCount)
}

func main() {
  playNumbersSequence(23, 32)
}
