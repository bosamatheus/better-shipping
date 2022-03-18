package date

import "time"

const yyyyMMdd = "2006-01-02"

// FormatDate formats a date in the format YYYY-MM-DD.
func FormatDate(t time.Time) string {
	return t.Format(yyyyMMdd)
}
