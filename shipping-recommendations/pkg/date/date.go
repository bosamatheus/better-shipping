package date

import "time"

const yyyyMMdd = "2006-01-02"

func FormatDate(t time.Time) string {
	return t.Format(yyyyMMdd)
}
