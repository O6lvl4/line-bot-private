package replyrouter

import (
	"fmt"
	"line-bot-private/anjuworkdata"
	"time"
)

func pointerWith(v string) *string {
	return &v
}

func Routing(text string) *string {
	if text == "シフト" {
		today := time.Now()
		workSchedule := anjuworkdata.WorkSchedule2022_9to10
		workDay := workSchedule.WorkDayWithSelectDate(today)
		if workDay == nil {
			return pointerWith("本日はお休みのようです")
		}
		workItem := workDay.Items[0]
		startLabel := fmt.Sprint(workItem.StartTime.Hour()) + "時" + fmt.Sprint(workItem.StartTime.Minute()) + "分"
		endLabel := fmt.Sprint(workItem.EndTime.Hour()) + "時" + fmt.Sprint(workItem.EndTime.Minute()) + "分"
		return pointerWith("本日は" + startLabel + "から" + endLabel + "までシフトが入っているようです")
	} else {
		return nil
	}
}
