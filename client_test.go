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
		Entry{
			// currently any fields which systemd prepends with "_" or "__" will be hidden

			//_CURSOR:                     "s=d8ba89471ff0454288d0589e3748af43;i=a;b=cb43988a622e43869c28d71431b0c3fc;m=254c539;t=511f5776c88b5;x=db9d099d8146ca84",
			//_REALTIMETIMESTAMP:         "1427120851814581",
			//_MONOTONICTIMESTAMP:        "39109945",
			//_BOOTID:                    "cb43988a622e43869c28d71431b0c3fc",
			PRIORITY: "6",
			//_MACHINEID:                 "fe39ba83b9244251b1704fc655fbff2f",
			//_HOSTNAME:                   "localhost.localdomain",
			//_TRANSPORT:                  "kernel",
			SYSLOGFACILITY:   "0",
			SYSLOGIDENTIFIER: "kernel",
			//_SOURCEMONOTONICTIMESTAMP: "36028765",
			MESSAGE: "SELinux: the above unknown classes and permissions will be allowed",
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
	entries, err := client.Entries()
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(expected, entries) {
		t.Errorf("Expected and actual entries differ:\n%+v\n%+v", expected, entries)
	}
}
