package connector

import (
	"errors"
	"io/ioutil"

	"github.com/cloudfoundry-incubator/candiedyaml"
)

type Config struct {
	Debug   bool           `yaml:"debug"`
	Balabit *BalabitConfig `yaml:"balabit"`
}

type PubSubConfig struct {
	GcpProjectName   string
	SubscriberId     string
	NumberOfMessages int
	KeyfileLocation  string
}

type BalabitConfig struct {
	EndpointAddress     string `yaml:"endpointAddress"`
	Tag                 string `yaml:"tag"`
	ImpersonateHostname string `yaml:"impersonateHostname"`
	CACert              string `yaml:"caCert"`
	ServerName          string `yaml:"serverName"`
	SkipSSLVerify       bool   `yaml:"skipSSLVerify"`
	UseTLS              bool   `yaml:"useTLS"`
}

func NewConfig(configPath string) (*Config, error) {
	file, err := ioutil.ReadFile(configPath)

	if err != nil {
		return nil, err
	}

	var config Config

	err = candiedyaml.Unmarshal(file, &config)

	if err != nil {
		return nil, err
	}

	if config.Balabit == nil {
		return nil, errors.New("downstream service key missing from config")
	}

	return &config, nil
}
