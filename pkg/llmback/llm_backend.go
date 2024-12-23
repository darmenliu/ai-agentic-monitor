package llmback

import (
	"context"
)

type LLMBackend interface {
	GenerateText(ctx context.Context, prompt string) (string, error)
}
