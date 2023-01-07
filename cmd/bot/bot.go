package bot

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"log"
	"math/rand"
	"time"
)

type DimxxBot struct {
	Bot *telebot.Bot
}

func InitBot(token string) *telebot.Bot {
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)

	if err != nil {
		log.Fatalf("Ошибка при инициализации бота %v", err)
	}

	return b
}

func (bot *DimxxBot) StartHandler(ctx telebot.Context) error {
	return ctx.Send("Привет," + ctx.Sender().FirstName)
}

func (bot *DimxxBot) SemesterHandle(ctx telebot.Context) error {
	result := ""
	startDate := time.Date(2023, time.January, 23, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 0, 15*7)

	now := time.Now()

	// Рассчитываем разницу между датами в неделях
	weeksSinceStart := int(now.Sub(startDate).Hours() / (24 * 7))

	// Рассчитываем текущую неделю в семестре
	currentWeek := weeksSinceStart % 15

	if currentWeek < 1 {
		result = fmt.Sprintf("Семестр начинается: %s,\nСеместр заканчивается: %s,\nТекущая неделя: Отдых", startDate.Format("02.01.2006"), endDate.Format("02.01.2006"))
		return ctx.Send(result)
	} else {
		result = fmt.Sprintf("Семестр начинается: %s,\nСеместр заканчивается: %s,\nТекущая неделя: %v", startDate.Format("02.01.2006"), endDate.Format("02.01.2006"), currentWeek+1)
	}
	return ctx.Send(result)
}

func (bot *DimxxBot) AskHandler(ctx telebot.Context) error {
	msg := connectAI(ctx.Text())
	return ctx.Send(msg)
}

func (bot *DimxxBot) RandomChoiceHandler(ctx telebot.Context) error {
	items := []string{"да", "нет"}
	randIndex := rand.Intn(len(items))
	randElement := items[randIndex]
	return ctx.Send("Результат : " + randElement)
}

func (bot *DimxxBot) RandomMotivationText(ctx telebot.Context) error {
	mtvntn := connectAI("дай 1 мотивационную фразу")
	return ctx.Send(mtvntn)
}
