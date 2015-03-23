package journalctl

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	_GET  = "GET"
	_POST = "POST"
)

func (c *Client) Entries() (entries []Entry, err error) {
	host := c.host()
	req, err := http.NewRequest(_GET, host, nil)
	if err != nil {
		return
	}

	req.Header.Add("Accept", "application/json")
	resp, err := c.http.Do(req)
	if err != nil {
		return
	}
	dec := json.NewDecoder(resp.Body)

	for {
		var entry Entry
		if err = dec.Decode(&entry); err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return
		}
		entries = append(entries, entry)
	}
	return entries, nil

}
