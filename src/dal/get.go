package db

import (
	"database/sql"
	"fmt"
	io "weather_project/src/models"
)

// get city data using cityid
func GetCityUsingID(conn *sql.DB, cityid string) (bool, error) {
	var col string
	sqlStatement := `SELECT city_id FROM master_city WHERE city_id=$1`
	row := conn.QueryRow(sqlStatement, cityid)
	err := row.Scan(&col)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return true, err
	}
	if col != "" {
		return true, err
	}
	return false, err
}

// get weather code  data using code id
func GetWeatherCodeUsingID(conn *sql.DB, codeid string) (bool, error) {
	var col string
	sqlStatement := `SELECT code FROM weather_codes WHERE code=$1`
	row := conn.QueryRow(sqlStatement, codeid)
	err := row.Scan(&col)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return true, err
	}
	if col != "" {
		return true, err
	}
	return false, err
}

// get all cities
func GetAllCities(conn *sql.DB, downloaded bool) ([]io.CityDataDb, error) {
	var cities []io.CityDataDb
	sqlStatement := `SELECT id, city_id, name, country, location FROM master_city WHERE downloaded=$1`
	rows, err := conn.Query(sqlStatement, downloaded)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var city io.CityDataDb
		err = rows.Scan(&city.ID, &city.CityID, &city.Name, &city.Country, &city.Location)
		if err != nil {
			panic(err)
		}
		cities = append(cities, city)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("total cities : ", len(cities))

	return cities, err

}
