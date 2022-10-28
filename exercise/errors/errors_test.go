package timeparse

import "testing"

//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{"22:07:13", true},
		{"1:3:44", true},
		{"bad", false},
		{"1:-3:44", false},
		{"0:59:59", true},
		{"", false},
		{"11:22", false},
		{"aa:bb:cc", false},
		{"13:13:", false},
		{"12:13:14", false},
	}

	for _, data := range table {
		_, err := ParseTime(data.time)
		if data.ok && err != nil {
			t.Errorf("%v: %v, error should be nil", data.time, err)
		}
	}
}