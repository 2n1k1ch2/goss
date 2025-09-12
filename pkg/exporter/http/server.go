package http

import (
	"encoding/json"
	"goss/pkg/exporter/exporter"
	"log"
	"net/http"
)

type Server struct {
	mux      *http.ServeMux
	exporter exporter.Exporter
}

func NewServer() *Server {
	return &Server{mux: http.NewServeMux()}
}

func (s *Server) Start() {
	err := http.ListenAndServe(":8080", s.mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)

	}
	s.mux.HandleFunc("/debug/goss/clusters", func(w http.ResponseWriter, r *http.Request) {
		metricsHandler(w, r, &s.exporter)
	})

}
func metricsHandler(w http.ResponseWriter, r *http.Request, exporter *exporter.Exporter) {
	data, err := json.Marshal(exporter.Cluster)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(data); err != nil {
		log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}
