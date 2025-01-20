package handlers

import (
	"Weatherbot/weather"
	"fmt"
	tele "gopkg.in/telebot.v4"
)


func Handlers(bot *tele.Bot ){

	bot.Handle("/start", func(ctx tele.Context) error {
		button := &tele.ReplyMarkup{}
		Tehran := button.Text("تهران")
		Isfahan := button.Text("اصفهان")
		Shiraz := button.Text("شیراز")
		Rasht := button.Text("رشت")
		Mashhad := button.Text("مشهد")
		Tabriz := button.Text("تبریز")
		Zabol := button.Text("زابل")
		Ahwaz := button.Text("اهواز")
		More_BTN := button.Text("...بیشتر")
	
		button.Reply(
			button.Row(Tehran,Rasht),
			button.Row(Isfahan, Shiraz,Mashhad),
			button.Row(Tabriz, Zabol,Ahwaz),
			button.Row(More_BTN),
		)
		return ctx.Send("شهر مورد نظر خود را انتخاب کنید:", button)
	})

	bot.Handle(tele.OnText, func(ctx tele.Context) error {
		text :=ctx.Text()
		cities := weather.GetCities()
		if text == "...بیشتر" {
			return ctx.Send("اسم شهر موردنظر خود را بنویسید مانند: \n کرج")
		}
		for k, v := range cities {
			if k == text {
				lon := v.Lon
				lat := v.Lat
				weatherData, err := weather.GetWeather(lat, lon)
				if err != nil {
					return ctx.Send("try later...")
				}
				return ctx.Send(fmt.Sprintln("Weather in "+ k,"\n\n", weatherData))
			}
		}
		return ctx.Send("شهر مورد نظر یافت نشد. لطفاً نام شهر دیگری وارد کنید.")
	})

	bot.Handle("/help", func(ctx tele.Context) error {
		return ctx.Send("یک شهر از منو زیر انتخاب کنید یا از اسم شهر مورد نظر را ارسال کنید مانند: \n سمنان")
	})


}