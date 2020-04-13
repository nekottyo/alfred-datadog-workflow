package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/nekottyo/alfred-datadog-workflow/pkg/dd"
	"gopkg.in/zorkian/go-datadog-api.v2"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/fuzzy"
	"github.com/deanishe/awgo/keychain"
)

const (
	appName            = "alfred-datadog-workflow"
	apiKeyName         = "apikey"
	appKeyName         = "appkey"
	dashboardCacheName = "dashboard.json"
)

var (
	Version     = "0.0.1"
	maxResults  = 200
	minScore    = 10.0
	maxCacheAge = 180 * time.Minute // How long to cache repo list for

	// Command-line flags
	apikey      string
	appkey      string
	doDashboard bool
	query       string

	// Workflow
	sopts  []fuzzy.Option
	wf     *aw.Workflow
	kc     *keychain.Keychain
	client *datadog.Client
)

func init() {
	flag.BoolVar(&doDashboard, "dashboard", false, "list dashboard")

	sopts = []fuzzy.Option{
		fuzzy.AdjacencyBonus(10.0),
		fuzzy.LeadingLetterPenalty(-0.1),
		fuzzy.MaxLeadingLetterPenalty(-3.0),
		fuzzy.UnmatchedLetterPenalty(-0.5),
	}

	wf = aw.New(aw.HelpURL("https://github.com/nekottyo/alfred-datadog-workflow"),
		aw.MaxResults(maxResults),
		aw.SortOptions(sopts...))

	kc = keychain.New("net.nekottyo.alfred-datadog.workflow")
	apikey, appkey = initSecrets(kc)
	client = datadog.NewClient(apikey, appkey)
}

func initSecrets(kc *keychain.Keychain) (string, string) {
	var apikey, appkey string

	apikey, err := kc.Get(apiKeyName)
	if err != nil {
		wf.FatalError(err)
	}
	appkey, _ = kc.Get(appKeyName)
	return apikey, appkey
}

func run() {
	wf.Args()
	flag.Parse()

	if args := flag.Args(); len(args) > 0 {
		command := args[0]
		if len(args) > 1 {
			query = args[1]
		}
		switch {
		case command == apiKeyName:
			if err := kc.Set(apiKeyName, args[1]); err != nil {
				wf.FatalError(err)
			}
		case command == appKeyName:
			if err := kc.Set(appKeyName, args[1]); err != nil {
				wf.FatalError(err)
			}
		}
	}
	if doDashboard {
		var dashboards []datadog.DashboardLite
		log.Printf("[main] cache dir %s", wf.Cache.Dir)

		if wf.Cache.Exists(dashboardCacheName) {
			if err := wf.Cache.LoadJSON(dashboardCacheName, &dashboards); err != nil {
				wf.FatalError(err)
			}
		}
		if wf.Cache.Expired(dashboardCacheName, maxCacheAge) {
			var err error
			dashboards, err = dd.QueryDashboard(client)
			if err != nil {
				wf.FatalError(err)
			}
			wf.Cache.StoreJSON(dashboardCacheName, dashboards)
		}

		for _, dash := range dashboards {
			url := fmt.Sprintf("https://app.datadoghq.com/dash/%d/datadog", dash.GetId())
			wf.NewItem(dash.GetTitle()).
				Subtitle(url).
				Arg(url).
				UID(dash.GetTitle()).
				Valid(true)
		}

		if query != "" {
			res := wf.Filter(query)
			log.Printf("[main] %d/%d match \"%s\"", len(dashboards), len(res), query)
		}
	}

	log.Printf("[main] query=%s", query)

	wf.WarnEmpty("No matching", "Try a different query")
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
