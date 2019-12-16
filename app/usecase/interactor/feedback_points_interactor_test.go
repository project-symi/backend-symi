package interactor

import (
	"testing"
	"time"
)

func fakeTime(year int, month time.Month) {
	now = func() time.Time {
		return time.Date(year, month, 7, 12, 0, 0, 0, time.Local)
	}
}

func TestCalculateExpireDate(t *testing.T) {
	var tests = []struct {
		year  int
		month time.Month
		want  string
	}{
		{2008, 1, "2008-03-31"},
		{2009, 2, "2009-03-31"},
		{2010, 3, "2010-03-31"},
		{2011, 4, "2011-06-31"},
		{2012, 5, "2012-06-31"},
		{2013, 6, "2013-06-31"},
		{2014, 7, "2014-09-30"},
		{2015, 8, "2015-09-30"},
		{2016, 9, "2016-09-30"},
		{2017, 10, "2017-12-31"},
		{2018, 11, "2018-12-31"},
		{2019, 12, "2019-12-31"},
	}
	for _, test := range tests {
		fakeTime(test.year, test.month)
		if got := calculateExpireDate(); got != test.want {
			t.Errorf("calculateExpireDate() != %q Year-Month = %d - %d", test.want, test.year, test.month)
		}
	}
	now = time.Now
}
