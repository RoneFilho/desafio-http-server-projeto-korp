package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define métricas
var (
	HttpRequestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api_http_request_total",
		Help: "Total number of requests processed by the API",
	}, []string{"path", "status"})

	HttpRequestErrorTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api_http_request_error_total",
		Help: "Total number of errors returned by the API",
	}, []string{"path", "status"})
)

// Registry customizada (sem métricas padrão Golang)
var customRegistry = prometheus.NewRegistry()

type Info struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

// Registrar métricas com registry customizada
func init() {
	customRegistry.MustRegister(HttpRequestTotal, HttpRequestErrorTotal)
}

// Handler de métricas customizadas com registry customizada
func PrometheusHandler() gin.HandlerFunc {
	h := promhttp.HandlerFor(customRegistry, promhttp.HandlerOpts{})
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Middleware para salvar requisições de métricas
func RequestMetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		c.Next()
		status := c.Writer.Status()
		if status < 400 {
			HttpRequestTotal.WithLabelValues(path, strconv.Itoa(status)).Inc()
		} else {
			HttpRequestErrorTotal.WithLabelValues(path, strconv.Itoa(status)).Inc()
		}
	}
}

func getProjetoKorp(c *gin.Context) {
	horario := Info{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.TimeOnly),
	}
	c.IndentedJSON(http.StatusOK, horario)
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Up and running!",
	})
}

func main() {
	router := gin.Default()
	router.Use(RequestMetricsMiddleware())
	router.GET("/metrics", PrometheusHandler())
	router.GET("/projeto-korp", getProjetoKorp)
	router.GET("/health", getHealth)

	router.Run("0.0.0.0:8080")
}
