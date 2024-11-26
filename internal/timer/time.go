package timer

import "time"

func GetNowTime() time.Time {
	return time.Now()
}

func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, error := time.ParseDuration(d)
	if error != nil {
		return time.Time{}, error
	}

	return currentTimer.Add(duration), nil
}
