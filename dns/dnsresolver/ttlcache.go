package dnsresolver

import (
	"encoding/json"
	"io"
	"os"
	"sync"
	"time"
)

type DNSRecord struct {
	IPAddress  string
	Expiration time.Time
}

type DNSCache struct {
	cache    map[string]DNSRecord
	mutex    sync.RWMutex
	filePath string
}

func NewDNSCache(filePath string) *DNSCache {
	return &DNSCache{
		cache:    make(map[string]DNSRecord),
		filePath: filePath,
	}
}

func (d *DNSCache) Add(domain string, recordType string, ip string, ttl time.Duration) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	key := domain + "." + recordType
	d.cache[key] = DNSRecord{
		IPAddress:  ip,
		Expiration: time.Now().Add(ttl),
	}
	d.SaveToDisk()
}

func (d *DNSCache) Get(domain string, recordType string) string {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	key := domain + "." + recordType
	record, found := d.cache[key]
	if !found {
		return ""
	}

	if time.Now().After(record.Expiration) {
		d.mutex.Lock()
		delete(d.cache, key)
		d.mutex.Unlock()
		d.SaveToDisk()
		return ""
	}

	return record.IPAddress
}

func (d *DNSCache) SaveToDisk() {
	data, err := json.Marshal(d.cache)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(d.filePath, data, 0644)
	if err != nil {
		panic(err)
	}
}

func (d *DNSCache) LoadFromDisk() {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	file, err := os.Open(d.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		panic(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	if len(data) == 0 {
		// Empty file, nothing to load
		return
	}

	err = json.Unmarshal(data, &d.cache)
	if err != nil {
		panic(err)
	}
}
