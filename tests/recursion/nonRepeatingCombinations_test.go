package recursion

import (
	"AlgorizmiGo/recursion"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNonRepeatingCombinations(t *testing.T) {
	// Arrange
	universe := []string{"A", "B", "C"}

	// Act
	var sequences []string = recursion.GetNonRepeatingCombinations(universe)

	// Assert
	assert.EqualValues(t, 6, len(sequences))
	assert.Equal(t, "ABC", sequences[0])
	assert.Equal(t, "ACB", sequences[1])
	assert.Equal(t, "BAC", sequences[2])
	assert.Equal(t, "BCA", sequences[3])
	assert.Equal(t, "CAB", sequences[4])
	assert.Equal(t, "CBA", sequences[5])
}

func TestNonRepeatingCombinationsWith4Chars(t *testing.T) {
	// Todo: Consider adding a test table with input and expected count
	// Arrange
	universe := []string{"A", "B", "C", "D"}

	// Act
	var sequences []string = recursion.GetNonRepeatingCombinations(universe)

	// Assert
	assert.EqualValues(t, 24, len(sequences))
	assert.Equal(t, "ABCD", sequences[0])
	assert.Equal(t, "ABDC", sequences[1])
}
