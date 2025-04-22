package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/mamoss-oss/label-exporter/internal/containers"
)

type DNSTagExporter struct {
	mu      sync.RWMutex
	records []containers.DNSRecord
}

func NewDNSTagExporter() *DNSTagExporter {
	return &DNSTagExporter{}
}

func (exp *DNSTagExporter) Start() {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			dnsrecords, err := containers.GetDockerDNS()
			if err != nil {
				log.Printf("failed fetching DNS records %s\n", err)
				continue
			}

			exp.mu.Lock()
			exp.records = dnsrecords
			exp.mu.Unlock()
		}
	}
}

func (exp *DNSTagExporter) DNSResponseHandler(w http.ResponseWriter, r *http.Request) {
	exp.mu.RLock()
	defer exp.mu.RUnlock()

	jsonData, err := json.Marshal(exp.records)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func main() {

	dnsTagExporter := NewDNSTagExporter()

	go dnsTagExporter.Start()

	http.HandleFunc("/dns", dnsTagExporter.DNSResponseHandler)
	port := ":8080"
	log.Printf("Server is running on http://localhost%s/dns\n", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
