package dnsrequest

import (
	"bytes"
	"encoding/binary"
)

func HeaderToBytes(header *DNSHeader) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, header)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
