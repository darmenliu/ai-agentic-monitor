package config

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

// LLMBackendConfig represents the configuration for the LLM backend
type LLMBackendYamlConfig struct {
	config LLMConfig
	path   string
}

func NewLLMBackendYamlConfig(configPath string) (*LLMBackendYamlConfig, error) {
	config := &LLMBackendYamlConfig{
		path: configPath,
	}
	if err := config.loadLLMConfig(configPath); err != nil {
		return nil, err
	}
	return config, nil
}

// LoadLLMConfig reads and parses the LLM configuration from the specified YAML file
func (c *LLMBackendYamlConfig) loadLLMConfig(configPath string) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	if err := yaml.Unmarshal(data, c.config); err != nil {
		return fmt.Errorf("error parsing config file: %w", err)
	}

	if err := c.validate(); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	return nil
}

// Validate checks if the configuration has all required fields
func (c *LLMBackendYamlConfig) validate() error {
	if c.config.Type == "" {
		return fmt.Errorf("LLM backend type is required")
	}
	if c.config.APIKey == "" {
		return fmt.Errorf("API key is required")
	}
	if c.config.Model == "" {
		return fmt.Errorf("model name is required")
	}
	return nil
}

func (c *LLMBackendYamlConfig) GetLLMType() string {
	return c.config.Type
}

func (c *LLMBackendYamlConfig) GetModel() string {
	return c.config.Model
}

func (c *LLMBackendYamlConfig) GetAPIKey() string {
	return c.config.APIKey
}

func (c *LLMBackendYamlConfig) GetBaseURL() string {
	return c.config.BaseURL
}

func (c *LLMBackendYamlConfig) GetTemperature() float64 {
	return c.config.Temperature
}
