package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	io1 "io"
	"os"
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

	// Open the file
	csvfile, err := os.Open("src/docs/Multilingual_Weather_Conditions.csv")
	if err != nil {
		fmt.Println("Couldn't open the csv file", err)
		panic(err.Error())
	}

	// Skip first row (line)
	row1, err := bufio.NewReader(csvfile).ReadSlice('\n')
	if err != nil {
		fmt.Println("Erro skipping first row of csv file", err)
		panic(err.Error())
	}
	_, err = csvfile.Seek(int64(len(row1)), io1.SeekStart)
	if err != nil {
		fmt.Println("Erro skipping first row of csv file", err)
		panic(err.Error())
	}

	// Skip first row (line)
	r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io1.EOF {
			break
		}

		if record[0] != "" {

			var WeatherCode io.WeatherCode
			WeatherCode.Code = record[0]
			WeatherCode.DayCondition = record[1]
			WeatherCode.NightCondition = record[2]

			// check if weather code already present
			exists, err := db.GetWeatherCodeUsingID(dbConn, WeatherCode.Code)
			if err != nil {
				fmt.Println("Error getting weather code from database using codeid - ", WeatherCode.Code)
				fmt.Println("WeatherCode object : ", WeatherCode)
				panic(err.Error())
			}

			// if weathercode data does not exists , insert record
			if exists == false {
				// save weather code data in database
				err = db.SaveWeatherCodeRecord(dbConn, WeatherCode)
				if err != nil {
					fmt.Println("Error saving record in database")
					fmt.Println("WeatherCode object : ", WeatherCode)
					panic(err.Error())
				}
			}
		}
	}

	fmt.Println("Finished uploading weather codes data----------------------- ")
}
