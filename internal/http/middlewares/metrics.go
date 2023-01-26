package middlewares

import (
	"reflect"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/AliRasoulinejad/cryptos-backend/internal/app"
)

// Config responsible to configure middleware
type Config struct {
	Namespace string
	Buckets   []float64
	Subsystem string
}

const (
	httpRequestsCount    = "requests_total"
	httpRequestsDuration = "request_duration_seconds"
	notFoundPath         = "/not-found"
)

var buckets = []float64{
	0.0005,
	0.001, // 1ms
	0.002,
	0.005,
	0.01, // 10ms
	0.02,
	0.05,
	0.1, // 100 ms
	0.2,
	0.5,
	1.0, // 1s
	2.0,
	5.0,
	10.0, // 10s
	15.0,
	20.0,
	30.0,
}

// DefaultConfig has the default instrumentation config
var DefaultConfig = Config{
	Namespace: app.AppName,
	Subsystem: "http",
	Buckets:   buckets,
}

func isNotFoundHandler(handler echo.HandlerFunc) bool {
	return reflect.ValueOf(handler).Pointer() == reflect.ValueOf(echo.NotFoundHandler).Pointer()
}

func MetricsMiddleware() func(next echo.HandlerFunc) echo.HandlerFunc {
	config := &DefaultConfig

	httpRequests := promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: config.Namespace,
		Subsystem: config.Subsystem,
		Name:      httpRequestsCount,
		Help:      "HTTP Request count",
	}, []string{"status", "method", "handler"})

	httpDuration := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: config.Namespace,
		Subsystem: config.Subsystem,
		Name:      httpRequestsDuration,
		Help:      "Duration of each request",
		Buckets:   config.Buckets,
	}, []string{"method", "handler"})

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			path := c.Path()

			if isNotFoundHandler(c.Handler()) {
				path = notFoundPath
			}

			timer := prometheus.NewTimer(httpDuration.WithLabelValues(req.Method, path))
			err := next(c)
			timer.ObserveDuration()

			if err != nil {
				c.Error(err)
			}

			status := strconv.Itoa(c.Response().Status)
			httpRequests.WithLabelValues(status, req.Method, path).Inc()

			return err
		}
	}
}
