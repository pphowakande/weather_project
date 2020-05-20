# The Go Programming Language

### Download and Install

#### Binary Distributions

Official binary distributions are available at https://golang.org/dl/.

After downloading a binary release, visit https://golang.org/doc/install
or load [doc/install.html](./doc/install.html) in your web browser for installation
instructions.

#### Install From Source

If a binary distribution is not available for your combination of
operating system and architecture, visit
https://golang.org/doc/install/source or load [doc/install-source.html](./doc/install-source.html)
in your web browser for source installation instructions.

### Contributing

Go is the work of thousands of contributors. We appreciate your help!

To contribute, please read the contribution guidelines:
	https://golang.org/doc/contribute.html

Note that the Go project uses the issue tracker for bug reports and
proposals only. See https://golang.org/wiki/Questions for a list of
places to ask questions about the Go language.

[rf]: https://reneefrench.blogspot.com/
[cc3-by]: https://creativecommons.org/licenses/by/3.0/

# WEATHER PROJECT

## Installation
To install the packages you need to the cmd folder and run the following command

`go get ./...`

After installing all the packages update the configuration file config.yaml with database and email configurations

## Run

To run the service execute the following command
`go run app.go`

To upload weather codes data execute the following command
`go run src/internal/upload_weather_codes.go`

To upload city data execute the following command
`go run src/internal/upload_city_data.go`

To get historical data execute the following command
`go run src/internal/get_historical_data.go`

## Build App

`docker build -t weather-app .`

## run docker image

`docker run -p 8080:8081 -it weather-app`



