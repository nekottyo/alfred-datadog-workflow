package dd

import (
	"fmt"
	"strconv"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	aw "github.com/deanishe/awgo"
)

// Monitor is workflow structure of the Monitor.
type Monitor struct {
	client    *Client
	cacheName string
	wf        *aw.Workflow
}

// NewMonitor returns a new dd.Monitor.
func NewMonitor(client *Client, wf *aw.Workflow) Monitor {
	return Monitor{
		cacheName: monitorCacheName(),
		client:    client,
		wf:        wf,
	}
}

// ListMonitors is fetch monitor and appends to workflow items.
func (d *Monitor) ListMonitors() error {
	var monitors []datadogV1.Monitor

	call := func() (interface{}, error) {
		monitors, _, err := d.client.Monitors.ListMonitors(d.client.Ctx)
		return monitors, err
	}

	if err := d.wf.Cache.LoadOrStoreJSON(d.cacheName, maxCacheAge, call, &monitors); err != nil {
		return err
	}

	for _, monitor := range monitors {
		url := fmt.Sprintf("%s/monitors/%d", baseURL, monitor.GetId())
		status := fmt.Sprintf("[%s] %s", string(monitor.GetOverallState()), monitor.GetName())
		d.wf.NewItem(status).
			Subtitle(url).
			Arg(url).
			UID(strconv.FormatInt(monitor.GetId(), 10)).
			Valid(true)
	}
	return nil
}

// monitorCacheName returns filename for the monitor's cache.
func monitorCacheName() string {
	return "monitor.json"
}
