// The iso8601 package encodes and decodes time.Time in JSON in
// ISO 8601 format, without subsecond resolution or time zone info.
package iso8601

import "time"

const Format = "2006-01-02T15:04:05"
const jsonFormat = `"` + Format + `"`

var fixedZone = time.FixedZone("", 0)

type Time time.Time

// New constructs a new iso8601.Time instance from an existing
// time.Time instance.  This causes the nanosecond field to be set to
// 0, and its time zone set to a fixed zone with no offset from UTC
// (but it is *not* UTC itself).
func New(t time.Time) Time {
	return Time(time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		0,
		fixedZone,
	))
}

func (it Time) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(it).Format(jsonFormat)), nil
}

func (it *Time) UnmarshalJSON(data []byte) error {
	t, err := time.ParseInLocation(jsonFormat, string(data), fixedZone)
	if err == nil {
		*it = Time(t)
	}

	return err
}

func (it Time) String() string {
	return time.Time(it).String()
}
