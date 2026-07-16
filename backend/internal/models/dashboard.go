package models

type DashboardSummary struct {
	RunningServices int     `json:"runningServices"`
	DeploymentsToday int    `json:"deploymentsToday"`
	ActiveAlerts    int     `json:"activeAlerts"`
	CloudCost       float64 `json:"cloudCost"`
	HealthyServices int     `json:"healthyServices"`
}