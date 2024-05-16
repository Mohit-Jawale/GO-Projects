package dnsrequest

import (
	"bytes"
)

func EncodeDNSName(domainName string) []byte {
	var encoded bytes.Buffer
	parts := bytes.Split([]byte(domainName), []byte("."))

	for _, part := range parts {
		if len(part) > 0 {
			encoded.WriteByte(byte(len(part)))
			encoded.Write(part)
		}
	}
	encoded.WriteByte(0) // to indicate end of name

	return encoded.Bytes()
}
