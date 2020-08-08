package anytime

import (
	"testing"
	"time"
)

var dateTests = []struct {
	year, month, day, hour, min, sec, nsec int
	z                                      *time.Location
	unix                                   int64
}{
	{2011, 11, 6, 1, 0, 0, 0, time.Local, 1320566400},   // 1:00:00 PDT
	{2011, 11, 6, 1, 59, 59, 0, time.Local, 1320569999}, // 1:59:59 PDT
	{2011, 11, 6, 2, 0, 0, 0, time.Local, 1320573600},   // 2:00:00 PST

	{2011, 3, 13, 1, 0, 0, 0, time.Local, 1300006800},   // 1:00:00 PST
	{2011, 3, 13, 1, 59, 59, 0, time.Local, 1300010399}, // 1:59:59 PST
	{2011, 3, 13, 3, 0, 0, 0, time.Local, 1300010400},   // 3:00:00 PDT
	{2011, 3, 13, 2, 30, 0, 0, time.Local, 1300008600},  // 2:30:00 PDT ≡ 1:30 PST
	{2012, 12, 24, 0, 0, 0, 0, time.Local, 1356336000},  // Leap year

	// Many names for Fri Nov 18 7:56:35 PST 2011
	{2011, 11, 18, 7, 56, 35, 0, time.Local, 1321631795},                      // Nov 18 7:56:35
	{2011, 11, 19, -17, 56, 35, 0, time.Local, 1321631795},                    // Nov 19 -17:56:35
	{2011, 11, 17, 31, 56, 35, 0, time.Local, 1321631795},                     // Nov 17 31:56:35
	{2011, 11, 18, 6, 116, 35, 0, time.Local, 1321631795},                     // Nov 18 6:116:35
	{2011, 10, 49, 7, 56, 35, 0, time.Local, 1321631795},                      // Oct 49 7:56:35
	{2011, 11, 18, 7, 55, 95, 0, time.Local, 1321631795},                      // Nov 18 7:55:95
	{2011, 11, 18, 7, 56, 34, 1e9, time.Local, 1321631795},                    // Nov 18 7:56:34 + 10⁹ns
	{2011, 12, -12, 7, 56, 35, 0, time.Local, 1321631795},                     // Dec -21 7:56:35
	{2012, 1, -43, 7, 56, 35, 0, time.Local, 1321631795},                      // Jan -52 7:56:35 2012
	{2012, int(time.January - 2), 18, 7, 56, 35, 0, time.Local, 1321631795},   // (Jan-2) 18 7:56:35 2012
	{2010, int(time.December + 11), 18, 7, 56, 35, 0, time.Local, 1321631795}, // (Dec+11) 18 7:56:35 2010
}

func TestDate(t *testing.T) {
	loc, err := time.LoadLocation("PDT")
	if err != nil {
		loc = time.FixedZone("PDT", -7*60*60)
	}
	time.Local = loc
	for _, tt := range dateTests {
		time := Date(tt.year, time.Month(tt.month), tt.day, tt.hour, tt.min, tt.sec, tt.nsec, tt.z)
		want := Unix(tt.unix, 0)
		if !time.Equal(want.Time) {
			t.Errorf("Date(%d, %d, %d, %d, %d, %d, %d, %s) = %v, want %v",
				tt.year, tt.month, tt.day, tt.hour, tt.min, tt.sec, tt.nsec, tt.z,
				time, want)
		}
	}
}
