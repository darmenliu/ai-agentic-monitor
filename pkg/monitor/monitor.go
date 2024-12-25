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
	Run(prompt string) error
}

type MonitorImpl struct {
	config config.LLMConfiger
}

func NewMonitor(config config.LLMConfiger) Monitor {
	return &MonitorImpl{
		config: config,
	}
}

func (m *MonitorImpl) Run(prompt string) error {

	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)
	llmbak, err := llmback.NewLLMBackend(context.Background(), m.config)
	if err != nil {
		logger.Error("NUWA TERMINAL: failed to get LLM backend,", logger.Args("err", err.Error()))
		return err
	}

	agentTools := []tools.Tool{
		&agents.ScriptExecutor{},
	}
	agent := agents.NewMonitorAgent(llmbak.GetModel(), agentTools, "output", nil)
	executor := lcagents.NewExecutor(agent)
	answer, err := chains.Run(context.Background(), executor, prompt)
	if err != nil {
		logger.Error("NUWA TERMINAL: failed to run agent,", logger.Args("err", err.Error()))
		return err
	}

	fmt.Println("NUWA: " + answer)
	return nil
}