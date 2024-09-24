package cdncheck

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
)
import "github.com/projectdiscovery/retryabledns"

type Client struct {
	sync.Once
	cdn          *providerScraper
	waf          *providerScraper
	cloud        *providerScraper
	retryAbleDns *retryabledns.Client
}

// noinspection GoUnusedGlobalVariable
var (
	CdnClient *Client
	once      sync.Once
)

func New() *Client {
	client, _ := NewWithOpts(3, []string{})
	return client
}

func Init() {
	once.Do(func() {
		CdnClient = New()
		if err := json.Unmarshal([]byte(data), &generatedData); err != nil {
			panic(fmt.Sprintf("Could not parse cidr data: %s", err))
		}
	})
}

// DefaultResolvers trusted (taken from fastdialer)
var DefaultResolvers = []string{
	"1.1.1.1:53",
	"1.0.0.1:53",
	"8.8.8.8:53",
	"8.8.4.4:53",
}

// NewWithOpts creates cdn check client with custom options
func NewWithOpts(MaxRetries int, resolvers []string) (*Client, error) {
	if MaxRetries <= 0 {
		MaxRetries = 3
	}
	if len(resolvers) == 0 {
		resolvers = DefaultResolvers
	}
	retryAbleDns, err := retryabledns.New(resolvers, MaxRetries)
	if err != nil {
		return nil, err
	}
	client := &Client{
		cdn:          newProviderScraper(generatedData.CDN),
		waf:          newProviderScraper(generatedData.WAF),
		cloud:        newProviderScraper(generatedData.Cloud),
		retryAbleDns: retryAbleDns,
	}
	return client, nil
}

// Check checks if ip belongs to one of CDN, WAF and Cloud . It is generic method for Checkxxx methods
func (c *Client) Check(ip net.IP) (matched bool, value string, itemType string, err error) {
	if matched, value, err = c.cdn.Match(ip); err == nil && matched && value != "" {
		return matched, value, "cdn", nil
	}
	if matched, value, err = c.waf.Match(ip); err == nil && matched && value != "" {
		return matched, value, "waf", nil
	}
	if matched, value, err = c.cloud.Match(ip); err == nil && matched && value != "" {
		return matched, value, "cloud", nil
	}
	return false, "", "", err
}
