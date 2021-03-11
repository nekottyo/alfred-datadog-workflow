package main

import (
	"flag"
	"log"
	"os"

	"github.com/nekottyo/alfred-datadog-workflow/pkg/dd"
	"gopkg.in/zorkian/go-datadog-api.v2"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/keychain"
	"go.deanishe.net/fuzzy"
)

const (
	apiKeyName = "apikey"
	appKeyName = "appkey"
)

var (
	maxResults = 200

	// Command-line flags
	flagAPIKey    bool
	flagAPPKey    bool
	flagDashboard bool
	flagMonitor   bool
	flagService   bool

	// credentials
	apiKey string
	appKey string

	arg string

	// Workflow
	sopts  []fuzzy.Option
	wf     *aw.Workflow
	kc     *keychain.Keychain
	client *datadog.Client
)

func init() {
	flag.BoolVar(&flagAPIKey, "apikey", false, "register apikey")
	flag.BoolVar(&flagAPPKey, "appkey", false, "register appkey")
	flag.BoolVar(&flagDashboard, "dashboard", false, "list dashboard")
	flag.BoolVar(&flagMonitor, "monitor", false, "list monitor")
	flag.BoolVar(&flagService, "service", false, "list service")

	initWorkflow()
}

func initWorkflow() {
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
}

func getSecrets(kc *keychain.Keychain) (string, string, error) {
	var apiKey, appKey string
	var err error

	apiKey, err = kc.Get(apiKeyName)
	if err != nil {
		return apiKey, appKey, err
	}
	appKey, err = kc.Get(appKeyName)
	if err != nil {
		return apiKey, appKey, err
	}

	return apiKey, appKey, nil
}

func setupDatadogClient() {

	var err error
	apiKey, appKey, err = getSecrets(kc)
	if err != nil {
		wf.FatalError(err)
	}
	client = datadog.NewClient(apiKey, appKey)

}

func run() {
	wf.Args()
	flag.Parse()

	arg = flag.Arg(0)

	if flagAPIKey {
		if err := kc.Set(apiKeyName, arg); err != nil {
			wf.FatalError(err)
		}
		wf.SendFeedback()
		os.Exit(0)
	}

	if flagAPPKey {
		if err := kc.Set(appKeyName, arg); err != nil {
			wf.FatalError(err)
		}
		wf.SendFeedback()
		os.Exit(0)
	}

	setupDatadogClient()

	if flagDashboard {
		d := dd.NewBoard(client, wf)
		if err := d.ListBoards(); err != nil {
			wf.FatalError(err)
		}
	}
	if flagMonitor {
		d := dd.NewMonitor(client, wf)
		if err := d.ListMonitors(); err != nil {
			wf.FatalError(err)
		}
	}
	if flagService {
		d, err := dd.NewServices("config/service.yaml", wf)
		if err != nil {
			wf.FatalError(err)
		}
		if err := d.ListServices(); err != nil {
			wf.FatalError(err)
		}
	}
	if arg != "" {
		wf.Filter(arg)
	}
	log.Printf("[main] query=%s", arg)

	wf.WarnEmpty("No matching", "Try a different query")
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
