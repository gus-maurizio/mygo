package types

import (
	"time"
)

// Context is designed to store relevant information for observability and
// tracing that will be needed to identify what is going on.
type Context struct {
	UserId        string  `json:"userid"`
	AccountId     string  `json:"accountid"`
	IPaddress     string  `json:"ipaddress"`
	ApplicationId string  `json:"applicationid"`
	ModuleId      string  `json:"moduleid"`
	VersionId     string  `json:"versionid"`
	EnvironmentId string  `json:"environmentid"`
	ComputeId     string  `json:"computeid"`
	RegionId      string  `json:"regionid"`
	ZoneId        string  `json:"zoneid"`
	TraceId       string  `json:"traceid"`
	RequestId     string  `json:"requestid"`
	ParentId      string  `json:"parentid"`
	Timestamp     float64 `json:"timestamp"`
}

// This is what gets loaded from the -f .yaml configuration file
type Config struct {
	A                string `yaml:"a"`
	DefaultUnit      string `yaml:"defaulttimeunit"`
	DefaultTick      int    `yaml:"defaulttimetick"`
	PrometheusPort   int    `yaml:"prometheusport"`
	PrometheusHandle string `yaml:"prometheushandle"`
	Plugins          []struct {
		PluginName   string `yaml:"pluginname"`
		PluginModule string `yaml:"pluginmodule"`
		PluginUnit   string `yaml:"plugintimeunit"`
		PluginTick   int    `yaml:"plugintimetick"`
	}
}

type PluginRuntime struct {
	Ticker     *time.Ticker
	PluginName string
}

type FuncMeasure func() string

type FuncPlugin func(string, *time.Ticker, FuncMeasure)
