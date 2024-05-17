package dnsresolver

import (
	"DNS/constants"
	"DNS/dnspacket"
	"DNS/dnsresponse"
)

func GetNameserver(records dnspacket.DNSPacket) string {

	for _, record := range records.Authorities {
		if record.Type == uint16(constants.TYPE_NS) {

			return dnsresponse.DecodeNSName(record.Data)

		} else if record.Type == uint16(constants.TYPE_CNAME) {

			return dnsresponse.DecodeNSName(record.Data)

		}

	}
	return ""
}
