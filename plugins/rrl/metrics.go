package rrl

import (
	"github.com/coredns/coredns/plugin"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	subsystem = "rrl"
	// droppedCount is counter of successfully filtered queries.
	droppedCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: subsystem,
		Name:      "dropped_total",
		Help:      "Counter of requests that was dropped.",
	}, []string{"server"})
)
