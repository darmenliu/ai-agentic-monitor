package main

import (
	"fmt"

	"github.com/darmenliu/ai-agentic-monitor/pkg/config"
	"github.com/darmenliu/ai-agentic-monitor/pkg/monitor"
)

func main() {
	config, err := config.NewLLMBackendYamlConfig("./config/llm_config.yml")
	if err != nil {
		fmt.Println(err)
		return
	}

	monitor := monitor.NewMonitor(config, "check the status of the system if there are any issues like performence, memory, etc.")
	err = monitor.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
