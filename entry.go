package journalctl

type Entry struct {
	MESSAGE                 string `json:"MESSAGE"`
	PRIORITY                string `json:"PRIORITY"`
	SYSLOGFACILITY          string `json:"SYSLOG_FACILITY"`
	SYSLOGIDENTIFIER        string `json:"SYSLOG_IDENTIFIER"`
	AUDITLOGINUID           string `json:"_AUDIT_LOGINUID"`
	AUDITSESSION            string `json:"_AUDIT_SESSION"`
	BOOTID                  string `json:"_BOOT_ID"`
	CAPEFFECTIVE            string `json:"_CAP_EFFECTIVE"`
	CMDLINE                 string `json:"_CMDLINE"`
	COMM                    string `json:"_COMM"`
	EXE                     string `json:"_EXE"`
	GID                     string `json:"_GID"`
	HOSTNAME                string `json:"_HOSTNAME"`
	MACHINEID               string `json:"_MACHINE_ID"`
	PID                     string `json:"_PID"`
	SELINUXCONTEXT          string `json:"_SELINUX_CONTEXT"`
	SOURCEREALTIMETIMESTAMP string `json:"_SOURCE_REALTIME_TIMESTAMP"`
	SYSTEMDCGROUP           string `json:"_SYSTEMD_CGROUP"`
	SYSTEMDOWNERUID         string `json:"_SYSTEMD_OWNER_UID"`
	SYSTEMDSESSION          string `json:"_SYSTEMD_SESSION"`
	SYSTEMDSLICE            string `json:"_SYSTEMD_SLICE"`
	SYSTEMDUNIT             string `json:"_SYSTEMD_UNIT"`
	TRANSPORT               string `json:"_TRANSPORT"`
	UID                     string `json:"_UID"`
	CURSOR                  string `json:"__CURSOR"`
	MONOTONICTIMESTAMP      string `json:"__MONOTONIC_TIMESTAMP"`
	REALTIMETIMESTAMP       string `json:"__REALTIME_TIMESTAMP"`
}
