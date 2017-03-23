package intersights

import "time"

const timeFormat = "Mon Jan 2 15:04:05.000 -0700 MST 2006"

type Time struct {
	time.Time
}

func (j Time) format() string {
	return j.Time.Format(timeFormat)
}

func (j Time) MarshalText() ([]byte, error) {
	return []byte(j.format()), nil
}

func (j Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + j.format() + `"`), nil
}
