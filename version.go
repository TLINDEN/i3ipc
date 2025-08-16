package swayipc

import (
	"encoding/json"
	"fmt"
)

// A version struct holding the sway wm version
type Version struct {
	HumanReadable string `json:"human_readable"`
	Major         int    `json:"major"`
	Minor         int    `json:"minor"`
	Patch         int    `json:"patch"`
}

// Get the sway software version
func (ipc *SwayIPC) GetVersion() (*Version, error) {
	payload, err := ipc.get(GET_VERSION)
	if err != nil {
		return nil, err
	}

	version := &Version{}
	if err := json.Unmarshal(payload.Payload, &version); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return version, nil
}
