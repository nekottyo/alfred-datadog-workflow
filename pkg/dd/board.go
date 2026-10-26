package dd

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	aw "github.com/deanishe/awgo"
)

// Board is workflow structure of the board.
type Board struct {
	client    *Client
	cacheName string
	wf        *aw.Workflow
}

// NewBoard returns a new dd.Board.
func NewBoard(client *Client, wf *aw.Workflow) Board {
	return Board{
		cacheName: boardCacheName(),
		client:    client,
		wf:        wf,
	}
}

// ListBoards is fetch board and appends to workflow items.
func (d *Board) ListBoards() error {
	var dashboards []datadogV1.DashboardSummaryDefinition

	call := func() (interface{}, error) {
		summary, _, err := d.client.Dashboards.ListDashboards(d.client.Ctx)
		if err != nil {
			return nil, err
		}
		return summary.GetDashboards(), nil
	}

	if err := d.wf.Cache.LoadOrStoreJSON(d.cacheName, maxCacheAge, call, &dashboards); err != nil {
		return err
	}

	for _, board := range dashboards {
		url := fmt.Sprintf("%s/dashboard/%s", baseURL, board.GetId())
		d.wf.NewItem(board.GetTitle()).
			Subtitle(url).
			Arg(url).
			UID(board.GetTitle()).
			Valid(true)
	}
	return nil
}

// monitorCacheName returns filename for the dashbaord's cache.
func boardCacheName() string {
	return "board.json"
}
