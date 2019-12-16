package interactor

import (
	"testing"
	"time"
)

func fakeTime (year int, month time.Month) {
	now = func() time.Time {
		return time.Date(year, month, 7, 12, 0, 0, 0, time.Local)
	}
}

func TestCalculateExpireDate(t *testing.T) {
	var tests = []struct{
		year int
		month time.Month
		want string
	}{
		{1989, 1, "1989-03-31"},
		{1989, 2, "1989-03-31"},
		{1989, 3, "1989-03-31"},
		{2019, 4, "2019-06-31"},
		{2019, 5, "2019-06-31"},
		{2019, 6, "2019-06-31"},
		{2015, 7, "2015-09-30"},
		{2015, 8, "2015-09-30"},
		{2015, 9, "2015-09-30"},
		{2012, 10, "2012-12-31"},
		{2012, 11, "2012-12-31"},
		{2012, 12, "2012-12-31"},
	}
	for _, test := range tests {
		fakeTime(test.year, test.month)
		if got := calculateExpireDate(); got != test.want {
			t.Errorf("calculateExpireDate() != %q Year-Month = %d - %d", test.want, test.year, test.month)
		}
	}
	now = time.Now
}