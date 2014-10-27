package openweathermap

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// IconData holds the relevant info for linking icons to conditions.
type IconData struct {
	Condition string
	Day       string
	Night     string
}

// ConditionData holds data structure for weather conditions information.
type ConditionData struct {
	ID      int
	Meaning string
	Icon1   string
	Icon2   string
}

// RetrieveIcon will get the specified icon from the API.
func RetrieveIcon(destination, iconFile string) (int64, error) {
	fullFilePath := fmt.Sprintf("%s/%s", destination, iconFile)

	if _, err := os.Stat(fullFilePath); err != nil {
		out, createErr := os.Create(fullFilePath)
		if createErr != nil {
			log.Fatalln(createErr)
		}
		defer out.Close()

		response, getErr := http.Get(fmt.Sprintf(iconUrl, iconFile))
		if getErr != nil {
			log.Fatalln(getErr)
		}
		defer response.Body.Close()

		n, copyErr := io.Copy(out, response.Body)
		if copyErr != nil {
			log.Fatalln(copyErr)
		}
		return n, nil
	}
	return 0, errors.New("File exists.  Using found file")
}

var IconList = []*IconData{
	&IconData{Condition: "clear sky", Day: "01d.png", Night: "01n.png"},
	&IconData{Condition: "few clouds", Day: "02d.png", Night: "02n.png"},
	&IconData{Condition: "scattered clouds", Day: "03d.png", Night: "03n.png"},
	&IconData{Condition: "broken clouds", Day: "04d.png", Night: "04n.png"},
	&IconData{Condition: "shower rain", Day: "09d.png", Night: "09n.png"},
	&IconData{Condition: "rain", Day: "10d.png", Night: "10n.png"},
	&IconData{Condition: "thunderstorm", Day: "11d.png", Night: "11n.png"},
	&IconData{Condition: "snow", Day: "13d.png", Night: "13n.png"},
	&IconData{Condition: "mist", Day: "50d.png", Night: "50n.png"},
}

var ThunderstormConditions = []*ConditionData{
	&ConditionData{ID: 200, Meaning: "thunderstorm with light rain", Icon1: "11d.png"},
	&ConditionData{ID: 201, Meaning: "thunderstorm with rain", Icon1: "11d.png"},
	&ConditionData{ID: 202, Meaning: "thunderstorm with heavy rain", Icon1: "11d.png"},
	&ConditionData{ID: 210, Meaning: "light thunderstorm", Icon1: "11d.png"},
	&ConditionData{ID: 211, Meaning: "thunderstorm", Icon1: "11d.png"},
	&ConditionData{ID: 212, Meaning: "heavy thunderstorm", Icon1: "11d.png"},
	&ConditionData{ID: 221, Meaning: "ragged thunderstorm", Icon1: "11d.png"},
	&ConditionData{ID: 230, Meaning: "thunderstorm with light drizzle", Icon1: "11d.png"},
	&ConditionData{ID: 231, Meaning: "thunderstorm with drizzle", Icon1: "11d.png"},
	&ConditionData{ID: 232, Meaning: "thunderstorm with heavy drizzle", Icon1: "11d.png"},
}

var DrizzleConditions = []*ConditionData{
	&ConditionData{ID: 300, Meaning: "light intensity drizzle", Icon1: "09d.png"},
	&ConditionData{ID: 301, Meaning: "drizzle", Icon1: "09d.png"},
	&ConditionData{ID: 302, Meaning: "heavy intensity drizzle", Icon1: "09d.png"},
	&ConditionData{ID: 310, Meaning: "light intensity drizzle rain", Icon1: "09d.png"},
	&ConditionData{ID: 311, Meaning: "drizzle rain", Icon1: "09d.png"},
	&ConditionData{ID: 312, Meaning: "heavy intensity drizzle rain", Icon1: "09d.png"},
	&ConditionData{ID: 313, Meaning: "shower rain and drizzle", Icon1: "09d.png"},
	&ConditionData{ID: 314, Meaning: "heavy shower rain and drizzle", Icon1: "09d.png"},
	&ConditionData{ID: 321, Meaning: "shower drizzle", Icon1: "09d.png"},
}

var RainConditions = []*ConditionData{
	&ConditionData{ID: 500, Meaning: "light rain", Icon1: "09d.png"},
	&ConditionData{ID: 501, Meaning: "moderate rain", Icon1: "09d.png"},
	&ConditionData{ID: 502, Meaning: "heavy intensity rain", Icon1: "09d.png"},
	&ConditionData{ID: 503, Meaning: "very heavy rain", Icon1: "09d.png"},
	&ConditionData{ID: 504, Meaning: "extreme rain", Icon1: "09d.png"},
	&ConditionData{ID: 511, Meaning: "freezing rain", Icon1: "13d.png"},
	&ConditionData{ID: 520, Meaning: "light intensity shower rain", Icon1: "09d.png"},
	&ConditionData{ID: 521, Meaning: "shower rain", Icon1: "09d.png"},
	&ConditionData{ID: 522, Meaning: "heavy intensity shower rain", Icon1: "09d.png"},
	&ConditionData{ID: 531, Meaning: "ragged shower rain", Icon1: "09d.png"},
}

