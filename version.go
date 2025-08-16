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
