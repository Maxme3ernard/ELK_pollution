package beater

import (
	"time"
  "github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
)

type ApiEvent struct {
	Timestamp    time.Time
	Fields       common.MapStr
  Data         Data
}

func (a *ApiEvent) ToBeatEvent() beat.Event {
	event := beat.Event{
		Timestamp: a.Timestamp,
	}
	/*
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
	*/
	event.Fields = common.MapStr{
      "Aqi":          a.Data.Aqi,
    	"Idx":          a.Data.Idx,
    	"Attributions": a.Data.Attributions,
    	"city":         a.Data.City,
    	"DominentPol":  a.Data.Dominentpol,
    	"iaqi":         a.Data.Iaqi,
    	"time":         a.Data.Time,
    	"forecast":     a.Data.Forecast,
    	"debug":        a.Data.Debug,
  }

	return event
}
