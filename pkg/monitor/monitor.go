package monitor

import (
	"context"
	"fmt"

	"github.com/darmenliu/ai-agentic-monitor/pkg/agents"
	"github.com/darmenliu/ai-agentic-monitor/pkg/config"
	"github.com/darmenliu/ai-agentic-monitor/pkg/llmback"

	"github.com/pterm/pterm"
	lcagents "github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/tools"
)

type Monitor interface {
	// Monitor will monitor the system and return the status of the system.
	Run() error
}

type MonitorImpl struct {
	config config.LLMConfiger
	prompt string
}

func NewMonitor(config config.LLMConfiger, prompt string) Monitor {
	return &MonitorImpl{
		config: config,
		prompt: prompt,
	}
}

func (m *MonitorImpl) Run() error {

	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)
	llmbak, err := llmback.NewLLMBackend(context.Background(), m.config)
	if err != nil {
		logger.Error("ai-agentic-monitor: failed to get LLM backend,", logger.Args("err", err.Error()))
		return err
	}

	agentTools := []tools.Tool{
		&agents.ScriptExecutor{},
	}
	agent := agents.NewMonitorAgent(llmbak.GetModel(), agentTools, "output", nil)
	executor := lcagents.NewExecutor(agent)
	answer, err := chains.Run(context.Background(), executor, m.prompt)
	if err != nil {
		logger.Error("ai-agentic-monitor: failed to run agent,", logger.Args("err", err.Error()))
		return err
	}

	fmt.Println("ai-agentic-monitor: " + answer)
	return nil
}