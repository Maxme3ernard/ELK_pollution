package beater

import (
	"time"
  "github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
)

type ApiEvent struct {
	Timestamp    time.Time
	Fields       common.MapStr
  Data         DataPoint
}

func (a *ApiEvent) ToBeatEvent() beat.Event {
	event := beat.Event{
		Timestamp: a.Timestamp,
	}

	event.Fields = common.MapStr{
      "Aqi":          a.Data.Aqi,
    	"Idx":          a.Data.Idx,
    	"Attributions": a.Data.Attributions,
    	"city":         a.Data.city,
    	"DominentPol":  a.Data.DominentPol,
    	"iaqi":         a.Data.iaqi,
    	"time":         a.Data.time,
    	"forecast":     a.Data.forecast,
    	"debug":        a.Data.debug,
  }
	
	return event
}
