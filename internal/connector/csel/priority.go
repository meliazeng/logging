package csel

import (
	syslog "github.com/RackSec/srslog"
)

func GetPriority(kvp map[string]string) syslog.Priority {
	pVal := kvp["ELCS"]
	switch pVal {
	case "0A":
		return syslog.LOG_LOCAL0
	case "0B":
		return syslog.LOG_LOCAL1
	case "0C":
		return syslog.LOG_LOCAL2
	case "0D":
		return syslog.LOG_LOCAL3
	case "0E":
		return syslog.LOG_LOCAL4
	case "0F":
		return syslog.LOG_LOCAL5
	default:
		return syslog.LOG_ERR
	}
}
