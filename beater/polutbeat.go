package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"

	"github.com/Maxme3ernard/polutbeat/config"
	"github.com/Maxme3ernard/polutbeat/sniffer"
)

// polutbeat configuration.
type polutbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of polutbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &polutbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts polutbeat.
func (bt *polutbeat) Run(b *beat.Beat) error {
	logp.Info("polutbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"type":    b.Info.Name,
				"counter": counter,
			},
		}
		// fetch data
		// compute data
		bt.client.Publish(event)
		logp.Info("Event sent")
		counter++
	}
}

// Stop stops polutbeat.
func (bt *polutbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
