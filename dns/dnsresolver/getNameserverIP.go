package dnsresolver

import (
	"DNS/dnspacket"
	"net"
)

func GetNameserverIP(records dnspacket.DNSPacket, recordType uint16) string {
	for _, record := range records.Additionals {
		if record.Type == recordType {
			return net.IP(record.Data).String()
		}

	}
	return ""
}
