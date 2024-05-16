package dnsresponse

import (
	"DNS/dnsrequest"
	"encoding/binary"
	"io"
)

func ParseHeader(reader io.Reader) (*dnsrequest.DNSHeader, error) {
	var header dnsrequest.DNSHeader
	err := binary.Read(reader, binary.BigEndian, &header)
	if err != nil {
		return nil, err
	}
	return &header, nil
}
