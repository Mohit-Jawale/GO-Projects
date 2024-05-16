package dnsquerybuilder

import (
	"DNS/dnsrequest"
	"math/rand"
)

func BuildQuery(domainName string, recordType uint16) ([]byte, error) {

	id := uint16(rand.Intn(65536))
	name := dnsrequest.EncodeDNSName(domainName)

	header := dnsrequest.DNSHeader{
		ID:           id,
		Flags:        0,
		NumQuestions: 1,
	}
	question := dnsrequest.DNSQuestion{
		Name:  name,
		Type:  recordType,
		Class: 1,
	}

	headerBytes, err := dnsrequest.HeaderToBytes(&header)
	if err != nil {
		return nil, err
	}
	questionBytes, err := dnsrequest.QuestionToBytes(&question)
	if err != nil {
		return nil, err
	}

	return append(headerBytes, questionBytes...), nil

}
