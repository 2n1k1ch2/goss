package goss_agent

import (
	"flag"
	"goss/pkg/alerting"
	"goss/pkg/config"
	"goss/pkg/exporter/http"
	"goss/pkg/store"
	"log"
	"os"
)

var (
	configPath string
	mode       string
	debug      bool
)

type Runner struct {
	logger  *log.Logger
	cfg     config.Config
	alerter *alerting.Alerter
	server  *http.Server
	store   *store.Store
}

func init() {
	flag.StringVar(&configPath, "config", "config.yaml", "path to config file")
	flag.StringVar(&mode, "mode", "embedded", "operation mode: embedded or scrape")
	flag.BoolVar(&debug, "debug", false, "enable debug logging")

}

func NewRunner(cfg config.Config) *Runner {
	var logger *log.Logger
	if debug {
		logger = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags|log.Lshortfile)
	} else {
		logger = log.New(os.Stdout, "[INFO] ", log.LstdFlags)
	}
	out := make(chan alerting.Alert)
	return &Runner{
		logger:  logger,
		cfg:     cfg,
		alerter: alerting.NewAlerter(1000, out),
		server:  http.NewServer(),
		store:   store.NewStore(10),
	}
}
