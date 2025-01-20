# IranWeather_Bot ☂️

**IranWeather_Bot** is a Telegram bot built with Golang that provides weather forecasts for cities in Iran. It fetches real-time weather data using the OpenWeatherMap API.

## Features ✨
- Provides weather forecasts for 349 cities across Iran.
- Users can:
  - Select cities from a predefined list at the start.
  - Type a city's name to get a forecast.
- Delivers weather forecasts for three hours ahead, including:
  - Temperature
  - Humidity percentage
  - Wind speed
  - Weather description (e.g., partly cloudy) in English.
- Ready for upgrade to support 5-day forecasts.

## How It Works ⚙️
1. At the start, the bot offers a menu with a few cities to choose from.
2. Users can also type the name of a city directly.
3. The bot uses pre-stored latitude and longitude coordinates for 349 cities to fetch accurate data.

## Requirements 🔧
- Golang installed on your system.
- An OpenWeatherMap API key.

## Installation 🛠️
1. Clone the repository:
   ```bash
   git clone https://github.com/SaeedSH02/IranWeather_Bot.git
   cd IranWeather_Bot
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Add your OpenWeatherMap API key to the configuration file (e.g., `config.go` or an environment variable).
4. Run the bot:
   ```bash
   go run main.go
   ```

## Commands 🖬️
- `/start`: Display a list of cities to choose from.
- Type a city's name: Get the three-hour weather forecast for the city.

## Data Details 🌧️
- Preloaded coordinates for 349 cities in Iran are included in the bot.
- Retrieves:
  - Temperature
  - Humidity percentage
  - Wind speed
  - Weather description in English.

## Future Enhancements 🌟
- Expand forecast capabilities to provide 5-day weather predictions.
- Add support for additional commands or city search features.

## Contributions 📚
Contributions are welcome! Feel free to fork the repo, make your changes, and submit a pull request.

## License 🔒
This project is licensed under the [MIT License](LICENSE).


