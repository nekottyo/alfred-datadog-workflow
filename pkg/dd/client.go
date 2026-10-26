package dd

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

// Client は公式 Datadog API クライアントと認証済み context を束ねる
type Client struct {
	Dashboards *datadogV1.DashboardsApi
	Monitors   *datadogV1.MonitorsApi
	Ctx        context.Context
}

// NewClient は API キーと APP キーから認証済みクライアントを組み立てる
// サイトは既定の US1 (app.datadoghq.com) を使う
func NewClient(apiKey, appKey string) *Client {
	cfgCtx := datadog.NewDefaultContext(context.Background())
	cfgCtx = context.WithValue(cfgCtx, datadog.ContextAPIKeys, map[string]datadog.APIKey{
		"apiKeyAuth": {Key: apiKey},
		"appKeyAuth": {Key: appKey},
	})
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	return &Client{
		Dashboards: datadogV1.NewDashboardsApi(apiClient),
		Monitors:   datadogV1.NewMonitorsApi(apiClient),
		Ctx:        cfgCtx,
	}
}
