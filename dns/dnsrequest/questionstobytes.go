package dnsrequest

import (
	"bytes"
	"encoding/binary"
)

func QuestionToBytes(question *DNSQuestion) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, err := buf.Write(question.Name)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, question.Type)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, question.Class)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
