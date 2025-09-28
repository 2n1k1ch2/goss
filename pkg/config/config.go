package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"time"
)

type Config struct {
	Mode               string        `yaml:"mode"`
	SampleInterval     time.Duration `yaml:"sample_interval"`
	HttpBind           string        `yaml:"http_bind"`
	PrometheusEnabled  bool          `yaml:"prometheus_enabled"`
	PrometheusBind     int           `yaml:"prometheus_bind"`
	RetentionSnapshots int           `yaml:"retention_snapshots"`
	RetentionWindow    time.Duration `yaml:"retention_window"`
	ReleaseEnv         []string      `yaml:"release_env"`

	CloudEnabled bool   `yaml:"cloud_enabled"`
	CloudURL     string `yaml:"cloud_url"`
	CloudAuth    string `yaml:"cloud_auth"`
	AgentID      string `yaml:"agent_id"`
	ReleaseTag   string `yaml:"release_tag"`

	RedactionEnabled bool `yaml:"redaction_enabled"`
}

func Load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	if err := validate(cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func DefaultConfig() *Config {
	return &Config{
		Mode:               "embedded",
		SampleInterval:     time.Minute,
		HttpBind:           "0.0.0.0:8080",
		PrometheusEnabled:  false,
		PrometheusBind:     0,
		RetentionSnapshots: 30,
		RetentionWindow:    time.Hour,
		ReleaseEnv:         []string{},
		CloudEnabled:       false,
		CloudURL:           "",
		CloudAuth:          "",
		AgentID:            "",
		ReleaseTag:         "",
		RedactionEnabled:   false,
	}
}

func validate(config Config) error {
	if config.Mode == "" {
		return errors.New("mode is required")
	}
	if config.HttpBind == "" {
		return errors.New("http_bind is required")
	}
	if config.SampleInterval < time.Second {
		return errors.New("sample_interval is too short")
	}
	if config.PrometheusEnabled && config.PrometheusBind == 0 {
		return errors.New("prometheus_bind is required when Prometheus is enabled")
	}
	if config.RetentionWindow < time.Second {
		return errors.New("retention_window is too short")
	}
	if config.CloudEnabled && config.CloudURL == "" && config.CloudAuth == "" {
		return errors.New("either cloud_url or cloud_auth must be provided when Cloud is enabled")
	}
	return nil
}
