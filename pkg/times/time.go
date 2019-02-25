package times

import (
	"time"
)

func StrToIntTime(str,level string) int{
	timeLayout  := "2006-01-02 15:04:05"
	switch level {
		case "y":
			timeLayout  = "2006"
		case "m":
			timeLayout  = "2006-01"
		case "d":
			timeLayout  = "2006-01-02"
		case "h":
			timeLayout  = "2006-01-02 15"
		case "i":
			timeLayout  = "2006-01-02 15:04"
		case "s":
			timeLayout  = "2006-01-02 15:04:05"
		case "ms":
			timeLayout  = "2006-01-02 15:04:05"
	}
	local,_ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout,str,local)
	return int(theTime.Unix())
}
