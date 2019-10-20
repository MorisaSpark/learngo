package main

import (
	"github.com/garyburd/redigo/redis"
	"strings"
	"time"
)

type StringMap map[string]string

func GetMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

var (
	timeTemplate         = "2006-01-02-15-04-05"
	timeTemplateStandard = "2006-01-02 15:04:05"
	WeekDayMap           = map[string]string{
		"Monday":    "2",
		"Tuesday":   "3",
		"Wednesday": "4",
		"Thursday":  "5",
		"Friday":    "6",
		"Saturday":  "7",
		"Sunday":    "1",
	}
)

const (
	UNREPEATED = "unrepeated"
	HOUR       = "@hourly "
	DAY        = "@everyday "
	WEEK       = "@weekly "
	MONTH      = "@monthly "
	YEAR       = "@yearly "
)

func CreateCronSpec(firstAt int64, frequency string) string {
	cronSpec := ""
	st, err := StampToTime(firstAt)
	if err != nil {
		return ""
	}
	// 指定某个时间点执行，好像已经被淘汰了，所以将每年和某时
	if frequency == UNREPEATED || frequency == YEAR {
		cronSpec = st.Minute + " " + st.Hour + " " + st.Day + " " + st.Month + " ?"
	} else if frequency == HOUR {
		cronSpec = st.Minute + " * * * ?"
	} else if frequency == DAY {
		cronSpec = st.Minute + " " + st.Hour + " * * ?"
	} else if frequency == WEEK {
		cronSpec = st.Minute + " " + st.Hour + " ? * " + st.Week
	} else if frequency == MONTH {
		cronSpec = st.Minute + " " + st.Hour + " " + st.Day + " * ?"
		if st.Day == "31" {
			cronSpec = st.Minute + " " + st.Hour + " L * ?"
		}
	}

	return cronSpec
}

type SplitTime struct {
	Year   string
	Month  string
	Day    string
	Hour   string
	Minute string
	Second string
	Week   string
}

func StampToTime(timestamp int64) (*SplitTime, error) {
	format := time.Unix(timestamp/1000, 0).Format(timeTemplate)
	split := strings.Split(format, "-")
	parse, err := time.Parse(timeTemplate, format)
	if err != nil {
		return nil, err
	}
	week := WeekDayMap[parse.Weekday().String()]
	splitTime := &SplitTime{Year: split[0], Month: split[1], Day: split[2], Hour: split[3], Minute: split[4], Second: split[5], Week: week}
	return splitTime, nil
}

func ToTimestamp(temp string) int64 {
	timeStruct := Str2Time(temp)
	millisecond := timeStruct.UnixNano() / 1e6
	return millisecond
}
func TimestampToString(temp int64) string {
	unix := time.Unix(temp, 0)
	return unix.Format(timeTemplateStandard)
}
func Str2Time(formatTimeStr string) time.Time {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, formatTimeStr, loc) //使用模板在对应时区转化为time.time类型

	return theTime

}

var Pool *redis.Pool

func init() {
	Pool = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   1024,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
