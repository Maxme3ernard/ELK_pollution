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
	Data   Global `json:"data"`
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
type Global struct {
	Status string
	Data  Data
}

type Data struct {
	Aqi int
	Idx int
	Attributions [] struct {
	Url string
	Name string
	Logo string
}
	City  City
	Dominentpol string
	Iaqi Iaqi
	Time Time
	Forecast Forecast
	Debug Debug

}


type City struct {
	Geo []float64
	Name string
	Url string
}

type Iaqi struct {
    H struct  {
    	V float64
    }
	P struct  {
    	V float64
    }
	Pm10 struct  {
    	V float64
    }
	Pm25  struct  {
    	V float64
    }
	T  struct  {
    	V float64
    }
	W  struct  {
    	V float64
    }
	Wg  struct  {
    	V float64
    }
}



type Time struct {
	S string
	Tz string
	V float64
	Iso string
}

type Forecast struct {
	Daily Daily
}

type Daily struct {
	O3 []  struct {
	Avg int
	Day string
	Max int
	Min int
}
	Pm10 []  struct {
	Avg int
	Day string
	Max int
	Min int
}
	Pm25 []  struct {
	Avg int
	Day string
	Max int
	Min int
}
	Uvi []  struct {
	Avg int
	Day string
	Max int
	Min int
}
}

type Debug struct {
	Sync string
}

func NewSniffer(pb *Polutbeat, url string, token string) *Sniffer {
	s := &Sniffer{
		pb: pb,
		URL: url,
		Token: token,
	}

	return s
}

func (s *Sniffer) Run() error{
	// fetch all Stations
	respTxt := getRequestResponseAsBytes(s.URL)
	var result StationsResponse
	json.Unmarshal(respTxt, &result)
	s.Stations = result.Data
	fmt.Println("nb  of stations %d", len(s.Stations))
	// s.getStationData(s.Stations[0].G[0], s.Stations[0].G[1])
	 for index, element := range s.Stations {
	 	if (index+1)%100 == 0 {
	 		// We are limited to 1000 calls a second
	 		// We stop every 500 stations and sleep for 1 second, just to be safe
	 		// fmt.Println("stop cond reached %d ", index)
	 		time.Sleep(1 * time.Second)
	 	}
	 	// fmt.Printf("station n %d: %s \n", index, element.N)
	 	go s.getStationData(element.G[0], element.G[1])
	 }
	 fmt.Println("run done")
	 return nil
}

func (s *Sniffer) getStationData(lat float64, lng float64) error {
	now := time.Now()
	requestURL := "https://api.waqi.info/feed/geo:" + strconv.FormatFloat(lat, 'f', -1, 64) + ";" + strconv.FormatFloat(lng, 'f', -1, 64) + "/?token=" + s.Token
	respTxt := getRequestResponseAsBytes(requestURL)
	var m Global
	json.Unmarshal(respTxt, &m)
	event := ApiEvent{
		Timestamp:	now,
		Data: 		m.Data,
	}
	s.pb.client.Publish(event.ToBeatEvent())
	return nil
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

func (s *Sniffer) Stop()error {
	fmt.Println("sniffer stop called")
	return nil
}
