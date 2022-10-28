//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
type Time struct {
	hour, minute, second int
}

type TimeParseError struct {
	msg   string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

func ParseTime(input string) (Time, error) {
	components := strings.Split(input, ":")
	if len(components) != 3 {
		return Time{0, 0, 0}, &TimeParseError{"invalid number of time components", input}
	} else {
		hour, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{0, 0, 0}, &TimeParseError{fmt.Sprintf("error parsing hour: %v to string", err), input}
		}
		minute, err := strconv.Atoi(components[1])
		if err != nil {
			return Time{0, 0, 0}, &TimeParseError{fmt.Sprintf("error parsing minute: %v to string", err), input}
		}
		second, err := strconv.Atoi(components[2])
		if err != nil {
			return Time{0, 0, 0}, &TimeParseError{fmt.Sprintf("error parsing second: %v to string", err), input}
		}

		if hour > 23 || hour < 0 {
			return Time{0, 0, 0}, &TimeParseError{"hour out of range: 0 <= hour <= 23", fmt.Sprintf("%v", hour)}
		}
		if minute > 59 || minute < 0 {
			return Time{0, 0, 0}, &TimeParseError{"minute out of range: 0 <= minute <= 59", fmt.Sprintf("%v", minute)}
		}
		if second > 59 || second < 0 {
			return Time{0, 0, 0}, &TimeParseError{"second out of range: 0 <= second <= 59", fmt.Sprintf("%v", second)}
		}

		return Time{hour, minute, second}, nil
	}
}
