package dd

import "time"

// How long to cache repo list for
var (
	maxCacheAge = 5 * time.Minute
	baseURL     = "https://app.datadoghq.com"
)
