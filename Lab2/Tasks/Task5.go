package Tasks

import (
	"fmt"
)

type CustomString string;

var SEPARATORS = []rune { ' ', '\t', '\n', '\r', ',', '.', '!', '?', '—' };

func isSeparator(r rune) bool {
	for _, v := range SEPARATORS {
		if r == v {
			return true;
		}
	}
	return false;
}

func (str CustomString) getWords() []CustomString {
	words := []CustomString { };
	runes := []rune(str);
	wordStart := -1;

	for i, v := range runes {
		separate, end := isSeparator(v), i == len(runes) - 1;

		if wordStart == -1 && !separate {
			wordStart = i;
		}

		if wordStart != -1 {
			if end {
				return append(words, CustomString(runes[wordStart:]));
			}

			if separate {
				words = append(words, CustomString(runes[wordStart:i]));
				wordStart = -1;
			}
		}
	}

	return words;
}

func join(strs []CustomString, separator CustomString) CustomString {
	if len(strs) == 0 {
		return "";
	}

	if len(strs) == 1 {
		return strs[0];
	}

	joined := strs[0];
	for i := 1; i < len(strs); i++ {
		joined += separator + strs[i];
	}

	return joined;
}

func Task5() {
	fmt.Print("Enter a phrase to parse: ");
	phrase := CustomString(readText());

	words := phrase.getWords();
	joined := join(words, ", ");

	fmt.Printf("Separated words: %s (%d words)", joined, len(words));
}
