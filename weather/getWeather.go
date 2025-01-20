package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/joho/godotenv"
)

func GetWeather(lat, lon float64) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", fmt.Errorf("error loading .env file: %v", err)
	}
	appid := os.Getenv("openWeatherAPI")

	
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&appid=%s&units=metric", lat, lon, appid)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("failed to get weather data:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("failed to get weather data:", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("cant Read json: ", err)
	}

	var weathers WeatherResponse
	if err:= json.Unmarshal(body, &weathers); err != nil {
		log.Fatal(err)
	}
	if len(weathers.List) == 0 {
		return fmt.Sprintln("No weather data available"), nil
	}
	temp := weathers.List[0].Main.Temp
	humidity := weathers.List[0].Main.Humidity
	wind := weathers.List[0].Wind.Speed
	description := weathers.List[0].Weather[0].Description
	// main := weathers.List[0].Weather[0].Main
	unixTime := int64(weathers.List[0].Dt)
	utcTime := time.Unix(unixTime, 0)
	utcFormat := utcTime.Format("2006-01-02 15:04:05")
	
	
	return fmt.Sprintf("🌡️ temperature: %.1f°C\n\n 🌤️ Description: %s \n\n 💧 Humidity: %v %%\n\n 💨 Wind: %vkm/h \n\n 🕓 Time: %s ",
	 temp, description,  humidity, wind, utcFormat), nil
}