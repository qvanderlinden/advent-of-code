package lib

func CountCharOccurrences(str string, char rune) (int) {
	count := 0
	for _, c := range str {
		if (c == char) {
			count++
		}
	}
	return count
}
