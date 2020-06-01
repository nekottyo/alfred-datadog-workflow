package dd

import (
	"fmt"
	"strconv"

	aw "github.com/deanishe/awgo"
	"gopkg.in/zorkian/go-datadog-api.v2"
)

// Monitor is workflow structure of the Monitor.
type Monitor struct {
	client    *datadog.Client
	cacheName string
	wf        *aw.Workflow
}

// NewMonitor returns a new dd.Monitor.
func NewMonitor(client *datadog.Client, wf *aw.Workflow) Monitor {
	return Monitor{
		cacheName: monitorCacheName(),
		client:    client,
		wf:        wf,
	}
}

// ListMonitors is fetch monitor and appends to workflow items.
func (d *Monitor) ListMonitors() error {
	var monitors []datadog.Monitor

	call := func() (interface{}, error) {
		return d.client.GetMonitors()
	}

	if err := d.wf.Cache.LoadOrStoreJSON(d.cacheName, maxCacheAge, call, &monitors); err != nil {
		return err
	}

	for _, monitor := range monitors {
		url := fmt.Sprintf("%s/monitors/%d", baseURL, monitor.GetId())
		status := fmt.Sprintf("[%s] %s", monitor.GetOverallState(), monitor.GetName())
		d.wf.NewItem(status).
			Subtitle(url).
			Arg(url).
			UID(strconv.Itoa(monitor.GetId())).
			Valid(true)
	}
	return nil
}

// monitorCacheName returns filename for the monitor's cache.
func monitorCacheName() string {
	return "monitor.json"
}
