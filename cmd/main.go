package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/darmenliu/ai-agentic-monitor/pkg/config"
	"github.com/darmenliu/ai-agentic-monitor/pkg/monitor"
)

func main() {
	config, err := config.NewLLMBackendYamlConfig("./config/llm_config.yml")
	if err != nil {
		fmt.Println(err)
		return
	}

	monitor := monitor.NewMonitor(config, "check the status of the system if there are any issues like performance, memory, etc.")

	manager := NewMonitorManager()
	manager.AddMonitor("monitor", monitor)
	err = manager.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Wait for terminate signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
