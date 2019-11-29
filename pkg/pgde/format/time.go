package format

import (
	"fmt"
	"time"
)

func (u TimeFormat_IntUnit) Time(t time.Time) int64 {
	switch u {
	case TimeFormat_nanosecond:
		return t.UnixNano()
	case TimeFormat_microsecond:
		return t.Unix()*(1000*1000) + int64(t.Nanosecond())/1000
	case TimeFormat_millisecond:
		return t.Unix()*(1000) + int64(t.Nanosecond())/(1000*1000)
	case TimeFormat_second:
		return t.Unix()
	case TimeFormat_minute:
		return t.Unix() / 60
	case TimeFormat_hour:
		return t.Unix() / (60 * 60)
	default:
		panic(fmt.Errorf("unkown TimeFormat value: %v", u))
	}
}

func (u TimeFormat_IntUnit) Duration(t time.Duration) int64 {
	switch u {
	case TimeFormat_nanosecond:
		return t.Nanoseconds()
	case TimeFormat_microsecond:
		return t.Microseconds()
	case TimeFormat_millisecond:
		return t.Milliseconds()
	case TimeFormat_second:
		return int64(t / time.Second)
	case TimeFormat_minute:
		return int64(t / time.Minute)
	case TimeFormat_hour:
		return int64(t / time.Hour)
	default:
		panic(fmt.Errorf("unkown TimeFormat value: %v", u))
	}
}
