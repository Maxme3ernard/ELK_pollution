// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period	time.Duration `config:"period"`
	URL 		string				`config:"url"`
	Token		string				`config:"token"`
	Timeout time.Duration `config:"timeout"`
}

var DefaultConfig = Config{
	Period: 15 * time.Second,
	URL:		"https://airnet.waqi.info/airnet/map/stations",
	Token:	"abe466e87b9df8832dfe2f08d96b915adbe4cdb1",
}
