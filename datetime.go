package datatypes

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"time"
)

type DateTime time.Time

func (date *DateTime) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*date = DateTime(nullTime.Time)
	return
}

func (date DateTime) Value() (driver.Value, error) {
	data := time.Time(date)
	return time.Date(data.Year(), data.Month(), data.Day(), data.Hour(), data.Minute(), data.Second(), data.Nanosecond(), time.Time(date).Location()), nil
}

// GormDataType gorm common data type
func (date DateTime) GormDataType() string {
	return "date"
}

func (date DateTime) GobEncode() ([]byte, error) {
	return time.Time(date).GobEncode()
}

func (date *DateTime) GobDecode(b []byte) error {
	return (*time.Time)(date).GobDecode(b)
}

func (date DateTime) MarshalJSON() ([]byte, error) {
	t := time.Time(date)
	if y := t.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len("2006-01-02T15:04:05")+2)
	b = append(b, '"')
	b = t.AppendFormat(b, "2006-01-02T15:04:05")
	b = append(b, '"')
	return b, nil
}

func (date *DateTime) UnmarshalJSON(b []byte) error {

	data := b
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	t, err := time.Parse(`"`+"2006-01-02T15:04:05"+`"`, string(data))
	if err != nil {
		t, err = time.Parse(`"`+"2006-01-02 15:04:05"+`"`, string(data))
	}
	if err == nil {
		*date = DateTime(t)
	}
	return err
}
