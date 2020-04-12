package dd

import (
	"fmt"

	"gopkg.in/zorkian/go-datadog-api.v2"
)

func QueryDashboard(client *datadog.Client) ([]datadog.DashboardLite, error) {
	dashboards, err := client.GetDashboards()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dashboards, nil
}