var SnowConditions = []*ConditionData{
	&ConditionData{ID: 600, Meaning: "light snow", Icon1: "13d.png"},
	&ConditionData{ID: 601, Meaning: "snow", Icon1: "13d.png"},
	&ConditionData{ID: 602, Meaning: "heavy snow", Icon1: "13d.png"},
	&ConditionData{ID: 611, Meaning: "sleet", Icon1: "13d.png"},
	&ConditionData{ID: 612, Meaning: "shower sleet", Icon1: "13d.png"},
	&ConditionData{ID: 615, Meaning: "light rain and snow", Icon1: "13d.png"},
	&ConditionData{ID: 616, Meaning: "rain and snow", Icon1: "13d.png"},
	&ConditionData{ID: 620, Meaning: "light shower snow", Icon1: "13d.png"},
	&ConditionData{ID: 621, Meaning: "shower snow", Icon1: "13d.png"},
	&ConditionData{ID: 622, Meaning: "heavy shower snow", Icon1: "13d.png"},
}

var AtmosphereConditions = []*ConditionData{
	&ConditionData{ID: 701, Meaning: "mist", Icon1: "50d.png"},
	&ConditionData{ID: 711, Meaning: "smoke", Icon1: "50d.png"},
	&ConditionData{ID: 721, Meaning: "haze", Icon1: "50d.png"},
	&ConditionData{ID: 731, Meaning: "sand, dust whirls", Icon1: "50d.png"},
	&ConditionData{ID: 741, Meaning: "fog", Icon1: "50d.png"},
	&ConditionData{ID: 751, Meaning: "sand", Icon1: "50d.png"},
	&ConditionData{ID: 761, Meaning: "dust", Icon1: "50d.png"},
	&ConditionData{ID: 762, Meaning: "volcanic ash", Icon1: "50d.png"},
	&ConditionData{ID: 771, Meaning: "squalls", Icon1: "50d.png"},
	&ConditionData{ID: 781, Meaning: "tornado", Icon1: "50d.png"},
}

var CloudConditions = []*ConditionData{
	&ConditionData{ID: 800, Meaning: "clear sky", Icon1: "01d.png", Icon2: "01n.png"},
	&ConditionData{ID: 801, Meaning: "few clouds", Icon1: "02d.png", Icon2: " 02n.png"},
	&ConditionData{ID: 802, Meaning: "scattered clouds", Icon1: "03d.png", Icon2: "03d.png"},
	&ConditionData{ID: 803, Meaning: "broken clouds", Icon1: "04d.png", Icon2: "03d.png"},
	&ConditionData{ID: 804, Meaning: "overcast clouds", Icon1: "04d.png", Icon2: "04d.png"},
}

var ExtremeConditions = []*ConditionData{
	&ConditionData{ID: 900, Meaning: "tornado", Icon1: ""},
	&ConditionData{ID: 901, Meaning: "tropical storm", Icon1: ""},
	&ConditionData{ID: 902, Meaning: "hurricane", Icon1: ""},
	&ConditionData{ID: 903, Meaning: "cold", Icon1: ""},
	&ConditionData{ID: 904, Meaning: "hot", Icon1: ""},
	&ConditionData{ID: 905, Meaning: "windy", Icon1: ""},
	&ConditionData{ID: 906, Meaning: "hail", Icon1: ""},
}

var AdditionalConditions = []*ConditionData{
	&ConditionData{ID: 951, Meaning: "calm", Icon1: ""},
	&ConditionData{ID: 952, Meaning: "light breeze", Icon1: ""},
	&ConditionData{ID: 953, Meaning: "gentle breeze", Icon1: ""},
	&ConditionData{ID: 954, Meaning: "moderate breeze", Icon1: ""},
	&ConditionData{ID: 955, Meaning: "fresh breeze", Icon1: ""},
	&ConditionData{ID: 956, Meaning: "strong breeze", Icon1: ""},
	&ConditionData{ID: 957, Meaning: "high wind, near gale", Icon1: ""},
	&ConditionData{ID: 958, Meaning: "gale", Icon1: ""},
	&ConditionData{ID: 959, Meaning: "severe gale", Icon1: ""},
	&ConditionData{ID: 960, Meaning: "storm", Icon1: ""},
	&ConditionData{ID: 961, Meaning: "violent storm", Icon1: ""},
	&ConditionData{ID: 962, Meaning: "hurricane", Icon1: ""},
}
