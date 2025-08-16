/*
Copyright © 2025 Thomas von Dein

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
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
