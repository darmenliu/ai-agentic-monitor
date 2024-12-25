package llmback

import (
	"context"
	"fmt"

	"github.com/darmenliu/ai-agentic-monitor/pkg/config"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/llms/openai"
)

type ContentGenerator interface {
	GenerateText(ctx context.Context, prompt string) (string, error)
}

type LLMBackend struct {
	model  llms.Model
	config config.LLMConfiger
}

func NewLLMBackend(ctx context.Context, config config.LLMConfiger) (*LLMBackend, error) {
	l := &LLMBackend{
		config: config,
	}
	if err := l.initClient(ctx); err != nil {
		return nil, err
	}
	return l, nil
}

func (l *LLMBackend) initClient(ctx context.Context) error {
	var model llms.Model
	var err error
	switch l.config.GetLLMType() {
	case "gemini":
		model, err = googleai.New(ctx, googleai.WithAPIKey(l.config.GetAPIKey()), googleai.WithDefaultModel(l.config.GetModel()))
	case "ollama":
		model, err = ollama.New(ollama.WithModel(l.config.GetModel()), ollama.WithServerURL(l.config.GetBaseURL()))
	case "groq":
		model, err = openai.New(
			openai.WithModel("llama3-8b-8192"),
			openai.WithBaseURL("https://api.groq.com/openai/v1"),
			openai.WithToken(l.config.GetAPIKey()),
		)
	case "deepseek":
		model, err = openai.New(
			openai.WithModel(l.config.GetModel()),
			openai.WithBaseURL(l.config.GetBaseURL()),
			openai.WithToken(l.config.GetAPIKey()),
		)
	case "claude":
		model, err = anthropic.New(
			anthropic.WithModel("claude-3-5-sonnet-20240620"),
			anthropic.WithToken(l.config.GetAPIKey()),
		)
	default:
		return fmt.Errorf("unknown LLM backend: %s", l.config.GetLLMType())
	}

	if err != nil {
		return err
	}
	l.model = model
	return nil
}

func (l *LLMBackend) GenerateText(ctx context.Context, prompt string) (string, error) {
	resp, err := llms.GenerateFromSinglePrompt(ctx, l.model, prompt, llms.WithTemperature(l.config.GetTemperature()))
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}
	return resp, nil
}

func (l *LLMBackend) GetModel() llms.Model {
	return l.model
}
