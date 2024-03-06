package utils

import (
	"fmt"
	"strings"
	"time"
)

// CustomTime provides an example of how to declare a new time Type with a custom formatter.
// Note that time.Time methods are not available, if needed you can add and cast like the String method does
// Otherwise, only use in the json struct at marshal/unmarshal time.
type CustomTime time.Time

const ctLayout = "2006-01-02"

// UnmarshalJSON Parses the json string in the custom format
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(ctLayout, s)
	*ct = CustomTime(nt)
	return
}

// MarshalJSON writes a quoted string in the custom format
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.String()), nil
}

// String returns the time in the custom format
func (ct *CustomTime) String() string {
	t := time.Time(*ct)
	return fmt.Sprintf("%q", t.Format(ctLayout))
}
