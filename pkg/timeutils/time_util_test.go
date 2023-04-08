package timeutils

import "testing"

func TestParse(t *testing.T) {
	startDate, err := Parse("2021-06-10", DefaultDateFormat)
	t.Log(startDate, err)

	t.Log(GetRemainderSecOfCurrentDate())

}
