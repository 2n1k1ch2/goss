package goss_agent

import (
	"flag"
	"goss/pkg/config"
	"log"
	"os"
)

var (
	configPath string
	mode       string
	debug      bool
)

type Runner struct {
	logger *log.Logger
	cfg    config.Config
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

	return &Runner{
		logger: logger,
		cfg:    cfg,
	}
}
