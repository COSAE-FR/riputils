package arp

import (
	"time"
)

type Entry struct {
	MacAddress string
	IPAddress  string
	Iface      string
	Permanent  bool
}

type ArpTable map[string]Entry

var (
	stop     = make(chan struct{})
	arpCache = &cache{
		table: make(ArpTable),
	}
)

func AutoRefresh(t time.Duration) {
	go func() {
		for {
			select {
			case <-time.After(t):
				arpCache.Refresh()
			case <-stop:
				return
			}
		}
	}()
}

func StopAutoRefresh() {
	stop <- struct{}{}
}

func CacheUpdate() {
	arpCache.Refresh()
}

func CacheLastUpdate() time.Time {
	return arpCache.Updated
}

func CacheUpdateCount() int {
	return arpCache.UpdatedCount
}

// Search looks up the MAC address for an IP address
// in the arp table
func Search(ip string) Entry {
	return arpCache.Search(ip)
}

func ReverseSearch(mac string) []Entry {
	return arpCache.ReverseSearch(mac)
}
