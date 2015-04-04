package journalctl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_GetEntries(t *testing.T) {

	var expected = []Entry{
		// currently any fields which systemd prepends with "_" or "__" will be hidden
		Entry{
			CURSOR:                  "s=c60ae6415a0e4531a2064c8cdb11a0f4;i=587;b=d5f050252af342a2817795b375889f78;m=b11f9d6;t=5120f62678669;x=74d4678ad474a6b6",
			REALTIMETIMESTAMP:       "1427232168314473",
			MONOTONICTIMESTAMP:      "185727446",
			BOOTID:                  "d5f050252af342a2817795b375889f78",
			UID:                     "0",
			MACHINEID:               "fe39ba83b9244251b1704fc655fbff2f",
			PRIORITY:                "5",
			CAPEFFECTIVE:            "3fffffffff",
			TRANSPORT:               "syslog",
			HOSTNAME:                "ip-172-30-0-229",
			SYSLOGFACILITY:          "10",
			GID:                     "1000",
			SYSTEMDOWNERUID:         "1000",
			SYSTEMDSLICE:            "user-1000.slice",
			AUDITLOGINUID:           "1000",
			SYSLOGIDENTIFIER:        "sudo",
			COMM:                    "sudo",
			EXE:                     "/usr/bin/sudo",
			SELINUXCONTEXT:          "unconfined_u:unconfined_r:unconfined_t:s0-s0:c0.c1023",
			AUDITSESSION:            "3",
			SYSTEMDCGROUP:           "/user.slice/user-1000.slice/session-3.scope",
			SYSTEMDSESSION:          "3",
			SYSTEMDUNIT:             "session-3.scope",
			MESSAGE:                 "fedora : TTY=pts/1 ; PWD=/home/fedora ; USER=root ; COMMAND=/bin/systemctl status systemd-journal-gatewayd",
			PID:                     "1125",
			CMDLINE:                 "sudo systemctl status systemd-journal-gatewayd",
			SOURCEREALTIMETIMESTAMP: "1427232168313982",
		},
	}

	bts, err := ioutil.ReadFile("json/journalctl_entry.json")
	if err != nil {
		t.Error(err)
		return
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, string(bts))
	}))
	defer ts.Close()

	client := &Client{Host: ts.URL}
	entries_ch, err := client.Entries(nil)
	if err != nil {
		t.Error(err)
		return
	}

    var entries []Entry
    for entry := range entries_ch{
        entries = append(entries, entry)
    }



	if !reflect.DeepEqual(expected, entries) {
		t.Errorf("Expected and actual entries differ:\n%+v\n%+v", expected, entries)
	}
}

func Test_GetEntriesList(t *testing.T) {
	bts, err := ioutil.ReadFile("json/journalctl_entry_list.json")
	if err != nil {
		t.Error(err)
		return
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, string(bts))
	}))
	defer ts.Close()
	client := &Client{Host: ts.URL}
	entries_ch, err := client.Entries(nil)
	if err != nil {
		t.Error(err)
		return
    }
	

    var entries []Entry
    for entry := range entries_ch{
        entries = append(entries, entry)
    }
	if len(entries) != 2 {
		t.Errorf("Expected 2 entries and found %d", len(entries))
		return
	}
    
    var i int
	for entry := range entries {
		if reflect.DeepEqual(entry, Entry{}) {
			t.Errorf("Received empty entry for entry %d", i)
		}
        i++
	}
}
