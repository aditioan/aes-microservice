package main

import (
	"github.com/aditioan/encryptService/helpers"
	kitlog "github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
)

func main()  {
	logger := kitlog.NewLogfmtLogger(os.Stderr)
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "encryption",
		Subsystem: "my_service",
		Name: "request_count",
		Help: "number of request received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "encryption",
		Subsystem: "my_system",
		Name: "request_latency_microseconds",
		Help: "Total duration request in microseconds",
	}, fieldKeys)
	var svc helpers.EncryptService
	svc = helpers.EncryptServiceInstance{}
	svc = helpers.LoggingMiddleware{Logger: logger, Next: svc}
	svc = helpers.IntrumentingMiddleware{RequestCount: requestCount, RequestLatency: requestLatency, Next: svc}
	encryptHandler := httptransport.NewServer(helpers.MakeEncryptEndopint(svc), helpers.DecodeEncryptRequest, helpers.EncodeResponse)
	decryptHandler := httptransport.NewServer(helpers.MakeDecrytEndpoint(svc), helpers.DecodeDecryptRequest, helpers.EncodeResponse)
	http.Handle("/encrypt", encryptHandler)
	http.Handle("/decrypt", decryptHandler)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8000", nil))
}
