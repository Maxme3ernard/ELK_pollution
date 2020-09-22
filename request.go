package main
import (
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	"strconv"
)
type StationsResponse struct {
	Data []Station `json:"data"`
	Status string `json:"status"`
}
type DataResponse struct {
	Data []interface{} `json:"data"`
	Status string `json:"status"`
}
type Geotag struct {
	Lat float64
	Lng float64
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
type DataPoint struct {
	
}

func main() {
	// We fetch all stations 
	
	resp, err := http.Get("https://airnet.waqi.info/airnet/map/stations")
	if err != nil {
		// handle error
		panic(err)
	}
	defer resp.Body.Close()
	respTxt, err := ioutil.ReadAll(resp.Body)

	var result StationsResponse
	json.Unmarshal(respTxt, &result)
	var stations = result.Data
	fmt.Printf("nb  of stations %d", len(stations))
	token := "abe466e87b9df8832dfe2f08d96b915adbe4cdb1"
	
	for index, element:= range stations{
		
		if (index+1)%500 == 0 {
			// We are limited to 1000 calls a second 
			// We stop every 500 stations and sleep for 1 second, just to be safe
			fmt.Printf("stop cond reached %d \n",index)
			time.Sleep(1 * time.Second)	
		}
		// fmt.Printf("station n %d: %s \n", index,element.N)
		fmt.Printf("station %d geotag (%f,%f) \n", index,element.G[0],element.G[1])
		var requestUrl = "https://api.waqi.info/feed/geo:"+strconv.FormatFloat(element.G[0], 'E', -1, 64)+";"+strconv.FormatFloat(element.G[1], 'E', -1, 64)+"/?token="+token
		resp, err = http.Get(requestUrl)
		if err != nil {
			// handle error
			panic(err)
		}
		defer resp.Body.Close()
		respTxt, err := ioutil.ReadAll(resp.Body)
	
		m := make(map[string]interface{})
		err := json.Unmarshal(data, &m)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(m["id"])
	}
}