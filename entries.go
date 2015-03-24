package journalctl

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
)

const (
	_GET  = "GET"
	_POST = "POST"
)

func (c *Client) Entries() (entries []Entry, err error) {
	const endpoint = "/entries"
	host := c.host()
	req, err := http.NewRequest(_GET, host+"/entries", nil)
	if err != nil {
		return
	}

	req.Header.Add("Accept", "application/json")
	resp, err := c.http.Do(req)
	if err != nil {
		return
	}

	// Defaults to scanning lines
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		var entry Entry
		bts := scanner.Bytes()
		if len(bts) == 0 {
			continue
		}
		err = json.Unmarshal(bts, &entry)
		if err != nil {
			return entries, err
		}
		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		return entries, err
	}
	return entries, nil
}
