package containers

import (
	"fmt"
	"net"
)

type DNSRecord struct {
	Name string
	IP   string
}

func convert_dns_to_json(dns DNSEntry) DNSRecord {
	return DNSRecord{
		Name: dns.Name,
		IP:   dns.IP.String(),
	}
}

// extract_dns_from_container parses the container labels to build dns entries. Validates IP address correctness.
func extract_dns_from_container(container Container) (DNSEntry, error) {
	var dnsEntry DNSEntry

	name, ok := container.Labels[dnsName]
	if ok {
		dnsEntry.Name = name
	} else {
		return dnsEntry, fmt.Errorf("no tag %s found.", dnsName)
	}

	ip, ok := container.Labels[dnsValue]
	if ok {
		ipaddr, err := net.ResolveIPAddr("ip", ip)
		if err != nil {
			return dnsEntry, fmt.Errorf("failed parsing IP address in %s for container %s", dnsValue, container.Names[0])
		}
		dnsEntry.IP = *ipaddr
	} else {
		return dnsEntry, fmt.Errorf("no tag %s found.", dnsValue)
	}
	return dnsEntry, nil
}
