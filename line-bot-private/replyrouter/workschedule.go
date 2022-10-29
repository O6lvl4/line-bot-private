package replyrouter

import (
	"line-bot-private/anju"
	"line-bot-private/anjuworkdata"
	"strings"
)

const (
	todayWorkScheduleMessageFormat     = "本日は%s〜%sまでシフトが入っているようです。"
	otherMemberScheduleAnnounceMessage = "他のメンバーのシフトはこのようになっております。"
)

type WorkScheduleRoute struct {
	workSchedule *anju.WorkSchedule
}

func (w *WorkScheduleRoute) IsHandle(text string) bool {
	return strings.Contains(text, "シフト")
}

func NewWorkScheduleRoute() *WorkScheduleRoute {
	return &WorkScheduleRoute{
		workSchedule: anjuworkdata.WorkSchedule2022_9to10,
	}
}

// Q. シフトを聞く
// If. 今日シフトがあるなら、そのシフトを表示する
// Else. 今日シフトがないなら、一番直近のシフトを表示する
