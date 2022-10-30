package replyrouter

import (
	"fmt"
	"line-bot-private/anju"
	"line-bot-private/anjuworkdata"
	"line-bot-private/dateutil"
	"regexp"
	"strconv"
	"time"
)

func pointerWith(v string) *string {
	return &v
}

func Routing(text string) *string {
	today := dateutil.Today()
	workSchedule := anjuworkdata.WorkSchedule2022_9to10
	if text == "今の時間" {
		timeStr := fmt.Sprint(time.Now().In(dateutil.TimeZoneJST))
		return pointerWith(timeStr)
	} else if text == "今日のシフト" {
		todaysWorkday := workSchedule.TimeWithSelectDate(today)
		if todaysWorkday != nil {
			return pointerWith(messageOfTodayWorking(todaysWorkday))
		}
		workDays := workSchedule.TimesWithSelectAfterDate(today)
		if len(workDays) > 0 {
			latestWorkDay := workDays[0]
			return pointerWith(messageOfAfterDayWorking(latestWorkDay))
		}
		return pointerWith("本日はお休みのようです。\nゆっくりとお休みくださいませ")
	}
	selectDayFormatPattern := "(\\d{1,2})月(\\d{1,2})日のシフト"
	selectDayRegex := regexp.MustCompile(selectDayFormatPattern)
	result := selectDayRegex.FindAllStringSubmatch(text, -1)
	if len(result) != 0 {
		month, err := strconv.Atoi(result[0][1])
		if err != nil {
			return nil
		}
		day, err := strconv.Atoi(result[0][2])
		if err != nil {
			return nil
		}
		year := today.Year()
		if month == 1 {
			year += 1
		}
		selectDate := dateutil.DateFromYearMonthDay(year, month, day)
		selectWorkDay := workSchedule.TimeWithSelectDate(selectDate)
		if selectWorkDay == nil {
			return pointerWith("その日はお休みのようです。")
		}
		return pointerWith(messageOfSelectWorkingDate(selectWorkDay))
	}
	return nil
}

func messageOfTodayWorking(workDay *anju.WorkDay) string {
	item := workDay.Items[0]
	startLabel := dateutil.Format24Hour(item.StartTime)
	endLabel := dateutil.Format24Hour(item.EndTime)
	contentMessage := fmt.Sprintf("本日は%s〜%sまでシフトが入っているようです。", startLabel, endLabel)
	if len(workDay.Items) >= 2 {
		otherItems := workDay.Items[1:]
		contentMessage += "\n\n他のメンバーのシフトはこのようになっております。"
		for _, otherItem := range otherItems {
			startLabel = dateutil.Format24Hour(otherItem.StartTime)
			endLabel = dateutil.Format24Hour(otherItem.EndTime)
			otherMemberContent := fmt.Sprintf("\n%s\n%s〜%sまで", otherItem.Member.Name, startLabel, endLabel)
			contentMessage += otherMemberContent
		}
		contentMessage += "\n以上になります。\n"
	}
	contentMessage += "\n本日もがんばっていきましょう"
	return contentMessage
}

func messageOfAfterDayWorking(workDay *anju.WorkDay) string {
	item := workDay.Items[0]
	afterWorkDayLabel := fmt.Sprintf("%d月%d日", int(workDay.Date.Month()), workDay.Date.Day())
	startLabel := dateutil.Format24Hour(item.StartTime)
	endLabel := dateutil.Format24Hour(item.EndTime)
	contentMessage := "本日はお休みのようです。"
	contentMessage += fmt.Sprintf("直近ですと%sに%s〜%sまでシフトが入っているようです。", afterWorkDayLabel, startLabel, endLabel)
	if len(workDay.Items) >= 2 {
		otherItems := workDay.Items[1:]
		contentMessage += "\n\n他のメンバーのシフトはこのようになっております。"
		for _, otherItem := range otherItems {
			startLabel = dateutil.Format24Hour(otherItem.StartTime)
			endLabel = dateutil.Format24Hour(otherItem.EndTime)
			otherMemberContent := fmt.Sprintf("\n%s\n%s〜%sまで", otherItem.Member.Name, startLabel, endLabel)
			contentMessage += otherMemberContent
		}
		contentMessage += "\n以上になります。\n"
	}
	contentMessage += "\nゆっくりとお休みくださいませ"
	return contentMessage
}

func messageOfSelectWorkingDate(workDay *anju.WorkDay) string {
	item := workDay.Items[0]
	startLabel := dateutil.Format24Hour(item.StartTime)
	endLabel := dateutil.Format24Hour(item.EndTime)
	contentMessage := fmt.Sprintf("その日は%s〜%sまでシフトが入っているようです。", startLabel, endLabel)
	if len(workDay.Items) >= 2 {
		otherItems := workDay.Items[1:]
		contentMessage += "\n\n他のメンバーのシフトはこのようになっております。"
		for _, otherItem := range otherItems {
			startLabel = dateutil.Format24Hour(otherItem.StartTime)
			endLabel = dateutil.Format24Hour(otherItem.EndTime)
			otherMemberContent := fmt.Sprintf("\n%s\n%s〜%sまで", otherItem.Member.Name, startLabel, endLabel)
			contentMessage += otherMemberContent
		}
		contentMessage += "\n以上になります。\n"
	}
	return contentMessage
}
