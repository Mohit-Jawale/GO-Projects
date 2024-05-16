package dnsresponse

import (
	"DNS/dnsrequest"
	"bytes"
	"encoding/binary"
)

func ParseQuestion(reader *bytes.Reader) (*dnsrequest.DNSQuestion, error) {
	name, err := DecodeName(reader)
	if err != nil {
		return nil, err
	}

	var type_, class_ uint16
	err = binary.Read(reader, binary.BigEndian, &type_)
	if err != nil {
		return nil, err
	}
	err = binary.Read(reader, binary.BigEndian, &class_)
	if err != nil {
		return nil, err
	}

	return &dnsrequest.DNSQuestion{
		Name:  name,
		Type:  type_,
		Class: class_,
	}, nil
}
