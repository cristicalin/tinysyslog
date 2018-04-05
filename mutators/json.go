package mutators

import (
	"strings"
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
	// merge processed message
	d := json.NewDecoder(strings.NewReader(logParts["message"].(string)))
	d.UseNumber()
	var message interface{}
	if err := d.Decode(&message); err != nil {
		m["message"] = logParts["message"]
		m["error"] = err
	} else {
		m["message"] = message
	}
	formatted, _ := json.Marshal(m)
	return string(formatted)
}
