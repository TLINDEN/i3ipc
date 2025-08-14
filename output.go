package i3ipc

import (
	"encoding/json"
	"fmt"
)

type Mode struct {
	Width   int `json:"width"`
	Height  int `json:"height"`
	Refresh int `json:"refresh"`
}

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

func (ipc *I3ipc) GetOutputs() ([]*Output, error) {
	err := ipc.sendHeader(GET_OUTPUTS, 0)
	if err != nil {
		return nil, err
	}

	payload, err := ipc.readResponse()
	if err != nil {
		return nil, err
	}

	workspaces := []*Output{}
	if err := json.Unmarshal(payload, &workspaces); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return workspaces, nil
}
