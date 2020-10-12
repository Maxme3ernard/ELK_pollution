package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"

	"github.com/Maxme3ernard/polutbeat/config"
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

	sniff = NewSniffer(bt, bt.config.URL, bt.config.Token)


	for {
		select {
		case <-bt.done:
			return nil
		}

		// compute data
		// bt.client.Publish(event)
		logp.Info("Data sent")
	}
}

// Stop stops polutbeat.
func (bt *polutbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
