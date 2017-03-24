package intersights

import "time"

type Time struct {
	time.Time
}

func (j Time) format() string {
	return j.Time.Format(time.RFC3339Nano)
}

func (j Time) MarshalText() ([]byte, error) {
	return []byte(j.format()), nil
}

func (j Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + j.format() + `"`), nil
}
