package time

import "time"

// The 5 laws of API dates and times
// Link: http://apiux.com/2013/03/20/5-laws-api-dates-and-times/
//
// Law #1: Use ISO-8601 for your dates (ISO 8601 ~ RFC3339)
// Law #2: Accept any timezone
// Law #3: Store it in UTC
// Law #4: Return it in UTC
// Law #5: Don’t use time if you don’t need it
//
// RFC3339 = "2006-01-02T15:04:05Z07:00"
// Where:
// - Date: 2006-01-02
// - Time: 15:04:05
// - Zone: 07:00

// Time wrapper for the default time
type Time struct {
	time.Time
}

// Now returns the current time in Coordinated Universal Time (UTC)
func Now() Time {
	return Time{time.Now().UTC()}
}

// Unix returns the time in Unix format (Seconds since Jan 01 1970. (UTC))
func Unix() int64 {
	return time.Now().Unix()
}

// String retuns the time into unified string format (~ISO 8601)
func (t Time) String() string {
	return t.Format(time.RFC3339)
}

// Parse returns the time (RFC3339 format)
func Parse(val string) (time.Time, error) {
	return time.Parse(time.RFC3339, val)
}

// Since returns the duration
func Since(t Time) time.Duration {
	return time.Since(t.Time)
}
