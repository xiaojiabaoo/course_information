package pkg

import (
	"strings"
	"time"
)

var (
	toBeCharge   = "2006-01-02T15:04:05+08:00"
	toBeChargeUT = "2006-01-02T15:04:05+00:00"
	toBeCharge2  = "2006-01-02 15:04:05"
	timeLayout   = "2006-01-02"
	timeLayoutDt = "2006-01-02T15:04:05+08:00"
	loc, _       = time.LoadLocation("Asia/Shanghai") //设置时区
)

//  yyyy-mm-dd  hh:mm:ss 格式的时间转换为TTL时间
func StringTimeToInt(str string) int {
	te, err := time.ParseInLocation(toBeCharge2, str, loc)
	if err != nil {
		return 0
	}
	return int(te.Unix())
}

//  yyyy-mm-dd  hh:mm:ss 格式的时间转换为TTL时间
func StringTimeToInt64(str string) int64 {
	te, err := time.ParseInLocation(toBeCharge2, str, loc)
	if err != nil {
		return 0
	}
	return te.Unix()
}

//TTL时间转Time YYYY-MM--HH
func TTLToYMDString(t int64) string {
	// go语言固定日期模版
	// time.Unix的第二个参数传递0或10结果一样，因为都不大于1e9
	timeStr := time.Unix(t, 0).Format(timeLayout)
	return timeStr
}

//TTL时间转Time YYYY-MM--HH MM:SS
func TTLToDateTime(t int64) time.Time {
	// go语言固定日期模版
	timeLayout := "2006-01-02 15:04:05"
	// time.Unix的第二个参数传递0或10结果一样，因为都不大于1e9
	timeStr := time.Unix(t, 0).Format(timeLayout)
	st, _ := time.Parse(timeLayout, timeStr) //string转time
	return st
}

func GetTimeNow() time.Time {
	return time.Now()
}

//TTL时间转Time YYYY-MM--HH MM:SS
func TTLToDateTimeString(t int64) string {
	if t <= 0 {
		return ""
	}
	// go语言固定日期模版
	timeLayout := "2006-01-02 15:04:05"
	// time.Unix的第二个参数传递0或10结果一样，因为都不大于1e9
	timeStr := time.Unix(t, 0).Format(timeLayout)
	return timeStr
}

//string转上海时间Time
func StringToTimeYMD(t string) time.Time {
	timeLayout := "2006-01-02"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	ymd, _ := time.ParseInLocation(timeLayout, t, loc)
	return ymd
}

//string转上海时间Time
func StringToTimeYMDHMS(t string) time.Time {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	ymd, _ := time.ParseInLocation(timeLayout, t, loc)
	return ymd
}

func TimeParse(str string) (*time.Time, error) {
	t, err := time.Parse(toBeCharge2, str)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func DateParse(str string) (*time.Time, error) {
	t, err := time.Parse(timeLayout, str)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// int 转time.Data
func Int2String(times int64) string {
	tm := time.Unix(times, 0)
	return tm.Format("2006-01-02 15:04:05")
}

// int 转time.Time
func Int2DataTime(times int64) time.Time {
	tm := time.Unix(times, 0)
	return tm
}

//获取范围为北京时间8:00至次日7:59,根据当前时间判断
func GetDayTime() (int, int) {
	Date := time.Now().Format("2006-01-02")
	tst := Date + " 00:00:00"
	_, todayStartTime := TimeStrToTimestamp(tst)

	tet := Date + " 23:59:59"
	_, todayEndTime := TimeStrToTimestamp(tet)

	return todayStartTime, todayEndTime
}

//获取范围为北京时间8:00至次日7:59,根据当前时间判断
func GetDateTime(date string) (int, int, bool) {
	tst := date + " 00:00:00"
	err, startTTL := TimeStrToTimestamp(tst)
	if err != nil {
		return 0, 0, false
	}
	tet := date + " 23:59:59"
	err, endTTL := TimeStrToTimestamp(tet)
	if err != nil {
		return 0, 0, false
	}
	return startTTL, endTTL, true
}

//获取当天日期
func GetNowDate() string {
	return time.Now().Format("2006-01-02")
}

//时间字符串转时间戳 入参格式:2006-01-02 15:04:05  返回值：时间戳
func TimeStrToTimestamp(timeStr string) (error, int) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return err, 0
	}
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	if err != nil {
		return err, 0
	}
	sr := int(theTime.Unix())
	return nil, sr
}

//true   2021-9-1  ->  2021-09-01
//false  2021-09-01 -> 2021-9-1
func DateChange(date string, b bool) string {
	data := ""
	strs := strings.Split(date, "-")
	if len(strs) != 3 {
		return data
	}
	data += strs[0]
	if b {
		if len(strs[1]) == 1 {
			data += "-0" + strs[1]
		} else {
			data += "-" + strs[1]
		}

		if len(strs[2]) == 1 {
			data += "-0" + strs[2]
		} else {
			data += "-" + strs[2]
		}
	} else {
		if strings.HasPrefix(strs[1], "0") {
			data += "-" + strs[1][1:]
		} else {
			data += "-" + strs[1]
		}

		if strings.HasPrefix(strs[2], "0") {
			data += "-" + strs[2][1:]
		} else {
			data += "-" + strs[2]
		}
	}
	return data
}