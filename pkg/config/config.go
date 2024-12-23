package config

type LLMConfiger interface {
	GetLLMType() string
	GetModel() string
	GetAPIKey() string
	GetBaseURL() string
	GetTemperature() float64
}

type LLMConfig struct {
	Type        string  `yaml:"type"`
	APIKey      string  `yaml:"api_key"`
	Model       string  `yaml:"model"`
	Temperature float64 `yaml:"temperature"`
	BaseURL     string  `yaml:"base_url"`
}
