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
		startLabel := fmt.Sprintf("%02d:%02d", workItem.StartTime.Hour(), workItem.StartTime.Minute())
		endLabel := fmt.Sprintf("%02d:%02d", workItem.EndTime.Hour(), workItem.EndTime.Minute())
		contentMessage := fmt.Sprintf("本日は%s〜%sまでシフトが入っているようです。", startLabel, endLabel)
		if len(workDay.Items) >= 2 {
			otherItems := workDay.Items[1:]
			contentMessage += "\n\n他のメンバーのシフトはこのようになっております。"
			for _, otherItem := range otherItems {
				startLabel = fmt.Sprintf("%02d:%02d", otherItem.StartTime.Hour(), workItem.StartTime.Minute())
				endLabel = fmt.Sprintf("%02d:%02d", otherItem.EndTime.Hour(), workItem.EndTime.Minute())
				otherMemberContent := fmt.Sprintf("\n%s\n%s〜%sまで", otherItem.Member.Name, startLabel, endLabel)
				contentMessage += otherMemberContent
			}
			contentMessage += "\n以上になります。\n"
		}
		contentMessage += "\n本日もがんばっていきましょう"
		return pointerWith(contentMessage)
	} else {
		return nil
	}
}
