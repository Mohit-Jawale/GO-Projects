package dnsresponse

type DNSRecord struct {
	Name  []byte
	Type  uint16
	Class uint16
	TTL   uint32
	Data  []byte
}

type RecordData struct {
	Type    uint16
	Class   uint16
	TTL     uint32
	DataLen uint16
}
