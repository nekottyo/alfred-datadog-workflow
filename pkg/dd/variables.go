package dd

import "time"

// How long to cache repo list for
var (
	maxCacheAge = 180 * time.Minute
	baseURL     = "https://app.datadoghq.com"
)
