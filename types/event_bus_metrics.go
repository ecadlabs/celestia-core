package types

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/discard"

	prometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

const (
	// MetricsSubsystem is a subsystem shared by all metrics exposed by this
	// package.
	MetricsSubsystem = "event_bus"
)

type EventBusMetrics struct {
	// Total number of events
	EventsTotal metrics.Counter
}

func PrometheusEventBusMetrics(namespace string, labelsAndValues ...string) *EventBusMetrics {
	labels := []string{}
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}
	return &EventBusMetrics{
		EventsTotal: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "events_total",
			Help:      "Number of events",
		}, append(labels, "event_type")).With(labelsAndValues...),
	}
}

// NopMetrics returns no-op Metrics.
func NopEventBusMetrics() *EventBusMetrics {
	return &EventBusMetrics{
		EventsTotal: discard.NewCounter(),
	}
}
