package Tasks

import "fmt"

func Task4() {
	str, substr := "abcdabc", "abc";
	count := countSubstrings(str, substr);
	fmt.Printf("%v is met in %v for %d times", substr, str, count);
}

func countSubstrings(str string, substr string) int {
	count, sublen, end := 0, len(substr), len(str) - len(substr);
	for i := 0; i <= end; i++ {
		if str[i:i + sublen] == substr {
			count++;
		}
	}
	return count;
}
