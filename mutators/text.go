package mutators

import (
	"fmt"
	"time"
)

// TextMutator represents a text mutator
type TextMutator struct{
	syslogFormat string
}

// NewTextMutator creates TextMutator instance
func NewTextMutator(syslogFormat string) Mutator {
	return Mutator(&TextMutator{syslogFormat: syslogFormat})
}

// Mutate mutates a slice of bytes
func (tm *TextMutator) Mutate(logParts map[string]interface{}) string {
	t := logParts["timestamp"].(time.Time)
	formatted := ""
	// will eventually need to support user-defined format
	if tm.syslogFormat == "5424" {
		formatted = fmt.Sprintf("%s %s %s[%s]: %s",
			t.Format("Jan _2 15:04:05"),
			logParts["hostname"],
			logParts["app_name"],
			logParts["proc_id"],
			logParts["message"])
	} else {
		formatted = fmt.Sprintf("%s %s %s: %s",
			t.Format("Jan _2 15:04:05"),
			logParts["hostname"],
			logParts["tag"],
			logParts["content"])
	}

	return formatted
}
