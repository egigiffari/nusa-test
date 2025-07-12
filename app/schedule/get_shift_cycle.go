package schedule

import "time"

func getShiftCycle(startDate time.Time, cycles []string, daysDifference int) string {
	startPointer := int(startDate.Weekday())
	// fmt.Printf("startPointer:  %v\n", startPointer)

	// startPattern := shiftPattern[startPointer-1]
	// fmt.Printf("startPattern:  %v\n", startPattern)

	endPointer := daysDifference % len(cycles)
	// fmt.Printf("endPointer (before): %v\n", endPointer)

	if startPointer > 0 {
		endPointer = (daysDifference + startPointer) % len(cycles)
	}

	if endPointer <= 0 {
		endPointer = len(cycles)
	}

	// fmt.Printf("endPointer (after): %v\n", endPointer)

	endPattern := cycles[endPointer-1]
	// fmt.Printf("endPattern:  %v\n", endPattern)
	return endPattern
}
