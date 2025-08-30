package http

import (
	"encoding/json"
	"goss/pkg/exporter/prometheus"
	"log"
	"net/http"
)

type Server struct {
	mux      *http.ServeMux
	exporter prometheus.Exporter
}

func NewServer() *Server {
	return &Server{mux: http.NewServeMux()}
}

func (s *Server) Start() {
	http.ListenAndServe(":8080", s.mux)
	s.mux.HandleFunc("/debug/goss/clusters", func(w http.ResponseWriter, r *http.Request) {
		metricsHandler(w, r, &s.exporter)
	})

}
func metricsHandler(w http.ResponseWriter, r *http.Request, exporter *prometheus.Exporter) {
	data, err := json.Marshal(exporter.Cluster)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}
