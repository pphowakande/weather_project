package io

// city struct which contains
type City struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Country     string `json:"country"`
	Coordinates Coord  `json:"coord"`
}

type Coord struct {
	Longitude float32 `json:"lon"`
	Latitude  float32 `json:"lat"`
}

type Location struct {
	Name string `json:"name"`
}

type Weather struct {
	Temperature int `json:"temperature"`
}

type Response struct {
	Location Location `json:"location"`
	Current  Weather  `json:"current"`
}

type CityDataDb struct {
	ID                       string `json:"id"`
	CityID                   string `json:"city_id"`
	Name                     string `json:"name"`
	Country                  string `json:"country"`
	HistoricalDataDownloaded bool   `json:"downloaded"`
	Location                 string `json:"location"`
}

type CityData struct {
	ID                       string  `json:"id"`
	CityID                   string  `json:"city_id"`
	Name                     string  `json:"name"`
	Country                  string  `json:"country"`
	Latitude                 float64 `json:"latitude"`
	Longitude                float64 `json:"longitude"`
	HistoricalDataDownloaded bool    `json:"downloaded"`
}

type WeatherCode struct {
	Code           string `json:"overhead_code"`
	DayCondition   string `json:"daycondition"`
	NightCondition string `json:"NightCondition"`
}

type HistoricalReq struct {
	Type     string `json:"type"`
	Query    string `json:"query"`
	Language string `json:"language"`
	Unit     string `json:"unit"`
}

type HistoricalLoc struct {
	Name           string `json:"name"`
	Country        string `json:"country"`
	Region         string `json:"region"`
	Lat            string `json:"lat"`
	Lon            string `json:"lon"`
	TimezoneID     string `json:"timezone_id"`
	Localtime      string `json:"localtime"`
	LocaltimeEpoch int    `json:"localtime_epoch"`
	UtcOffset      string `json:"utc_offset"`
}

type HistoCurrent struct {
	ObservationTime     string   `json:"observation_time"`
	Temperature         int      `json:"temperature"`
	WeatherCode         int      `json:"weather_code"`
	WeatherIcons        []string `json:"weather_icons"`
	WeatherDescriptions []string `json:"weather_descriptions"`
	WindSpeed           int      `json:"wind_speed"`
	WindDegree          int      `json:"wind_degree"`
	WindDir             string   `json:"wind_dir"`
	Pressure            int      `json:"pressure"`
	Precip              float32  `json:"precip"`
	Humidity            float32  `json:"humidity"`
	Cloudcover          float32  `json:"cloudcover"`
	Feelslike           float32  `json:"feelslike"`
	UvIndex             float32  `json:"uv_index"`
	Visibility          float32  `json:"visibility"`
}

type HistoAstro struct {
	Sunrise          string `json:"sunrise"`
	Sunset           string `json:"sunset"`
	Moonrise         string `json:"moonrise"`
	Moonset          string `json:"moonset"`
	MoonPhase        string `json:"moon_phase"`
	MoonIllumination int    `json:"moon_illumination"`
}

type HistoHourly struct {
	Time        string `json:"time"`
	Temperature int    `json:"temperature"`
	WindSpeed   int    `json:"wind_speed"`
	WindDegree  int    `json:"wind_degree"`
	WindDir     string `json:"wind_dir"`

	WeatherCode         int      `json:"weather_code"`
	WeatherIcons        []string `json:"weather_icons"`
	WeatherDescriptions []string `json:"weather_descriptions"`

	Precip     float32 `json:"precip"`
	Humidity   float32 `json:"humidity"`
	Visibility float32 `json:"visibility"`
	Pressure   float32 `json:"pressure"`

	Cloudcover float32 `json:"cloudcover"`
	HeatIndex  float32 `json:"heatindex"`
	DewPoint   float32 `json:"dewpoint"`
	WindChill  float32 `json:"windchill"`
	WindGust   float32 `json:"windgust"`
	FeelsLike  float32 `json:"feelslike"`

	ChanceOfRain     float32 `json:"chanceofrain"`
	ChanceOfRemdry   float32 `json:"chanceofremdry"`
	ChanceOfWindy    float32 `json:"chanceofwindy"`
	ChanceOfOvercast float32 `json:"chanceofovercast"`
	ChanceOfSunshine float32 `json:"chanceofsunshine"`
	ChanceOffRost    float32 `json:"chanceoffrost"`
	ChanceOfHightemp float32 `json:"chanceofhightemp"`

	ChanceOfFog     float32 `json:"chanceoffog"`
	ChanceOfSnow    float32 `json:"chanceofsnow"`
	ChanceOfThunder float32 `json:"chanceofthunder"`
	UvIndex         float32 `json:"uv_index"`
}

type Historical struct {
	CityID    string        `json:"city_id"`
	Date      string        `json:"date"`
	DateEpoch int           `json:"date_epoch"`
	Astro     HistoAstro    `json:"astro"`
	Mintemp   float32       `json:"mintemp"`
	Maxtemp   float32       `json:"maxtemp"`
	Avgtemp   float32       `json:"avgtemp"`
	Totalsnow float32       `json:"totalsnow"`
	Sunhour   float32       `json:"sunhour"`
	UvIndex   float32       `json:"uv_index"`
	Hourly    []HistoHourly `json:"hourly"`
}

type HistoricalWeatherResponse struct {
	HistoReq   HistoricalReq         `json:"request"`
	HistoLoc   HistoricalLoc         `json:"location"`
	HistoCur   HistoCurrent          `json:"current"`
	Historical map[string]Historical `json:"historical"`
}

type HistoGenericSave struct {
	CityID           string  `json:"city_id"`
	Latitude         string  `json:"lat"`
	Longitude        string  `json:"lon"`
	Date             string  `json:"date"`
	Sunrise          string  `json:"sunrise"`
	Sunset           string  `json:"sunset"`
	Moonrise         string  `json:"moonrise"`
	Moonset          string  `json:"moonset"`
	MoonPhase        string  `json:"moon_phase"`
	MoonIllumination int     `json:"moon_illumination"`
	Mintemp          float32 `json:"mintemp"`
	Maxtemp          float32 `json:"maxtemp"`
	Avgtemp          float32 `json:"avgtemp"`
	Totalsnow        float32 `json:"totalsnow"`
	Sunhour          float32 `json:"sunhour"`
	UvIndex          float32 `json:"uv_index"`
}
