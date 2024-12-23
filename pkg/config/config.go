package config

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

// LLMConfig represents the top-level configuration structure
type LLMConfig struct {
	LLMBackend LLMBackendConfig `yaml:"llm_backend"`
}

// LLMBackendConfig represents the configuration for the LLM backend
type LLMBackendConfig struct {
	Type        string  `yaml:"type"`
	APIKey      string  `yaml:"api_key"`
	Model       string  `yaml:"model"`
	Temperature float64 `yaml:"temperature"`
	BaseURL     string  `yaml:"base_url"`
}

// LoadLLMConfig reads and parses the LLM configuration from the specified YAML file
func LoadLLMConfig(configPath string) (*LLMConfig, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config LLMConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	return &config, nil
}

// Validate checks if the configuration has all required fields
func (c *LLMConfig) Validate() error {
	if c.LLMBackend.Type == "" {
		return fmt.Errorf("LLM backend type is required")
	}
	if c.LLMBackend.APIKey == "" {
		return fmt.Errorf("API key is required")
	}
	if c.LLMBackend.Model == "" {
		return fmt.Errorf("Model name is required")
	}
	return nil
}
