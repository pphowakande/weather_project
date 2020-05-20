package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	db "weather_project/src/dal"
	io "weather_project/src/models"

	_ "github.com/lib/pq"
)

func main() {
	dbConn, err := db.CreateDBConn()
	if err != nil {
		fmt.Println("Error connecting to database - ", err)
		panic(err.Error())
	}
	fmt.Println("Connection Established")
	defer dbConn.Close()

	// Open our jsonFile
	jsonFile, err := os.Open("src/docs/city.list.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println("Error opening city json file")
		panic(err.Error())
	}
	fmt.Println("Successfully Opened city.list.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Cities array
	var users []io.City

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	// we iterate through every user within our users array
	for i := 0; i < len(users); i++ {
		var CityData io.CityData

		CityData.CityID = strconv.Itoa(users[i].ID)
		CityData.Name = users[i].Name
		CityData.Country = users[i].Country
		CityData.Latitude = float64(users[i].Coordinates.Latitude)
		CityData.Longitude = float64(users[i].Coordinates.Longitude)
		CityData.HistoricalDataDownloaded = false

		// check if citydata already present
		exists, err := db.GetCityUsingID(dbConn, CityData.CityID)
		if err != nil {
			fmt.Println("Error getting city record from database for cityid - ", CityData.CityID)
			fmt.Println("CityData object : ", CityData)
			panic(err.Error())
		}

		// if citydata does not exists , insert record
		if exists == false {
			// save city data
			err = db.SaveCityRecord(dbConn, CityData)
			if err != nil {
				fmt.Println("Error saving record in database for cityid - ", CityData.CityID)
				fmt.Println("CityData object : ", CityData)
				panic(err.Error())
			}
		}
	}
	fmt.Println("Finished uploading city data----------------------- ")
}
