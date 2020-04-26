package dd

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"gopkg.in/zorkian/go-datadog-api.v2"
)

// Dashboard is workflow structure of the dashboard.
type Dashboard struct {
	client    *datadog.Client
	cacheName string
	wf        *aw.Workflow
}

// NewDashboard returns a new dd.Dashboard.
func NewDashboard(client *datadog.Client, wf *aw.Workflow) Dashboard {
	return Dashboard{
		cacheName: dashboardCacheName(),
		client:    client,
		wf:        wf,
	}
}

// ListDashboards is fetch dashboard and appends to workflow items.
func (d *Dashboard) ListDashboards() error {
	var dashboards []datadog.DashboardLite

	call := func() (interface{}, error) {
		return d.client.GetDashboards()
	}

	if err := d.wf.Cache.LoadOrStoreJSON(d.cacheName, maxCacheAge, call, &dashboards); err != nil {
		return err
	}

	for _, dash := range dashboards {
		url := fmt.Sprintf("https://app.datadoghq.com/dash/%d/datadog", dash.GetId())
		d.wf.NewItem(dash.GetTitle()).
			Subtitle(url).
			Arg(url).
			UID(dash.GetTitle()).
			Valid(true)
	}
	return nil
}

// monitorCacheName returns filename for the dashbaord's cache.
func dashboardCacheName() string {
	return "dashboard.json"
}
