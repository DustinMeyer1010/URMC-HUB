package utils

import "strconv"

func ExtractStartEndRange(start, end string) (int, int) {
	newStart := 0
	newEnd := 1500

	// Converts the string value to an int and update start if no error
	if value, convError := strconv.Atoi(start); convError == nil {
		newStart = value
	}

	// Converts the string value to an int and update end if no error
	if value, convError := strconv.Atoi(end); convError == nil {
		newEnd = value
	}

	// Checks for positive start and end values
	if newStart < 0 {
		newStart *= -1
	}

	if newEnd < 0 {
		newStart *= -1
	}

	// Checks to make sure start is less than end. If they are not
	// the values are swapped
	if newStart > newEnd {
		temp := newEnd
		newEnd = newStart
		newStart = temp
	}

	return newStart, newEnd

}
