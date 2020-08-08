package anytime

import (
	"testing"
	"time"
)

func init() {
	loc, err := time.LoadLocation("PDT")
	if err != nil {
		loc = time.FixedZone("PDT", -7*60*60)
	}
	time.Local = loc
}

func TestJSON_repeatability(t *testing.T) {
	at := &Time{}
	data := []byte(`"2020-08-20T01:02:03+09:00"`)

	if err := at.UnmarshalJSON(data); err != nil {
		t.Errorf("unexpected Time.UnmarshalJSON error. error = %v", err)
	}
	got, err := at.MarshalJSON()
	if err != nil {
		t.Errorf("unexpected Time.MarshalJSON error. error = %v", err)
	}
	if string(data) != string(got) {
		t.Errorf("got must be equal data. got = %s, data = %s", string(got), string(data))
	}
}
