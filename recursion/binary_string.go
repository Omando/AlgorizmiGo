package recursion

import "errors"

func BinaryString(number int) (string, error) {
	if number < 0 {
		return "NOT SUPPORTED!", errors.New("negative numbers are not supported")
	}

	if number == 0 {
		return "0", nil
	}

	var result []byte = doBinaryString(number, 0)
	return string(result), nil
}

func doBinaryString(number int, index int) []byte {
	// Exit condition
	if number == 0 {
		return make([]byte, index)
	}

	// Calculations
	newNumber := number / 2
	var remainder byte = byte(number % 2)
	var binaries []byte = doBinaryString(newNumber, index+1)

	// Assign results
	if remainder == 0 {
		binaries[len(binaries)-index-1] = '0'
	} else {
		binaries[len(binaries)-index-1] = '1'
	}

	return binaries
}
