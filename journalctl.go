package journalctl

import (
	"net/http"
)

//go:generate gojson -o entry.go -name Entry -pkg "journalctl" -input json/journalctl_entry.json

const DefaultHost = "http://localhost:19531"

// A Client is used for accessing the journal (through systemd-journal-gatewayd.service)
// Like http.DefaultClient, its zero value (DefaultClient) is a usable client that connects to localhost on the default port (http://localhost:19531)
type Client struct {
	Host string

	http http.Client
}

func (c *Client) host() string {
	if c.Host == "" {
		return DefaultHost
	}
	return c.Host
}
