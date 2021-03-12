package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomnumber := rand.Intn(5) + 1
	//fmt.Print("secret is ", randomnumber)
	fmt.Println("\nA secret number was created between 1-5, go ahead and guess.")

	var guess int
	//fmt.Scanln(&guess)

	fmt.Println("Guess was : ", guess)

	if guess == randomnumber {
		fmt.Println("Congrats, you were right!")
	} else if guess > randomnumber {
		fmt.Println("Your guess was too high!")
	} else if guess < randomnumber {
		fmt.Println("Your guess was too low!")
	}

	if guess > 5 {
		fmt.Println("Your number must be between 1-5!")
	} else if guess < 0 {
		fmt.Println("Your number must be between 1-5!")
	}
	fmt.Println("The secret number was ", randomnumber)
}
