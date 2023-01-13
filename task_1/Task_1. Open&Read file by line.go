package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var userIN string
	fmt.Scan(&userIN)

	file, err := os.Open(userIN)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmp := scanner.Text()
		words = append(words, strings.ToLower(tmp))

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for {
		in := bufio.NewReader(os.Stdin)
		userStr, err := in.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		lowerUserStr := strings.ToLower(strings.TrimSpace(userStr))

		if lowerUserStr == "exit" {
			fmt.Println("Bye!")
			break
		}

		wordsAr := strings.Split(userStr, " ")
		for _, userWord := range wordsAr {
			found := false
			tabooLen := 0
			tmpWord := strings.ToLower(userWord)
			for _, taboo := range words {
				if strings.HasPrefix(tmpWord, taboo) {
					found = true
					tabooLen = len(taboo)
					break
				}
			}

			if !found {
				fmt.Printf("%s ", userWord)
				continue
			}

			for i := 0; i < len(userWord); i++ {
				if i < tabooLen {
					fmt.Print("*")
				} else {
					fmt.Print(string(userWord[i]))
				}
			}

			fmt.Print(" ")

		}

	}

}
