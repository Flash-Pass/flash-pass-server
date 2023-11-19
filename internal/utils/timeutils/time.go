package timeutils

func GetStartTimestampOfDay(timestamp int64) int64 {
	moreThanOneDay := timestamp % 86400
	return timestamp - moreThanOneDay
}

func GetEndTimestampOfDay(timestamp int64) int64 {
	moreThanOneDay := timestamp % 86400
	return timestamp + (86400 - moreThanOneDay)
}

func TimestampLag(timestamp int64, lag int64) int64 {
	return timestamp - lag
}
