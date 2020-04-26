package main

import (
	"flag"
	"log"

	"github.com/nekottyo/alfred-datadog-workflow/pkg/dd"
	"gopkg.in/zorkian/go-datadog-api.v2"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/fuzzy"
	"github.com/deanishe/awgo/keychain"
)

const (
	apiKeyName = "apikey"
	appKeyName = "appkey"
)

var (
	maxResults = 200

	// Command-line flags
	apikey      string
	appkey      string
	doDashboard bool
	doMonitor   bool
	query       string

	// Workflow
	sopts  []fuzzy.Option
	wf     *aw.Workflow
	kc     *keychain.Keychain
	client *datadog.Client
)

func init() {
	flag.BoolVar(&doDashboard, "dashboard", false, "list dashboard")
	flag.BoolVar(&doMonitor, "monitor", false, "list monitor")

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
		query = args[0]

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
		d := dd.NewDashboard(client, wf)
		if err := d.ListDashboards(); err != nil {
			wf.FatalError(err)
		}
	}
	if doMonitor {
		d := dd.NewMonitor(client, wf)
		if err := d.ListMonitors(); err != nil {
			wf.FatalError(err)
		}
	}
	if query != "" {
		wf.Filter(query)
	}
	log.Printf("[main] query=%s", query)

	wf.WarnEmpty("No matching", "Try a different query")
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
