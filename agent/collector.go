package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

//Define the metrics we wish to expose
var fooMetric = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "agent_foometric",
	Help: "Shows whether a foo has occurred in our cluster",
})

var messageMetric = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "agent_plugin_ticks",
		Help: "Number of times plugin has executed.",
	},
	[]string{"plugin"},
)

var bytesMetric = prometheus.NewCounterVec(
        prometheus.CounterOpts{
                Name: "agent_bytes_sent",
                Help: "Number of bytes plugin has generated.",
        },
        []string{"plugin"},
)


func init() {
	//Register metrics with prometheus
	prometheus.MustRegister(fooMetric)
	prometheus.MustRegister(messageMetric)
	prometheus.MustRegister(bytesMetric)

	//Set fooMetric to 1
	fooMetric.Set(0)

}
