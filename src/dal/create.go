package db

import (
	"database/sql"
	io "weather_project/src/models"

	"github.com/cridenour/go-postgis"
)

// save record
func SaveCityRecord(conn *sql.DB, CityData io.CityData) error {
	sqlStatement := `INSERT INTO master_city ( city_id, name, country, downloaded, location)
	VALUES ($1, $2, $3, $4, ST_GeomFromEWKB($5))
`

	point := postgis.PointS{4326, CityData.Longitude, CityData.Latitude}

	_, err := conn.Exec(sqlStatement, CityData.CityID, CityData.Name, CityData.Country, CityData.HistoricalDataDownloaded, point)

	return err
}

// save weather code record
func SaveWeatherCodeRecord(conn *sql.DB, codeData io.WeatherCode) error {

	sqlStatement := `
	INSERT INTO weather_codes ( code, day_condition, night_condition) VALUES ($1, $2, $3)`

	_, err := conn.Exec(sqlStatement, codeData.Code, codeData.DayCondition, codeData.NightCondition)

	return err
}
