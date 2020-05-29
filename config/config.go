package config

import (
	"errors"
	"fmt"
	"github.com/InoGo-Software/downtime-notifier/healthcheck"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	HealthChecks []healthcheck.HealthCheck
}

func Load() (*Config, error) {
	// Read the config file.
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Printf("error reading yaml file: %s", err)
		return nil, err
	}

	// Parse config.
	var yamlConfig Config
	if err := yaml.UnmarshalStrict(yamlFile, &yamlConfig); err != nil {
		fmt.Printf("error parsing yaml file: %s", err)
		return nil, err
	}

	// Validate the config.
	if err := validate(yamlConfig); err != nil {
		fmt.Printf("error parsing yaml file: %s", err)
		return nil, err
	}

	return &yamlConfig, nil
}

func validate(cfg Config) error {
	if len(cfg.HealthChecks) == 0 {
		return errors.New("health checks are required")
	}

	for _, healthCheck := range cfg.HealthChecks {
		if err := healthCheck.Validate(); err != nil {
			return err
		}
	}
	return nil
}
