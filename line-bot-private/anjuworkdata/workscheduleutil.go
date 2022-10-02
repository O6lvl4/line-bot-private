package anjuworkdata

import (
	"line-bot-private/anju"
	"log"
	"time"
)

func newWorkDate(monthNumber int, day int) time.Time {
	return time.Date(2022, newMonth(monthNumber), day, 0, 0, 0, 0, time.Local)
}

func newWorkTime(date time.Time, hour int, minutes int) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), hour, minutes, 0, 0, time.Local)
}

type workDayItem struct {
	member           *anju.WorkMember
	startWorkHour    int
	startWorkMinutes int
	endWorkHour      int
	endWorkMinutes   int
}

func newWorkDay(date time.Time, items []*workDayItem) *anju.WorkDay {
	var replaceItems []*anju.WorkDayItem
	for _, item := range items {
		replaceItems = append(replaceItems, &anju.WorkDayItem{
			Member:    item.member,
			StartTime: newWorkTime(date, item.startWorkHour, item.startWorkMinutes),
			EndTime:   newWorkTime(date, item.endWorkHour, item.endWorkMinutes),
		})
	}
	return &anju.WorkDay{
		Date:  date,
		Items: replaceItems,
	}
}

func newMonth(monthNumber int) time.Month {
	switch monthNumber {
	case 1:
		return time.January
	case 2:
		return time.February
	case 3:
		return time.March
	case 4:
		return time.April
	case 5:
		return time.May
	case 6:
		return time.June
	case 7:
		return time.July
	case 8:
		return time.August
	case 9:
		return time.September
	case 10:
		return time.October
	case 11:
		return time.November
	case 12:
		return time.December
	default:
		log.Fatal("Invalid month number value: ", monthNumber)
		return -1
	}
}
