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
	//var wg sync.WaitGroup
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}
	for {
		fmt.Println("new loop")
		select {
			case <-bt.done:
				logp.Info("bt.done")
				return nil
			default:

		}
		fmt.Println("previous wg done")
		fmt.Println("Wait time %d",bt.config.Period)
		sniff := NewSniffer(bt, bt.config.URL, bt.config.Token)
		fmt.Println("running sniff")
		//wg.Add(1)
		go sniff.Run()
		fmt.Println("wait")
		time.Sleep(bt.config.Period)
		fmt.Println("go on")
		//wg.Wait()

		//defer time.Sleep(bt.config.Period)
	}
	return nil

}

// Stop stops polutbeat.
func (bt *Polutbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
