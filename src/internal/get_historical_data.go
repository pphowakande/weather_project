package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	config "weather_project/src/config"
	db "weather_project/src/dal"
	io "weather_project/src/models"

	_ "github.com/lib/pq"
)

func main() {

	fmt.Println("Starting the program to get historical data----------------")

	dbConn, err := db.CreateDBConn()
	if err != nil {
		fmt.Println("Error connecting to database - ", err)
		panic(err.Error())
	}
	fmt.Println("Connection Established")
	defer dbConn.Close()

	historicalDateStart := "2008-01-01"
	currentTime := time.Now()
	historicalDateEnd := currentTime.Format("2006-01-02")

	// take city from database one by one
	citiesData, err := db.GetAllCities(dbConn, false)
	if err != nil {
		fmt.Println("Error getting city data from database")
		panic(err.Error())
	}

	httpClient := http.Client{}
	conf := config.GetConfig()

	for _, eachCity := range citiesData {
		// process record here
		cityName := eachCity.Name
		//cityID := eachCity.CityID

		req, err := http.NewRequest("GET", conf.WeatherEndpoint, nil)
		if err != nil {
			fmt.Println("Error requesting weather endpoint - ", err)
			fmt.Println("weather endpoint : ", conf.WeatherEndpoint)
			fmt.Println("http request : ", req)
			panic(err.Error())
		}

		q := req.URL.Query()
		q.Add("access_key", conf.APIKey)
		q.Add("query", cityName)
		q.Add("historical_date_start", historicalDateStart)
		q.Add("historical_date_end", historicalDateEnd)
		q.Add("hourly", "1")
		q.Add("interval", "1")

		fmt.Println("query : ", q)

		req.URL.RawQuery = q.Encode()

		res, err := httpClient.Do(req)
		fmt.Println("res : ", res)
		fmt.Println("res status code : ", res.StatusCode)
		if err != nil {
			fmt.Println("Error requesting http weather endpoint")
			panic(err)
		}
		defer res.Body.Close()

		if res.StatusCode == 200 {

			var apiResponse io.HistoricalWeatherResponse
			json.NewDecoder(res.Body).Decode(&apiResponse)

			fmt.Println("apiResponse request: ", apiResponse.HistoReq)
			fmt.Println("apiResponse location: ", apiResponse.HistoLoc)
			fmt.Println("apiResponse current: ", apiResponse.HistoCur)
			fmt.Println("apiResponse historical: ", apiResponse.Historical)

			// rearrange the data

			// check if data already exists

			// save the data

			// update city record for  historicalData=1

			// save last scraped date
		}
		os.Exit(0)
	}

}
