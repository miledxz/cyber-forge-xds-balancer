package processor

import (
	"fmt"
	"os"

	"github.com/miledxz/cyber-forge-xds-balancer/apis/v1alpha1"
	"gopkg.in/yaml.v2"
)

// parseYaml takes in a yaml envoy config and returns a typed version
func parseYaml(file string) (*v1alpha1.EnvoyConfig, error) {
	var config v1alpha1.EnvoyConfig

	yamlFile, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file: %s", err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
