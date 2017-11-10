package mutators

import (
	"encoding/json"
	"time"

	"github.com/clearbit/tinysyslog/util"
)

// JSONMutator represents a JSON mutator
type JSONMutator struct{}

// NewJSONMutator creates a JSONMutator instance
func NewJSONMutator() Mutator {
	return Mutator(&JSONMutator{})
}

// HAproxy JSON
type HAProxy struct {
	Pid            int    `json:"pid"`
	Actconn        int    `json:"actconn"`
	Feconn         int    `json:"feconn"`
	Beconn         int    `json:"beconn"`
	BackendQueue   int    `json:"backend_queue"`
	SrvConn        int    `json:"srv_conn"`
	Retry          int    `json:"retry"`
	Tw             int    `json:"tw"`
	Tc             int    `json:"tc"`
	Tt             string `json:"tt"`
	Tsc            string `json:"tsc"`
	ClientAddr     string `json:"client_addr"`
	ClientPort     int    `json:"client_port"`
	FrontAddr      string `json:"front_addr"`
	FrontPort      int    `json:"front_port"`
	FrontTransport string `json:"front_transport"`
	BackName       string `json:"back_name"`
	BackServer     string `json:"back_server"`
	BytesUploaded  int    `json:"bytes_uploaded"`
	BytesRead      string `json:"bytes_read"`
}

// Mutate mutates a slice of bytes
func (jm *JSONMutator) Mutate(logParts map[string]interface{}) string {
	t := logParts["timestamp"].(time.Time)
	m := map[string]interface{}{
		"timestamp": t.Format(time.RFC3339Nano),
		"hostname":  logParts["hostname"].(string),
		"app_name":  logParts["app_name"].(string),
		"proc_id":   logParts["proc_id"].(string),
		"severity":  util.SeverityNumToString(logParts["severity"].(int)),
	}
	// merge processed message from haproxy
	haproxy := &HAProxy{}
	if err := json.Unmarshal([]byte(logParts["message"].(string)), &haproxy); err != nil {
		m["message"] = logParts["message"]
		m["error"] = err
	} else {
		m["message"] = haproxy
	}
	formatted, _ := json.Marshal(m)
	return string(formatted)
}
