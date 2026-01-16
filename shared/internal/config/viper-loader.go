package config

import (
	"sync"

	"github.com/spf13/viper"
)

type Loader struct {
	ConfigService interface{}
	ServiceName   string
	viper         *viper.Viper
	sync          sync.Mutex
	watchers      []func() error
	//еще 1 поле
}

func NewViperLoader(cfgPtr interface{}, serviceName string) (*Loader, error) {
	l := &Loader{
		ConfigService: cfgPtr,
		ServiceName:   serviceName,
		viper:         viper.New(),
	}
	if configFilesExist() == false {

	}

	loadPriority()

	if err := viper.Unmarshal(l.ConfigService); err != nil {
		return nil, err
	}

	return l, nil
}
