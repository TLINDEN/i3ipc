package i3ipc

import (
	"encoding/json"
	"fmt"
)

type Version struct {
	HumanReadable string `json:"human_readable"`
	Major         int    `json:"major"`
	Minor         int    `json:"minor"`
	Patch         int    `json:"patch"`
}

func (ipc *I3ipc) GetVersion() (*Version, error) {
	payload, err := ipc.get(GET_VERSION)
	if err != nil {
		return nil, err
	}

	version := &Version{}
	if err := json.Unmarshal(payload, &version); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return version, nil
}
