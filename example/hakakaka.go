package main

import (
	"fmt"
	"strings"
)

func main() {
	tabooWords := []string{
		"first",
		"sec",
	}

	str := "qwqe, q first, q eq qw eqwe sec qwe qwe "
	wordsArray := strings.Split(str, " ")

	for _, wordFromUserInputString := range wordsArray {
		found := false
		tabooWordLen := 0
		for _, tabooWord := range tabooWords {
			if strings.HasPrefix(wordFromUserInputString, tabooWord) {
				found = true
				tabooWordLen = len(tabooWord)
				break
			}
		}

		if !found {
			fmt.Printf("%s ", wordFromUserInputString)
			continue
		}

		for i := 0; i < len(wordFromUserInputString); i++ {
			if i < tabooWordLen {
				fmt.Print("*")
			} else {
				fmt.Print(string(wordFromUserInputString[i]))
			}
		}
		fmt.Print(" ")
	}

}


in := bufio.NewReader(os.Stdin)
line, err := in.ReadString('\n')
if err != nil {
fmt.Println(err.Error())
return
}

fmt.Println(line)