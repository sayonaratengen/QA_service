package mapper

import "time"

func toRFC3339ms(t time.Time) string {
	return t.Truncate(time.Millisecond).Format("2006-01-02T15:04:05.000Z07:00")
}
