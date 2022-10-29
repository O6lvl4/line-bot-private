package anjuworkdata

import (
	"line-bot-private/anju"
	"line-bot-private/dateutil"
	"time"
)

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
			Member: item.member,
			StartTime: dateutil.DateFromDateWithHourMinutes(
				date, item.startWorkHour, item.startWorkMinutes),
			EndTime: dateutil.DateFromDateWithHourMinutes(
				date, item.endWorkHour, item.endWorkMinutes),
		})
	}
	return &anju.WorkDay{
		Date:  date,
		Items: replaceItems,
	}
}
