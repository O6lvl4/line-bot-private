package anjuworkdata

import (
	"line-bot-private/anju"
	"line-bot-private/dateutil"
)

var WorkSchedule2022_9to10 *anju.WorkSchedule = &anju.WorkSchedule{
	StartDate: dateutil.DateFromYearMonthDay(2022, 10, 15),
	EndDate:   dateutil.DateFromYearMonthDay(2022, 11, 14),
	Days: []*anju.WorkDay{
		newWorkDay(dateutil.DateFromYearMonthDay(2022, 11, 2), []*workDayItem{
			{
				member:           WorkMemberあんじゅ,
				startWorkHour:    14,
				startWorkMinutes: 0,
				endWorkHour:      21,
				endWorkMinutes:   00,
			},
			{
				member:           WorkMemberふくもとさん,
				startWorkHour:    10,
				startWorkMinutes: 00,
				endWorkHour:      18,
				endWorkMinutes:   00,
			},
			{
				member:           WorkMemberこぞのさん,
				startWorkHour:    13,
				startWorkMinutes: 10,
				endWorkHour:      21,
				endWorkMinutes:   10,
			},
			{
				member:           WorkMemberしまださん,
				startWorkHour:    10,
				startWorkMinutes: 00,
				endWorkHour:      18,
				endWorkMinutes:   00,
			},
			{
				member:           WorkMemberあきたさん,
				startWorkHour:    9,
				startWorkMinutes: 40,
				endWorkHour:      17,
				endWorkMinutes:   40,
			},
			{
				member:           WorkMemberなかむらさん,
				startWorkHour:    13,
				startWorkMinutes: 0,
				endWorkHour:      21,
				endWorkMinutes:   00,
			},
			{
				member:           WorkMemberきたいちさん,
				startWorkHour:    9,
				startWorkMinutes: 40,
				endWorkHour:      17,
				endWorkMinutes:   40,
			},
		}),
		newWorkDay(dateutil.DateFromYearMonthDay(2022, 11, 5), []*workDayItem{
			{
				member:           WorkMemberあんじゅ,
				startWorkHour:    14,
				startWorkMinutes: 0,
				endWorkHour:      21,
				endWorkMinutes:   00,
			},
		}),
		newWorkDay(dateutil.DateFromYearMonthDay(2022, 11, 7), []*workDayItem{
			{
				member:           WorkMemberあんじゅ,
				startWorkHour:    8,
				startWorkMinutes: 0,
				endWorkHour:      15,
				endWorkMinutes:   00,
			},
		}),
		newWorkDay(dateutil.DateFromYearMonthDay(2022, 11, 8), []*workDayItem{
			{
				member:           WorkMemberあんじゅ,
				startWorkHour:    9,
				startWorkMinutes: 40,
				endWorkHour:      16,
				endWorkMinutes:   40,
			},
		}),
		newWorkDay(dateutil.DateFromYearMonthDay(2022, 11, 9), []*workDayItem{
			{
				member:           WorkMemberあんじゅ,
				startWorkHour:    9,
				startWorkMinutes: 40,
				endWorkHour:      16,
				endWorkMinutes:   40,
			},
		}),
		newWorkDay(dateutil.DateFromYearMonthDay(2022, 11, 12), []*workDayItem{
			{
				member:           WorkMemberあんじゅ,
				startWorkHour:    14,
				startWorkMinutes: 10,
				endWorkHour:      21,
				endWorkMinutes:   10,
			},
		}),
		newWorkDay(dateutil.DateFromYearMonthDay(2022, 11, 13), []*workDayItem{
			{
				member:           WorkMemberあんじゅ,
				startWorkHour:    14,
				startWorkMinutes: 0,
				endWorkHour:      21,
				endWorkMinutes:   0,
			},
		}),
		newWorkDay(dateutil.DateFromYearMonthDay(2022, 11, 16), []*workDayItem{
			{
				member:           WorkMemberあんじゅ,
				startWorkHour:    10,
				startWorkMinutes: 0,
				endWorkHour:      17,
				endWorkMinutes:   0,
			},
		}),
		newWorkDay(dateutil.DateFromYearMonthDay(2022, 11, 18), []*workDayItem{
			{
				member:           WorkMemberあんじゅ,
				startWorkHour:    10,
				startWorkMinutes: 0,
				endWorkHour:      17,
				endWorkMinutes:   0,
			},
		}),
		newWorkDay(dateutil.DateFromYearMonthDay(2022, 11, 20), []*workDayItem{
			{
				member:           WorkMemberあんじゅ,
				startWorkHour:    14,
				startWorkMinutes: 10,
				endWorkHour:      21,
				endWorkMinutes:   10,
			},
		}),
	},
}
