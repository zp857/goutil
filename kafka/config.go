package kafka

import "time"

type Config struct {
	Consumer struct {
		Urls []string `yaml:"urls" json:"urls"`
	} `yaml:"consumer" json:"consumer"`
	Producer struct {
		Urls []string `yaml:"urls" json:"urls"`
	} `yaml:"producer" json:"producer"`
}

const (
	defaultSleep = 3 * time.Second
)
