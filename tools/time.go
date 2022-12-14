package tools

import (
	log "github.com/sirupsen/logrus"
	"time"
)

// TimeStrToTime 时间格式字符串转时间格式
func TimeStrToTime(timeStr string) time.Time {
	var layout = "2006-01-02 15:04:05" //转换的时间字符串带秒则 为 2006-01-02 15:04:05
	timeVal, err := time.ParseInLocation(layout, timeStr, time.Local)
	if err != nil {
		log.Errorf("TimeStr To Time failed, err: %v", err)
	}
	return timeVal
}

// UnixTimeToTime 时间戳转时间格式
func UnixTimeToTime(unixTime int64) time.Time {
	return time.Unix(unixTime, 0)
}

/*
	TimeDayStartTime 获取参数时间当天零点零分零秒的时间
*/
func TimeDayStartTime(baseTime time.Time) time.Time {
	return time.Date(baseTime.Year(), baseTime.Month(), baseTime.Day(), 0, 0, 0, 0, baseTime.Location())
}

/*
	BeforeDayStartTime 获取以参数时间为基准的前n天零点零分零秒的时间
*/
func BeforeDayStartTime(baseTime time.Time, n int64) time.Time {
	t1 := time.Date(baseTime.Year(), baseTime.Month(), baseTime.Day(), 0, 0, 0, 0, baseTime.Location())
	return t1.Add(-24 * time.Hour * time.Duration(n))
}

/*
	GetDateTime 获取系统当前时间，格式化日期
*/
func GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

/*
	GetUnixTime 获取系统当前事件，格式化时间戳
*/
func GetUnixTime() int64 {
	return time.Now().Unix()
}

/*
	GetMillisecondUnixTime 获取系统当前毫秒时间戳
*/
func GetMillisecondUnixTime() int64 {
	return time.Now().UnixNano() / 1e6
}

/*
	GetTimestampByDuration 获取时间戳 - 指定多少时间前后
 	@param duration string 描述取前后多少时间。比如：-1m=前一分钟，1h=后一小时
 	@example tools.GetTimestampByDuration("-24h")
*/
func GetTimestampByDuration(duration string) int64 {
	currentTime := time.Now()
	d, _ := time.ParseDuration(duration)
	res := currentTime.Add(d)

	return res.Unix()
}

/*
	DateToTimestamp 日期转时间戳
*/
func DateToTimestamp(date string) int64 {
	if date == "" {
		return 0
	}

	loc, _ := time.LoadLocation("Local")
	timestamp, _ := time.ParseInLocation("2006-01-02 15:04:05", date, loc)

	return timestamp.Unix()
}

/*
	TimestampToDate 时间戳转日期
*/
func TimestampToDate(timestamp int64) string {
	if timestamp == 0 {
		return ""
	}

	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

/*
	TodayStartTime 获取今天零点零分零秒的时间
*/
func TodayStartTime() time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

/*
	TodayStartUnix 获取今天零点零分零秒的时间
*/
func TodayStartUnix() int64 {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
}

/*
	YesterdayStartUnix 获取昨天零点零分零秒的时间戳
*/
func YesterdayStartUnix() int64 {
	t := time.Now()
	t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return t1.Add(-24 * time.Hour).Unix()
}

/*
	TomorrowStartTime 获取明天零点零分零秒的时间
*/
func TomorrowStartTime() time.Time {
	t := time.Now()
	t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return t1.Add(24 * time.Hour)
}
