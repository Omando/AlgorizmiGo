package recursion

func GetCharacterCount(s string, c byte) int {
	if len(s) == 0 {
		return 0
	}

	return doGetCharacterCount(s, c, 0, 0)

}

func doGetCharacterCount(s string, c byte, index int, runningCount int) int {
	// Exit condition
	if len(s) == index {
		return runningCount
	}

	// Calculation
	if s[index] == c {
		runningCount++
	}
	return doGetCharacterCount(s, c, index+1, runningCount)
}
