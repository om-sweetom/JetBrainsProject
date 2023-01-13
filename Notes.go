package main

import (
	"fmt"
	"strings"
)

func main() {
	/*var i int
	var out int = 1
	fmt.Scan(&i)

	for ; i > 1; i-- {
		out *= i


	fmt.Println(out)

	 }*/

	/*
	   matrix := [][]int{
	   {1, 2, 3},
	   {4, 5, 6},
	   {7, 8, 9},
	   }
	   var num int
	   fmt.Scanln(&num)

	   // Use a nested for...range loop to print each element
	   // of the matrix in a new line, multiplied by 'num'
	   for _, row := range matrix {
	   for _, ? := range ? {
	   fmt.Println(? * num)
	   }
	   }
	*/
	/*file, err := os.Open("numbers.txt") // open 'numbers.txt' here with the os.Open function
		if err != nil {
			panic(err)
		}
		defer file.Close()

		var oddNumberCount int // oddNumberCount will be used to count the amount of odd numbers in the file

		scanner := bufio.NewScanner(file) // create a new scanner for the file

		// Use the for scanner.Scan() loop to read the file line by line
		for scanner.Scan() {
			// DO NOT delete the code block below, it parses the read number into an integer:
			var number int
			number, err = strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}

			// Write the code that checks if 'number' is odd and then increments 'oddNumberCount' below:
			if number%2 == 0 {
				oddNumberCount++
			}
		}
		fmt.Println(oddNumberCount) // print the total oddNumberCount here!
	}
	*/

	/*injuredPlayers := map[string]int{
		"Aguero":  93,
		"Dembele": 45,
		"Pedri":   4,
	}

	var player string

	fmt.Scanln(&player) // take input of the player name here within the 'player' variable

	// implement the val, ok check below to verify if the player is on the map
	if val, ok := injuredPlayers[player]; ok {
		fmt.Println("The player", player, "will recover in", val, "days")
		return
	}
	fmt.Println("The player", player, "is not injured")

	*/

	var quote string
	quote = "hey, lousy. go fuck yourself"

	fmt.Println(quote[1:4])
	index := strings.Index(quote, "lousy")

	fmt.Println(index)

	contains := strings.Contains(quote, "lousy")
	fmt.Println(contains)

}
