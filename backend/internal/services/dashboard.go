package services

import "github.com/Vaishnavi9843/engineering-command-center/internal/models"

func GetDashboardSummary() models.DashboardSummary {
	return models.DashboardSummary{
		RunningServices: 112,
		DeploymentsToday: 15,
		ActiveAlerts: 4,
		CloudCost: 126,
		HealthyServices: 5,
	}
}