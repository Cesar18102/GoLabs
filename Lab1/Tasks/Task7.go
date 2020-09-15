package Tasks

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Task7() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter count of words: ")
	count := readInt(scanner);

	var words []string;
	for i := 0; i < count; i++ {
		fmt.Print("Enter word: ")
		words = append(words, readText(scanner));
	}

	var palindromes []string;
	for i := 0; i < len(words); i++ {
		if isPalindrome(words[i]) {
			palindromes = append(palindromes, words[i]);
		}
	}

	fmt.Printf("Palindrome words: %v", palindromes);
}

func readText(scanner *bufio.Scanner) string {
	scanner.Scan();
	return scanner.Text();
}

func readInt(scanner *bufio.Scanner) int {
	num, _ := strconv.Atoi(readText(scanner));
	return num;
}

func isPalindrome(str string) bool {
	return str == reverse(str);
}

func reverse(str string) string {
	var reversed string;
	for _, rune := range str {
		reversed = string(rune) + reversed;
	}
	return reversed;
}