package swayipc

import (
	"encoding/json"
	"fmt"
)

// Store an output mode.
type Mode struct {
	Width   int `json:"width"`
	Height  int `json:"height"`
	Refresh int `json:"refresh"`
}

// An output object (i.e. a physical monitor)
type Output struct {
	Name              string  `json:"name"`
	Make              string  `json:"make"`
	Serial            string  `json:"serial"`
	Active            bool    `json:"active"`
	Primary           bool    `json:"primary"`
	SubpixelHinting   string  `json:"subpixel_hinting"`
	Transform         string  `json:"transform"`
	Current_workspace string  `json:"current_workspace"`
	Modes             []*Mode `json:"modes"`
	CurrentMode       *Mode   `json:"current_mode"`
}

// Get a list of currently available and usable outputs.
func (ipc *SwayIPC) GetOutputs() ([]*Output, error) {
	err := ipc.sendHeader(GET_OUTPUTS, 0)
	if err != nil {
		return nil, err
	}

	payload, err := ipc.readResponse()
	if err != nil {
		return nil, err
	}

	workspaces := []*Output{}
	if err := json.Unmarshal(payload.Payload, &workspaces); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return workspaces, nil
}
