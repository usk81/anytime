package anytime

import "time"

// Date returns the Time corresponding to
//	yyyy-mm-dd hh:mm:ss + nsec nanoseconds
// in the appropriate zone for that time in the given location.
//
// The month, day, hour, min, sec, and nsec values may be outside
// their usual ranges and will be normalized during the conversion.
// For example, October 32 converts to November 1.
//
// A daylight savings time transition skips or repeats times.
// For example, in the United States, March 13, 2011 2:15am never occurred,
// while November 6, 2011 1:15am occurred twice. In such cases, the
// choice of time zone, and therefore the time, is not well-defined.
// Date returns a time that is correct in one of the two zones involved
// in the transition, but it does not guarantee which.
//
// Date panics if loc is nil.
func Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	t := Time{}
	t.Time = time.Date(year, month, day, hour, min, sec, nsec, loc)
	return t
}

// Now returns the current local time.
func Now() Time {
	t := Time{}
	t.Time = time.Now()
	return t
}

// Parse parses a formatted string and returns the time value it represents.
// The layout defines the format by showing how the reference time.
func Parse(layout, value string) (Time, error) {
	t := Time{}
	tt, err := time.Parse(layout, value)
	if err != nil {
		return t, err
	}
	t.Time = tt
	return t, nil
}

// ParseInLocation is like Parse but differs in two important ways.
// First, in the absence of time zone information, Parse interprets a time as UTC;
// ParseInLocation interprets the time as in the given location.
// Second, when given a zone offset or abbreviation, Parse tries to match it
// against the Local location; ParseInLocation uses the given location.
func ParseInLocation(layout, value string, loc *time.Location) (Time, error) {
	t := Time{}
	tt, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		return t, err
	}
	t.Time = tt
	return t, nil
}

// Unix returns the local Time corresponding to the given Unix time,
// sec seconds and nsec nanoseconds since January 1, 1970 UTC.
// It is valid to pass nsec outside the range [0, 999999999].
// Not all sec values have a corresponding time value. One such
// value is 1<<63-1 (the largest int64 value).
func Unix(sec int64, nsec int64) Time {
	t := Time{}
	t.Time = time.Unix(sec, nsec)
	return t
}
