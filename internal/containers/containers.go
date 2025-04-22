package containers

import (
	"fmt"
	"log"
	"net"
)

const (
	defaultDockerSockPath string = "/var/run/docker.sock"
	dnsName               string = "dns.name"
	dnsValue              string = "dns.value"
)

type labels map[string]string

type Container struct {
	Names  []string
	Image  string
	State  string
	Labels map[string]string
}

type Containers []Container

type DNSEntry struct {
	Name string
	IP   net.IPAddr
}

func GetDockerDNS() ([]DNSRecord, error) {
	var dnsEntries []DNSEntry
	containers, err := fetchAllDockerContainers()
	if err != nil {
		return []DNSRecord{}, fmt.Errorf("failed fetching containers")
	}

	for _, c := range containers {
		dns, err := extract_dns_from_container(c)
		if err != nil {
			log.Printf("error extracting dns from container tags %s\n", err)
			continue
		}
		dnsEntries = append(dnsEntries, dns)
	}

	dnsRecords := make([]DNSRecord, len(dnsEntries))
	for i, e := range dnsEntries {
		dnsRecords[i] = convert_dns_to_json(e)
	}

	return dnsRecords, nil
}
