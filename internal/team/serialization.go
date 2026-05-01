package team

import (
	"encoding/json"
	"fmt"
)

// ToJSON serializes the team to JSON bytes.
func (t *Team) ToJSON() ([]byte, error) {
	return json.Marshal(t)
}

// FromJSON deserializes a team from JSON bytes.
func FromJSON(data []byte) (*Team, error) {
	var t Team
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, fmt.Errorf("invalid team JSON: %w", err)
	}

	if t.Version == "" {
		t.Version = "1.0"
	}

	if t.Version != "1.0" {
		return nil, fmt.Errorf("unsupported team version: %q", t.Version)
	}

	return &t, nil
}
