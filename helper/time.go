package helper

import "time"

func DateTime(t time.Time, layout string) (time.Time, error) {

	s := t.Local().Format(layout)
	return time.Parse(layout, s)
}

// string to unix 2006-01-02 =>
func String2Unix(timeStr, layout string, local *time.Location) (int64, error) {
	if local == nil {
		local = time.Local
	}
	location, err := time.ParseInLocation(layout, timeStr, local)
	if err != nil {
		return -1, err
	}
	return location.Unix(), nil
}
