package config

import "time"

type AthenaConfig struct {
	Region         string
	OutputLocation string
	DbName         string
	PollFrequency  time.Duration
}

func NewConfig(region string, dbName string, outputLocation string, pollFrequency time.Duration) *AthenaConfig {
	return &AthenaConfig{
		Region:         region,
		OutputLocation: outputLocation,
		DbName:         dbName,
		PollFrequency:  pollFrequency,
	}
}
