package dnsresolver

import (
	"DNS/dnspacket"
	"net"
)

func GetAnswer(records dnspacket.DNSPacket) string {
	for _, record := range records.Answers {
		return net.IP(record.Data).String()

	}
	return ""
}
