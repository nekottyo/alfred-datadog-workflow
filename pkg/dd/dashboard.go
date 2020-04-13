package dd

import (
	"gopkg.in/zorkian/go-datadog-api.v2"
)

func QueryDashboard(client *datadog.Client) ([]datadog.DashboardLite, error) {
	dashboards, err := client.GetDashboards()
	if err != nil {
		return nil, err
	}

	return dashboards, nil
}
