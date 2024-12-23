package config

import (
	"fmt"
	"os"
	"strconv"
)

type LLMBackendEnvConfig struct {
	config LLMConfig
}

func NewLLMBackendEnvConfig() (*LLMBackendEnvConfig, error) {
	config := &LLMBackendEnvConfig{}
	if err := config.loadLLMConfig(); err != nil {
		return nil, err
	}
	return config, nil
}

func (c *LLMBackendEnvConfig) loadLLMConfig() error {
	c.config.Type = os.Getenv("LLM_TYPE")
	c.config.Model = os.Getenv("LLM_MODEL")
	c.config.APIKey = os.Getenv("LLM_API_KEY")
	c.config.BaseURL = os.Getenv("LLM_BASE_URL")

	// Convert temperature string to float64
	if temp := os.Getenv("LLM_TEMPERATURE"); temp != "" {
		tempFloat, err := strconv.ParseFloat(temp, 64)
		if err != nil {
			return fmt.Errorf("invalid temperature value: %w", err)
		}
		c.config.Temperature = tempFloat
	}

	if err := c.validate(); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	return nil
}

func (c *LLMBackendEnvConfig) validate() error {
	if c.config.Type == "" {
		return fmt.Errorf("LLM_TYPE is required")
	}
	if c.config.Model == "" {
		return fmt.Errorf("LLM_MODEL is required")
	}
	if c.config.APIKey == "" {
		return fmt.Errorf("LLM_API_KEY is required")
	}
	if c.config.BaseURL == "" {
		return fmt.Errorf("LLM_BASE_URL is required")
	}
	if c.config.Temperature == 0 {
		return fmt.Errorf("LLM_TEMPERATURE is required")
	}
	return nil
}

func (c *LLMBackendEnvConfig) GetLLMType() string {
	return c.config.Type
}

func (c *LLMBackendEnvConfig) GetModel() string {
	return c.config.Model
}

func (c *LLMBackendEnvConfig) GetAPIKey() string {
	return c.config.APIKey
}

func (c *LLMBackendEnvConfig) GetBaseURL() string {
	return c.config.BaseURL
}

func (c *LLMBackendEnvConfig) GetTemperature() float64 {
	return c.config.Temperature
}
