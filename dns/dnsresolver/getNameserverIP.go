package dnsresolver

import (
	"DNS/dnspacket"
	"net"
)

func GetNameserverIP(records dnspacket.DNSPacket, recordType uint16) (string, string) {
	for _, record := range records.Additionals {
		if record.Type == recordType {

			return net.IP(record.Data).String(), string(record.Name)
		}

	}
	return "", ""
}
