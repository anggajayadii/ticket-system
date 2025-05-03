package entity

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

type Date time.Time

// Implement MarshalJSON untuk format JSON
func (d Date) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	if t.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, t.Format("2006-01-02"))), nil
}

// Implement UnmarshalJSON untuk parsing dari JSON
func (d *Date) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		*d = Date(time.Time{})
		return nil
	}

	// Remove quotes if present
	if len(s) > 0 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

// Implement Scanner interface untuk database
func (d *Date) Scan(value interface{}) error {
	if value == nil {
		*d = Date(time.Time{})
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		*d = Date(v)
	case []byte:
		t, err := time.Parse("2006-01-02", string(v))
		if err != nil {
			return err
		}
		*d = Date(t)
	default:
		return errors.New("invalid type for Date")
	}
	return nil
}

// Implement Valuer interface untuk database
func (d Date) Value() (driver.Value, error) {
	t := time.Time(d)
	if t.IsZero() {
		return nil, nil
	}
	return t.Format("2006-01-02"), nil
}

// Format untuk keperluan string representation
func (d Date) String() string {
	return time.Time(d).Format("2006-01-02")
}
