package anju

import "time"

type WorkSchedule struct {
	StartDate time.Time
	EndDate   time.Time
	Days      []*WorkDay
}

func (w *WorkSchedule) WorkDayWithSelectDate(date time.Time) *WorkDay {
	compareDay := func(lhs time.Time, rhs time.Time) bool {
		return lhs.Year() == rhs.Year() &&
			lhs.Month() == rhs.Month() &&
			lhs.Day() == rhs.Day()
	}
	for _, workDay := range w.Days {
		if compareDay(date, workDay.Date) {
			return workDay
		}
	}
	return nil
}

type WorkDay struct {
	Date  time.Time
	Items []*WorkDayItem
}

type WorkDayItem struct {
	Member    *WorkMember
	StartTime time.Time
	EndTime   time.Time
}

type WorkMember struct {
	Name string
}
