package dnsresponse

import (
	"bytes"
	"encoding/binary"
)

func ParseRecord(reader *bytes.Reader) (*DNSRecord, error) {
	name, err := DecodeName(reader)
	if err != nil {
		return nil, err
	}

	var recordData RecordData

	if err := binary.Read(reader, binary.BigEndian, &recordData); err != nil {
		return nil, err
	}

	data := make([]byte, recordData.DataLen)
	if _, err := reader.Read(data); err != nil {
		return nil, err
	}

	return &DNSRecord{
		Name:  name,
		Type:  recordData.Type,
		Class: recordData.Class,
		TTL:   recordData.TTL,
		Data:  data,
	}, nil
}
