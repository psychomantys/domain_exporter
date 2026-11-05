package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-kit/kit/log/level"
	"github.com/zonedb/zonedb"
)

type rdapResponse struct {
	Events []rdapEvent `json:"events"`
}

type rdapEvent struct {
	EventAction string `json:"eventAction"`
	EventDate   string `json:"eventDate"`
}

var rdapHTTPClient = &http.Client{Timeout: 15 * time.Second}

func lookupRDAP(domain string) (float64, error) {
	endpoints := rdapQueryURLs(domain)
	if len(endpoints) == 0 {
		return -1, fmt.Errorf("no RDAP endpoint for domain: %s", domain)
	}

	var lastErr error
	for _, endpoint := range endpoints {
		date, err := fetchRDAPExpiration(domain, endpoint)
		if err == nil {
			return date, nil
		}
		lastErr = err
		level.Warn(logger).Log("msg", "RDAP lookup failed", "domain", domain, "url", endpoint, "err", err)
	}
	return -1, lastErr
}

func rdapQueryURLs(domain string) []string {
	zone := zonedb.PublicZone(domain)
	if zone == nil {
		return nil
	}

	var endpoints []string
	for _, base := range zone.RDAPURLs() {
		endpoint, err := url.JoinPath(base, "domain", domain)
		if err != nil {
			continue
		}
		endpoints = append(endpoints, endpoint)
	}
	return endpoints
}

func fetchRDAPExpiration(domain, endpoint string) (float64, error) {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return -1, err
	}
	req.Header.Set("Accept", "application/rdap+json, application/json")

	resp, err := rdapHTTPClient.Do(req)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return -1, err
	}
	if resp.StatusCode != http.StatusOK {
		return -1, fmt.Errorf("RDAP HTTP %d for %s", resp.StatusCode, domain)
	}

	return parseRDAP(domain, body)
}

func parseRDAP(domain string, body []byte) (float64, error) {
	var res rdapResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return -1, err
	}

	for _, event := range res.Events {
		if !strings.EqualFold(event.EventAction, "expiration") {
			continue
		}

		dateStr := strings.TrimSpace(event.EventDate)
		layouts := append([]string{time.RFC3339Nano, time.RFC3339}, formats...)
		for _, layout := range layouts {
			if date, err := time.Parse(layout, dateStr); err == nil {
				level.Info(logger).Log("domain:", domain, "date", date, "source", "rdap")
				return float64(date.Unix()), nil
			}
		}
		return -1, fmt.Errorf("Unable to parse RDAP date: %s, for %s", dateStr, domain)
	}

	return -1, fmt.Errorf("no RDAP expiration event for domain: %s", domain)
}
