package humanTime

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func FormatDate(input string) (string, error) {
	// Parse the input string as UTC time
	t, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return "", err
	}

	// Extract the day and determine the correct suffix (st, nd, rd, th)
	day := t.Day()
	suffix := "th"
	if day%10 == 1 && day != 11 {
		suffix = "st"
	} else if day%10 == 2 && day != 12 {
		suffix = "nd"
	} else if day%10 == 3 && day != 13 {
		suffix = "rd"
	}

	// Format time to AM/PM
	formattedDate := fmt.Sprintf("%d %s %d%s %02d:%02d:%02d %s",
		t.Year(),
		t.Month().String(), // Full month name
		day, suffix,
		t.Hour()%12, t.Minute(), t.Second(),
		strings.ToUpper(t.Format("PM")), // Convert to uppercase
	)

	return formattedDate, nil
}

func DiffSecFromNow(t time.Time) float64 {
	currentTime := time.Now().Local()
	loc, _ := time.LoadLocation("Asia/Jakarta")
	oldTime := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, loc)
	fmt.Println(currentTime)
	fmt.Println(oldTime)
	fmt.Println(oldTime.Sub(currentTime))
	return oldTime.Sub(currentTime).Seconds()
}

func StringTimestampToTime(timeString string) time.Time {
	timeString = strings.ReplaceAll(timeString, "T", " ")
	timeString = strings.ReplaceAll(timeString, "Z", "")
	fmt.Println(timeString)

	year, _ := strconv.Atoi(timeString[:4])
	month, _ := strconv.Atoi(timeString[5:7])
	day, _ := strconv.Atoi(timeString[8:10])
	hour, _ := strconv.Atoi(timeString[11:13])
	minute, _ := strconv.Atoi(timeString[14:16])
	sec, _ := strconv.Atoi(timeString[17:19])

	loc, _ := time.LoadLocation("Asia/Jakarta")
	timestamp := time.Date(year, time.Month(month), day, hour, minute, sec, 0, loc)

	return timestamp
}
