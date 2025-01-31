package alerts

import (
	"fmt"
	"sync"
)

const (
	Info    = "info"
	Warning = "warning"
	Error   = "error"
	Fatal   = "fatal"
)

// Alert struct represents an alert
type Alert struct {
	Level       string `json:"level"`       // Alert level
	Summary     string `json:"summary"`     // Alert summary
	Description string `json:"description"` // Alert description
}

type AlertsManager interface {
	AddAlert(level, summary, description string) int
	UpdateAlert(id int, level, summary, description string) bool
	DeleteAlert(id int) bool
	GetAlert(id int) (Alert, bool)
	ListAlerts() []Alert
	ProcessAlert(id int) (Alert, bool)
}

// AlertsManager manages alerts
type AlertsManagerImpl struct {
	alerts map[int]Alert // Use map to store alerts, key is alert ID
	nextID int           // Next alert ID
	mu     sync.Mutex    // Mutex to ensure concurrency safety
}

// NewAlertsManager creates a new AlertsManager
func NewAlertsManager() AlertsManager {
	return &AlertsManagerImpl{
		alerts: make(map[int]Alert),
		nextID: 1,
	}
}

// AddAlert adds a new alert
func (am *AlertsManagerImpl) AddAlert(level, summary, description string) int {
	am.mu.Lock()
	defer am.mu.Unlock()

	id := am.nextID
	am.alerts[id] = Alert{
		Level:       level,
		Summary:     summary,
		Description: description,
	}
	am.nextID++
	return id
}

// UpdateAlert updates an existing alert
func (am *AlertsManagerImpl) UpdateAlert(id int, level, summary, description string) bool {
	am.mu.Lock()
	defer am.mu.Unlock()

	if _, exists := am.alerts[id]; !exists {
		return false
	}
	am.alerts[id] = Alert{
		Level:       level,
		Summary:     summary,
		Description: description,
	}
	return true
}

// DeleteAlert deletes an alert
func (am *AlertsManagerImpl) DeleteAlert(id int) bool {
	am.mu.Lock()
	defer am.mu.Unlock()

	if _, exists := am.alerts[id]; !exists {
		return false
	}
	delete(am.alerts, id)
	return true
}

// GetAlert retrieves an alert
func (am *AlertsManagerImpl) GetAlert(id int) (Alert, bool) {
	am.mu.Lock()
	defer am.mu.Unlock()

	alert, exists := am.alerts[id]
	return alert, exists
}

// ListAlerts retrieves all alerts
func (am *AlertsManagerImpl) ListAlerts() []Alert {
	am.mu.Lock()
	defer am.mu.Unlock()

	alerts := make([]Alert, 0, len(am.alerts))
	for _, alert := range am.alerts {
		alerts = append(alerts, alert)
	}
	return alerts
}

// ProcessAlert processes an alert (simply prints alert information here)
func (am *AlertsManagerImpl) ProcessAlert(id int) (Alert, bool) {
	am.mu.Lock()
	defer am.mu.Unlock()

	alert, exists := am.alerts[id]
	if !exists {
		return Alert{}, false
	}

	fmt.Printf("Processing Alert ID: %d\n", id)
	fmt.Printf("Level: %s\n", alert.Level)
	fmt.Printf("Summary: %s\n", alert.Summary)
	fmt.Printf("Description: %s\n", alert.Description)

	// Delete the alert after processing
	delete(am.alerts, id)
	return alert, true
}
