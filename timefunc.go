package main

import (
	"time"
)

// PeriodOfUse 함수는 RFC3339 시간문자열을 받아서 사용기간을 반환한다.
func PeriodOfUse(timestr string) (int, error) {
	past, err := time.Parse(time.RFC3339, timestr)
	if err != nil {
		return 0, err
	}
	now := time.Now()
	total := 0
	if now.Year()-past.Year() > 1 {
		// 구매일이 2년이상 이면
		total += (now.Year() - past.Year()) * 12
	} else if (now.Year() - past.Year()) == 1 {
		// 작년에 산것
		total += 12 - int(past.Month())
	}
	if now.Year() == past.Year() {
		// 올해 산것
		total += int(now.Month() - past.Month())
	} else {
		total += int(now.Month())
	}
	return total, nil
}
