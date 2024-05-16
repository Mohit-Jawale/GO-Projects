package dnsrequest

type DNSQuestion struct {
	Name  []byte
	Type  uint16
	Class uint16
}
