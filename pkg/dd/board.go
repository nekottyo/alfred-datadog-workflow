package dd

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"gopkg.in/zorkian/go-datadog-api.v2"
)

// Board is workflow structure of the board.
type Board struct {
	client    *datadog.Client
	cacheName string
	wf        *aw.Workflow
}

// NewBoard returns a new dd.Board.
func NewBoard(client *datadog.Client, wf *aw.Workflow) Board {
	return Board{
		cacheName: boardCacheName(),
		client:    client,
		wf:        wf,
	}
}

// ListBoards is fetch board and appends to workflow items.
func (d *Board) ListBoards() error {
	var boards []datadog.BoardLite

	call := func() (interface{}, error) {
		return d.client.GetBoards()
	}

	if err := d.wf.Cache.LoadOrStoreJSON(d.cacheName, maxCacheAge, call, &boards); err != nil {
		return err
	}

	for _, board := range boards {
		url := fmt.Sprintf("%s/dashboard/%s", baseUrl, board.GetId())
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
