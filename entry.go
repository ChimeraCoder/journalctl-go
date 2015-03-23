package journalctl

type Entry struct {
	MESSAGE                   string `json:"MESSAGE"`
	PRIORITY                  string `json:"PRIORITY"`
	SYSLOGFACILITY            string `json:"SYSLOG_FACILITY"`
	SYSLOGIDENTIFIER          string `json:"SYSLOG_IDENTIFIER"`
	_BOOTID                   string `json:"_BOOT_ID"`
	_HOSTNAME                 string `json:"_HOSTNAME"`
	_MACHINEID                string `json:"_MACHINE_ID"`
	_SOURCEMONOTONICTIMESTAMP string `json:"_SOURCE_MONOTONIC_TIMESTAMP"`
	_TRANSPORT                string `json:"_TRANSPORT"`
	_CURSOR                   string `json:"__CURSOR"`
	_MONOTONICTIMESTAMP       string `json:"__MONOTONIC_TIMESTAMP"`
	_REALTIMETIMESTAMP        string `json:"__REALTIME_TIMESTAMP"`
}
