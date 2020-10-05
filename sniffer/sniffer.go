package sniffer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type StationsResponse struct {
	Data   []Station `json:"data"`
	Status string    `json:"status"`
}
type DataResponse struct {
	Data   DataPoint `json:"data"`
	Status string    `json:"status"`
}
type Geotag struct {
	Lat float64
	Lng float64
}
type DataPoint struct {
	Aqi          int
	Idx          int
	Attributions []interface{}
	city         []interface{}
	DominentPol  string
	iaqi         []interface{}
	time         []interface{}
	forecast     []interface{}
	debug        []interface{}
}
type Station struct {
	G []float64
	N string
	U int
	C int
	X string
	A int
	S string
	Z string
}

type Sniffer struct {
  token
}

func main() {
	// We fetch all stations

	respTxt := getRequestResponseAsBytes("https://airnet.waqi.info/airnet/map/stations")

	var result StationsResponse
	json.Unmarshal(respTxt, &result)
	var stations = result.Data
	fmt.Println("nb  of stations %d", len(stations))

	for index, element := range stations {

		if (index+1)%100 == 0 {
			// We are limited to 1000 calls a second
			// We stop every 500 stations and sleep for 1 second, just to be safe
			// fmt.Println("stop cond reached %d ", index)
			time.Sleep(10 * time.Second)
		}
		fmt.Printf("station n %d: %s \n", index, element.N)
		go getStationData(element.G[0], element.G[1])

	}
}
func getStationData(lat float64, lng float64) {
	var requestURL = getAPIURL(lat, lng)
	respTxt := getRequestResponseAsBytes(requestURL)
	var m DataResponse
	json.Unmarshal(respTxt, &m)
	fmt.Printf("Air quality index : %s\n", strconv.Itoa(m.Data.Aqi))
}

func getAPIURL(lat float64, lng float64) string {
	token := "abe466e87b9df8832dfe2f08d96b915adbe4cdb1"
	return "https://api.waqi.info/feed/geo:" + strconv.FormatFloat(lat, 'f', -1, 64) + ";" + strconv.FormatFloat(lng, 'f', -1, 64) + "/?token=" + token
}

func getRequestResponseAsBytes(requestURL string) []byte {
	resp, err := http.Get(requestURL)
	if err != nil {
		// handle error
		panic(err)
	}
	defer resp.Body.Close()
	respTxt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		panic(err)
	}
	return respTxt
}
