package main

import (
	"time"

	"github.com/darmenliu/ai-agentic-monitor/pkg/monitor"
	"github.com/pterm/pterm"
)

const (
	// interval of all the monitors to run
	interval_of_monitors = 5 // in minutes
)

type AIMonitors interface {
	run() error
}

type MonitorManager struct {
	monitors map[string]monitor.Monitor
}

func NewMonitorManager() *MonitorManager {
	return &MonitorManager{
		monitors: make(map[string]monitor.Monitor),
	}
}

func (m *MonitorManager) AddMonitor(name string, mon monitor.Monitor) {
	m.monitors[name] = mon
}

func (m *MonitorManager) Run() error {
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)
	for _, mon := range m.monitors {
		go func(mon monitor.Monitor) {
			ticker := time.NewTicker(interval_of_monitors * time.Second)
			defer ticker.Stop()
			for range ticker.C {
				err := mon.Run()
				if err != nil {
					// handle error (e.g., log it)
					logger.Error("ai-agentic-monitor: failed to run monitor,", logger.Args("err", err.Error()))
				}
			}
		}(mon)
	}
	return nil
}
