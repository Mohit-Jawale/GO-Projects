package dnspacket

import (
	"DNS/dnsrequest"
	"DNS/dnsresponse"
)

type DNSPacket struct {
	Header      dnsrequest.DNSHeader
	Questions   []dnsrequest.DNSQuestion
	Answers     []dnsresponse.DNSRecord
	Authorities []dnsresponse.DNSRecord
	Additionals []dnsresponse.DNSRecord
}
