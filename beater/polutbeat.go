package beater

import (
	"fmt"

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
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	sniff := NewSniffer(bt, bt.config.URL, bt.config.Token)
	sniff.Run()

	select {
	case <-bt.done:
		logp.Info("bt.done")
		return nil
	}

}

// Stop stops polutbeat.
func (bt *Polutbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
