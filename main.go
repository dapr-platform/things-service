package main

import (
	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"things-service/api"
	_ "things-service/docs"
	"things-service/eventsub"
	_ "things-service/prom"
	_ "things-service/service"
)

var (
	PORT = 80
)

func init() {

	if val := os.Getenv("LISTEN_PORT"); val != "" {
		PORT, _ = strconv.Atoi(val)
	}
	log.Println("use PORT ", PORT)

	go func() {
		if err := http.ListenAndServe(":81", nil); err != nil {
			log.Println(err)
		}
	}()
}

// @title things-service RESTful API
// @version 1.0
// @description things-service API 文档.
// @BasePath /swagger/things-service
func main() {
	mux := chi.NewRouter()
	api.InitRoute(mux)
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/swagger*", httpSwagger.WrapHandler)

	s := daprd.NewServiceWithMux(":"+strconv.Itoa(PORT), mux)
	eventsub.Sub(s)
	log.Println("server start")
	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error: %v", err)
	}
}
