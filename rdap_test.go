package main

import (
	"io/ioutil"
	"path"
	"testing"
	"time"
)

func TestParseRDAP(t *testing.T) {
	cases := []struct {
		filename string
		domain   string
		date     time.Time
	}{
		{
			filename: "weni.ai.rdap.json",
			domain:   "weni.ai",
			date:     time.Date(2028, 10, 31, 5, 47, 30, 0, time.UTC),
		},
	}

	for i := range cases {
		body, err := ioutil.ReadFile(path.Join("testdata", cases[i].filename))
		if err != nil {
			t.Fatalf("read %s: %v", cases[i].filename, err)
		}

		parsedDate, err := parseRDAP(cases[i].domain, body)
		if err != nil {
			t.Fatalf("%s: %v", cases[i].filename, err)
		}

		want := float64(cases[i].date.Unix())
		if parsedDate != want {
			t.Errorf("cases[%d]: got=%.0f want=%.0f", i, parsedDate, want)
		}
	}
}

func TestParseRDAPFractionalSeconds(t *testing.T) {
	body := []byte(`{"events":[{"eventAction":"expiration","eventDate":"2027-02-05T20:13:47.38Z"}]}`)
	got, err := parseRDAP("suiteshare.io", body)
	if err != nil {
		t.Fatal(err)
	}
	want := float64(time.Date(2027, 2, 5, 20, 13, 47, 380000000, time.UTC).Unix())
	if got != want {
		t.Fatalf("got=%.0f want=%.0f", got, want)
	}
}

func TestRDAPQueryURLs(t *testing.T) {
	urls := rdapQueryURLs("weni.ai")
	if len(urls) == 0 {
		t.Fatal("expected RDAP URL for .ai")
	}
	want := "https://rdap.identitydigital.services/rdap/domain/weni.ai"
	if urls[0] != want {
		t.Errorf("got %q want %q", urls[0], want)
	}

	urls = rdapQueryURLs("suiteshare.io")
	if len(urls) == 0 {
		t.Fatal("expected RDAP URL for .io")
	}
	want = "https://rdap.identitydigital.services/rdap/domain/suiteshare.io"
	if urls[0] != want {
		t.Errorf("got %q want %q", urls[0], want)
	}
}
