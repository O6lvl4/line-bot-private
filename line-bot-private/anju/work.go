package anju

import "time"

type WorkSchedule struct {
	StartDate time.Time
	EndDate   time.Time
	Days      []*WorkDay
}

func (w *WorkSchedule) TimeWithSelectDate(date time.Time) *WorkDay {
	equal_yyyyMMdd := func(lhs time.Time, rhs time.Time) bool {
		return lhs.Year() == rhs.Year() &&
			lhs.Month() == rhs.Month() &&
			lhs.Day() == rhs.Day()
	}
	for _, workDay := range w.Days {
		if equal_yyyyMMdd(date, workDay.Date) {
			return workDay
		}
	}
	return nil
}

func (w *WorkSchedule) TimesWithSelectAfterDate(date time.Time) []*WorkDay {
	equal_yyyyMMdd := func(lhs time.Time, rhs time.Time) bool {
		return lhs.Year() == rhs.Year() &&
			lhs.Month() == rhs.Month() &&
			lhs.Day() == rhs.Day()
	}
	var days = make([]*WorkDay, 0)
	for _, workDay := range w.Days {
		if equal_yyyyMMdd(date, workDay.Date) || date.Before(workDay.Date) {
			days = append(days, workDay)
		}
	}
	return days
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
