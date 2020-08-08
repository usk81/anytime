package anytime

import (
	"errors"
	"time"
)

// Time ...
type Time struct {
	time.Time
}

var (
	// MarshalerFormat is used by MarshalText and MarshalJSON.
	MarshalerFormat = time.RFC3339Nano
	// UnmarshalerFormat is used by UnmarshalText and UnmarshalJSON.
	UnmarshalerFormat = time.RFC3339
)

// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in the format specified by MarshalerFormat.
func (t Time) MarshalJSON() ([]byte, error) {
	if y := t.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}
	return []byte(`"` + t.Format(MarshalerFormat) + `"`), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in the format specified by UnmarshalerFormat.
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return
	}
	// Fractional seconds are handled implicitly by Parse.
	t.Time, err = time.Parse(`"`+UnmarshalerFormat+`"`, string(data))
	return
}

// MarshalText implements the encoding.TextMarshaler interface.
// The time is formatted in the format specified by MarshalerFormat.
func (t Time) MarshalText() ([]byte, error) {
	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalText: year outside of range [0,9999]")
	}
	return []byte(t.Format(MarshalerFormat)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The time is expected to be the format specified by UnmarshalerFormat.
func (t *Time) UnmarshalText(data []byte) (err error) {
	// Fractional seconds are handled implicitly by Parse.
	t.Time, err = time.Parse(UnmarshalerFormat, string(data))
	return
}
