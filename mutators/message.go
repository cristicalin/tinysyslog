package mutators

import (
	"fmt"
)

// MessageMutator represents a message mutator
type MessageMutator struct{
	syslogFormat string
}

// NewMessageMutator creates MessageMutator instance
func NewMessageMutator(syslogFormat string) Mutator {
	return Mutator(&MessageMutator{syslogFormat: syslogFormat})
}

// Mutate mutates a slice of bytes
func (tm *MessageMutator) Mutate(logParts map[string]interface{}) string {
	if tm.syslogFormat == "5424" {
		return fmt.Sprintf("%s", logParts["message"])
	} else {
		return fmt.Sprintf("%s", logParts["content"])
	}
}
