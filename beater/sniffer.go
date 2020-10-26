package beater

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Sniffer struct {
	pb			*Polutbeat
	URL			string
	Token 	string
	Timeout time.Duration
	Stations   []Station
}

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

func NewSniffer(pb *Polutbeat, url string, token string) *Sniffer {
	s := &Sniffer{
		pb: pb,
		URL: url,
		Token: token,
	}

	return s
}

func (s *Sniffer) Run() {
	// fetch all Stations
	respTxt := getRequestResponseAsBytes(s.URL)
	var result StationsResponse
	json.Unmarshal(respTxt, &result)
	s.Stations = result.Data
	fmt.Println("nb  of stations %d", len(s.Stations))
	//s.getStationData(s.Stations[0].G[0], s.Stations[0].G[1])
	 for index, element := range s.Stations {
	 	if (index+1)%100 == 0 {
	 		// We are limited to 1000 calls a second
	 		// We stop every 500 stations and sleep for 1 second, just to be safe
	 		// fmt.Println("stop cond reached %d ", index)
	 		//time.Sleep(10 * time.Second)
	 	}
	 	// fmt.Printf("station n %d: %s \n", index, element.N)
	 	go s.getStationData(element.G[0], element.G[1])
	 }
}

func (s *Sniffer) getStationData(lat float64, lng float64) {
	now := time.Now()
	requestURL := "https://api.waqi.info/feed/geo:" + strconv.FormatFloat(lat, 'f', -1, 64) + ";" + strconv.FormatFloat(lng, 'f', -1, 64) + "/?token=" + s.Token
	respTxt := getRequestResponseAsBytes(requestURL)
	var m DataResponse
	json.Unmarshal(respTxt, &m)
	event := ApiEvent{
		Timestamp:	now,
		Data: 		m.Data,
	}
	s.pb.client.Publish(event.ToBeatEvent())
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

func (s *Sniffer) Stop(){
}
