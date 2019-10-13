package time

import "time"

// CurrentDate returns current date
func CurrentDate() string {
	return time.Now().Format("02-01-2006")
}

// CurrentTime returns current time
func CurrentTime() string {
	return time.Now().Format("15:04:05")
}

// CurrentLongTime returns current time till microsecond
func CurrentLongTime() string {
	return time.Now().Format("15:04:05.000000")
}

// DateTime returns current date-time
func DateTime(format string) string {
	return time.Now().Format(format)
}

// Unix returns current time in UnixNano format
func Unix() int64 {
	return time.Now().UnixNano()
}
