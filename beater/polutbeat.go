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
type Polutbeat struct {
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

	bt := &Polutbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts polutbeat.
func (bt *Polutbeat) Run(b *beat.Beat) error {
	logp.Info("polutbeat is running! Hit CTRL-C to stop it.")

	var err error
	var sniff = nil
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}


	for {
		select {
		case <-bt.done:
			logp.Info("bt.done")

			return nil
		default:
				timeout := bt.config.Timeout
				time.Sleep(timeout)
				logp.Info("running new sniffer after %d",timeout)
				if sniff!=nil {
					sniff := NewSniffer(bt, bt.config.URL, bt.config.Token)
					sniff.Run()

				}

		}
	}
}

// Stop stops polutbeat.
func (bt *Polutbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
