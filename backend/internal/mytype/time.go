package mytype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type JsonDate time.Time

// Implement Marshaler and Unmarshaler interface
func (j *JsonDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonDate(t)
	return nil
}

func (j JsonDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

// Maybe a Format function for printing your date
func (j JsonDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

func (j *JsonDate) Scan(value interface{}) error {
	if value == nil {
		*j = JsonDate{}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*j = JsonDate(v)
		return nil
	case []byte:
		t, err := time.Parse("2006-01-02", string(v))
		if err != nil {
			return fmt.Errorf("failed to parse date from []byte: %w", err)
		}
		*j = JsonDate(t)
		return nil
	case string:
		t, err := time.Parse("2006-01-02", v)
		if err != nil {
			return fmt.Errorf("failed to parse date from string: %w", err)
		}
		*j = JsonDate(t)
		return nil
	default:
		return fmt.Errorf("unsupported type %T for JsonDate.Scan", value)
	}
}

// Value implements the driver.Valuer interface
func (j JsonDate) Value() (driver.Value, error) {
	return time.Time(j), nil
}
