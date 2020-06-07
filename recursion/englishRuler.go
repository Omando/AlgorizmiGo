package recursion

import (
	"fmt"
)

func DrawEnglishRuler(dashCount int) {
	print(dashCount)
	drawInternal(dashCount)
	print(dashCount)
}

func drawInternal(dashCount int) {
	// Exit condition
	if dashCount == 0 {
		return
	}

	drawInternal(dashCount - 1)
	print(dashCount)
	drawInternal(dashCount - 1)
}

func print(dashCount int) {
	for i := 0; i < dashCount; i++ {
		fmt.Print("-")
	}
	fmt.Println("-")
}
