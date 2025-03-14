package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordfrequency(sentence string) map[string]int {

	frequencyCheck := make(map[string]int)
	words := strings.Fields(strings.ToLower(sentence))

	for _, word := range words {
		word = strings.Trim(word, ",.?!")
		frequencyCheck[word]++
	}

	return frequencyCheck
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a sentence: ")
	sentence, _ := reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence)

	fmt.Print("Enter the word to check the frequency: ")
	var word string
	fmt.Scanln(&word)

	frequency := wordfrequency(sentence)

	count, exists := frequency[strings.ToLower(word)]
	if exists {
		fmt.Printf("The word '%s' appears %d times.\n", word, count)
	} else {
		fmt.Printf("The word '%s' is not in the sentence.\n", word)
	}

}
