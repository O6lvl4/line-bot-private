package anjuworkdata

import (
	"line-bot-private/anju"
)

var WorkSchedule2022_9to10 *anju.WorkSchedule = &anju.WorkSchedule{
	StartDate: newWorkDate(9, 15),
	EndDate:   newWorkDate(10, 14),
	Days: []*anju.WorkDay{
		newWorkDay(newWorkDate(10, 2), []*workDayItem{
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
		newWorkDay(newWorkDate(10, 5), []*workDayItem{
			{
				member:           WorkMemberあんじゅ,
				startWorkHour:    14,
				startWorkMinutes: 0,
				endWorkHour:      21,
				endWorkMinutes:   00,
			},
		}),
		newWorkDay(newWorkDate(10, 7), []*workDayItem{
			{
				member:           WorkMemberあんじゅ,
				startWorkHour:    8,
				startWorkMinutes: 0,
				endWorkHour:      15,
				endWorkMinutes:   00,
			},
		}),
		newWorkDay(newWorkDate(10, 8), []*workDayItem{
			{
				member:           WorkMemberあんじゅ,
				startWorkHour:    9,
				startWorkMinutes: 40,
				endWorkHour:      16,
				endWorkMinutes:   40,
			},
		}),
		newWorkDay(newWorkDate(10, 9), []*workDayItem{
			{
				member:           WorkMemberあんじゅ,
				startWorkHour:    9,
				startWorkMinutes: 40,
				endWorkHour:      16,
				endWorkMinutes:   40,
			},
		}),
		newWorkDay(newWorkDate(10, 12), []*workDayItem{
			{
				member:           WorkMemberあんじゅ,
				startWorkHour:    14,
				startWorkMinutes: 10,
				endWorkHour:      21,
				endWorkMinutes:   10,
			},
		}),
		newWorkDay(newWorkDate(10, 13), []*workDayItem{
			{
				member:           WorkMemberあんじゅ,
				startWorkHour:    14,
				startWorkMinutes: 0,
				endWorkHour:      21,
				endWorkMinutes:   0,
			},
		}),
		newWorkDay(newWorkDate(10, 16), []*workDayItem{
			{
				member:           WorkMemberあんじゅ,
				startWorkHour:    10,
				startWorkMinutes: 0,
				endWorkHour:      17,
				endWorkMinutes:   0,
			},
		}),
		newWorkDay(newWorkDate(10, 18), []*workDayItem{
			{
				member:           WorkMemberあんじゅ,
				startWorkHour:    10,
				startWorkMinutes: 0,
				endWorkHour:      17,
				endWorkMinutes:   0,
			},
		}),
		newWorkDay(newWorkDate(10, 20), []*workDayItem{
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
