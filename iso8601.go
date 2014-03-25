// The iso8601 package encodes and decodes time.Time in JSON in
// ISO 8601 format, without subsecond resolution or time zone info.
package iso8601

import "time"

const Format = "2006-01-02T15:04:05"
const jsonFormat = `"` + Format + `"`

type Time time.Time

func (it Time) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(it).Format(jsonFormat)), nil
}

func (it *Time) UnmarshalJSON(data []byte) error {
	t, err := time.Parse(jsonFormat, string(data))
	if err == nil {
		*it = Time(t)
	}

	return err
}

func (it Time) String() string {
	return time.Time(it).String()
}
