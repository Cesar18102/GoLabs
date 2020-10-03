package Tasks

import "fmt"

type RuneString []rune;

func (str *RuneString) invertBinary(start int) {
	for i := start; i < len(*str); i++ {
		if (*str)[i] == '0' {
			(*str)[i] = '1';
		} else if (*str)[i] == '1' {
			(*str)[i] = '0';
		}
	}
}

func Task3() {
	fmt.Print("Enter source text: ");

	input := RuneString(readText());
	input.invertBinary(0);

	fmt.Printf("Inverted: %v", string(input));
}